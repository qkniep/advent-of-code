use crate::core::*;
use crate::day_tests;

pub const NAME: &str = "Printing Department";
pub const YEAR: u16 = 2025;
pub const DAY: u16 = 4;

pub type Parsed = grid::Grid<Stack>;
pub type Output = u64;

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum Stack {
    Empty,
    Stack(Option<u8>),
}

pub fn parse(input: &str) -> Parsed {
    let mut vec = Vec::new();
    for line in input.lines() {
        let mut row = Vec::new();
        for c in line.bytes() {
            if c == b'.' {
                row.push(Stack::Empty);
            } else if c == b'@' {
                row.push(Stack::Stack(None));
            } else {
                unreachable!()
            }
        }
        vec.push(row);
    }
    grid::Grid::from_vec(vec)
}

pub fn part1(grid: Parsed) -> Output {
    let mut accessible = 0;
    for ((x, y), cell) in grid.iter() {
        if *cell == Stack::Empty {
            continue;
        }
        let mut count_neighbors = 0;
        for (nx, ny) in grid.neighbors8(x, y) {
            if let Stack::Stack(_) = grid.get(nx, ny).unwrap() {
                count_neighbors += 1;
            }
        }
        if count_neighbors < 4 {
            accessible += 1;
        }
    }
    accessible
}

pub fn part1_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![]
}

pub fn part2(mut grid: Parsed) -> Output {
    let mut accessible = 0;
    let mut prev_accessible = 0;
    loop {
        for y in 0..grid.height {
            for x in 0..grid.width {
                if grid.get(x, y).unwrap() == Stack::Empty {
                    continue;
                }
                let mut count_neighbors = 0;
                for (x, y) in grid.neighbors8(x, y) {
                    if let Stack::Stack(_) = grid.get(x, y).unwrap() {
                        count_neighbors += 1;
                    }
                }
                if count_neighbors < 4 {
                    accessible += 1;
                    *grid.get_mut(x, y).unwrap() = Stack::Empty;
                }
            }
        }
        if accessible == prev_accessible {
            break;
        }
        prev_accessible = accessible;
    }
    accessible
}

pub fn part2_single_pass(mut grid: Parsed) -> Output {
    let mut accessible = 0;
    let mut to_check = Vec::with_capacity(grid.width * grid.height);
    for y in 0..grid.height {
        for x in 0..grid.width {
            if let Stack::Stack(_) = grid.get(x, y).unwrap() {
                to_check.push((x, y));
            }
        }
    }
    while let Some((x, y)) = to_check.pop() {
        if grid.get(x, y).unwrap() == Stack::Empty {
            continue;
        }
        let mut count_neighbors = 0;
        for (x, y) in grid.neighbors8(x, y) {
            if let Stack::Stack(_) = grid.get(x, y).unwrap() {
                count_neighbors += 1;
            }
        }
        if count_neighbors < 4 {
            for (x, y) in grid.neighbors8(x, y) {
                if let Stack::Stack(_) = grid.get(x, y).unwrap() {
                    to_check.push((x, y));
                }
            }
            accessible += 1;
            *grid.get_mut(x, y).unwrap() = Stack::Empty;
        }
    }
    accessible
}

pub fn part2_2(mut grid: Parsed) -> Output {
    let mut accessible = 0;
    let mut to_remove = Vec::with_capacity(1600);
    for y in 0..grid.height {
        for x in 0..grid.width {
            if grid.get(x, y).unwrap() == Stack::Empty {
                continue;
            }
            let mut count_neighbors = 0;
            for (x, y) in grid.neighbors8(x, y) {
                if let Stack::Stack(_) = grid.get(x, y).unwrap() {
                    count_neighbors += 1;
                }
            }
            *grid.get_mut(x, y).unwrap() = Stack::Stack(Some(count_neighbors));
            if count_neighbors < 4 {
                to_remove.push((x, y));
            }
        }
    }
    while let Some((x, y)) = to_remove.pop() {
        if grid.get(x, y).unwrap() == Stack::Empty {
            continue;
        }
        for (nx, ny) in grid.neighbors8(x, y) {
            if let Stack::Stack(Some(n)) = grid.get_mut(nx, ny).unwrap() {
                *n -= 1;
                if *n < 4 {
                    to_remove.push((nx, ny));
                }
            }
        }
        accessible += 1;
        *grid.get_mut(x, y).unwrap() = Stack::Empty;
    }
    accessible
}

pub fn part2_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![
        ("single pass", part2_single_pass),
        ("to-remove queue", part2_2),
    ]
}

day_tests!("1493", "9194");
