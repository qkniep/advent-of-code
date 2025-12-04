use rustc_hash::FxHashSet;

use crate::core::*;
use crate::day_tests;

pub const NAME: &str = "Guard Gallivant";
pub const YEAR: u16 = 2024;
pub const DAY: u16 = 6;

pub type Parsed = Map;
pub type Output = u32;

#[derive(Clone, Copy, Debug, Hash, PartialEq, Eq)]
pub enum Orientation {
    Up,
    Down,
    Left,
    Right,
}

#[derive(Clone, Debug)]
pub struct Map {
    grid: grid::Grid<bool>,
    guard_pos: (usize, usize),
    guard_orientation: Orientation,
}

impl Map {
    fn next_guard_pos(&self) -> Option<(usize, usize)> {
        let (mut x, mut y) = self.guard_pos;
        match self.guard_orientation {
            Orientation::Up => y -= 1,
            Orientation::Down => y += 1,
            Orientation::Left => x -= 1,
            Orientation::Right => x += 1,
        }
        if self.grid.get(x, y)? {
            None
        } else {
            Some((x, y))
        }
    }

    fn guard_step(&mut self) -> Option<(usize, usize)> {
        let (mut x, mut y) = self.guard_pos;
        match self.guard_orientation {
            Orientation::Up => y -= 1,
            Orientation::Down => y += 1,
            Orientation::Left => x -= 1,
            Orientation::Right => x += 1,
        }
        if self.grid.get(x, y)? {
            self.guard_orientation = match self.guard_orientation {
                Orientation::Up => Orientation::Right,
                Orientation::Down => Orientation::Left,
                Orientation::Left => Orientation::Up,
                Orientation::Right => Orientation::Down,
            };
        } else {
            self.guard_pos = (x, y);
        }
        Some(self.guard_pos)
    }

    fn causes_loop(&self, obstacle_placement: (usize, usize)) -> bool {
        let mut states = FxHashSet::default();
        let mut map = self.clone();
        let cell = map
            .grid
            .get_mut(obstacle_placement.0, obstacle_placement.1)
            .unwrap();
        *cell = true;
        states.insert((map.guard_pos, map.guard_orientation));
        loop {
            let Some(_) = map.guard_step() else {
                break;
            };
            if !states.insert((map.guard_pos, map.guard_orientation)) {
                return true;
            }
        }
        false
    }
}

pub fn parse(input: &str) -> Parsed {
    let mut guard_pos = None;
    let mut grid = Vec::new();
    for (y, line) in input.lines().enumerate() {
        let mut row = Vec::new();
        for (x, c) in line.chars().enumerate() {
            row.push(c == '#');
            if c == '^' {
                guard_pos = Some((x, y));
            }
        }
        grid.push(row);
    }
    Map {
        grid: grid::Grid::from_vec(grid),
        guard_pos: guard_pos.unwrap(),
        guard_orientation: Orientation::Up,
    }
}

pub fn part1(mut map: Parsed) -> Output {
    let mut visited = FxHashSet::default();
    visited.insert(map.guard_pos);
    loop {
        let Some(new_pos) = map.guard_step() else {
            break;
        };
        visited.insert(new_pos);
    }
    visited.len() as u32
}

pub fn part1_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![]
}

pub fn part2(mut map: Parsed) -> Output {
    let mut possible_positions = 0;
    let mut visited = FxHashSet::default();
    visited.insert(map.guard_pos);
    loop {
        if let Some(new_pos) = map.next_guard_pos()
            && !visited.contains(&new_pos)
            && map.causes_loop(new_pos)
        {
            possible_positions += 1;
        };
        let Some(new_pos) = map.guard_step() else {
            break;
        };
        visited.insert(new_pos);
    }
    possible_positions
}

pub fn part2_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![]
}

day_tests!("4374", "1705");
