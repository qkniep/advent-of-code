use aoc::core::*;

use crate::day_tests;

pub const NAME: &str = "Ceres Search";
pub const YEAR: u16 = 2024;
pub const DAY: u16 = 4;

type Parsed = grid::Grid<Letter>;
type Output = u32;

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum Letter {
    X,
    M,
    A,
    S,
}

impl TryFrom<char> for Letter {
    type Error = ();
    fn try_from(c: char) -> Result<Self, Self::Error> {
        match c {
            'X' => Ok(Letter::X),
            'M' => Ok(Letter::M),
            'A' => Ok(Letter::A),
            'S' => Ok(Letter::S),
            _ => Err(()),
        }
    }
}

impl From<Letter> for char {
    fn from(l: Letter) -> char {
        match l {
            Letter::X => 'X',
            Letter::M => 'M',
            Letter::A => 'A',
            Letter::S => 'S',
        }
    }
}

pub fn parse(input: &str) -> Parsed {
    grid::Grid::<Letter>::try_from(input).unwrap()
}

pub fn part1(grid: Parsed) -> Output {
    const OFFSETS: [[(i32, i32); 3]; 8] = [
        [(0, 1), (0, 2), (0, 3)],
        [(0, -1), (0, -2), (0, -3)],
        [(1, 0), (2, 0), (3, 0)],
        [(-1, 0), (-2, 0), (-3, 0)],
        [(1, 1), (2, 2), (3, 3)],
        [(1, -1), (2, -2), (3, -3)],
        [(-1, 1), (-2, 2), (-3, 3)],
        [(-1, -1), (-2, -2), (-3, -3)],
    ];
    const LETTERS: [Letter; 3] = [Letter::M, Letter::A, Letter::S];
    let mut count = 0;
    for y in 0..grid.height {
        for x in 0..grid.width {
            if grid.get(x, y) != Some(Letter::X) {
                continue;
            }
            'outer: for offsets in OFFSETS {
                for (i, offset) in offsets.iter().enumerate() {
                    let (dx, dy) = offset;
                    let xx = x as i32 + dx;
                    let yy = y as i32 + dy;
                    if xx < 0
                        || xx >= grid.width as i32
                        || yy < 0
                        || yy >= grid.height as i32
                        || grid.get(xx as usize, yy as usize) != Some(LETTERS[i])
                    {
                        continue 'outer;
                    }
                }
                count += 1;
            }
        }
    }
    count
}

pub fn part1_streaming(grid: Parsed) -> Output {
    0
}

pub fn part1_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![]
}

pub fn part2(grid: Parsed) -> Output {
    let mut count = 0;
    for y in 1..grid.height - 1 {
        'x: for x in 1..grid.width - 1 {
            if grid.get(x, y) != Some(Letter::A) {
                continue;
            }
            for (xx, yy) in grid.neighbors8(x, y) {
                if xx == x || yy == y {
                    continue;
                }
                let dx = xx as i32 - x as i32;
                let dy = yy as i32 - y as i32;
                let ox = (x as i32 - dx) as usize;
                let oy = (y as i32 - dy) as usize;
                if grid.get(xx, yy) != Some(Letter::M) && grid.get(xx, yy) != Some(Letter::S)
                    || grid.get(xx, yy) == Some(Letter::M) && grid.get(ox, oy) != Some(Letter::S)
                    || grid.get(xx, yy) == Some(Letter::S) && grid.get(ox, oy) != Some(Letter::M)
                {
                    continue 'x;
                }
            }
            count += 1;
        }
    }
    count
}

pub fn part2_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![]
}

day_tests!("2575", "2041");
