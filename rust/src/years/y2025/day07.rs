use std::collections::BTreeMap;
use std::collections::BTreeSet;

use crate::core::*;
use crate::day_tests;

pub const NAME: &str = "Laboratories";
pub const YEAR: u16 = 2025;
pub const DAY: u16 = 7;

pub type Parsed = ManifoldDiagram;
pub type Output = u64;

#[derive(Debug, Clone)]
pub struct ManifoldDiagram {
    grid: binary_grid::BinaryGrid,
    origin: (usize, usize),
}

pub fn parse(input: &str) -> Parsed {
    let height = input.lines().count();
    let first_line = input.lines().next().unwrap();
    let width = first_line.len();

    // read origin position
    let origin_x = first_line.chars().position(|c| c == 'S').unwrap();
    let origin = (origin_x, 0);

    // read grid
    let mut grid = binary_grid::BinaryGrid::new(width, height, false);
    for (y, line) in input.lines().enumerate() {
        for (x, c) in line.chars().enumerate() {
            if c == '^' {
                grid.set(x, y, true);
            }
        }
    }

    ManifoldDiagram { grid, origin }
}

pub fn part1(diagram: Parsed) -> Output {
    let mut splits = 0;
    let mut active_beams = BTreeSet::new();
    active_beams.insert(diagram.origin.0);

    for y in 0..diagram.grid.height() {
        let mut new_active_beams = BTreeSet::new();
        for beam_x in &active_beams {
            if diagram.grid.get(*beam_x, y) {
                splits += 1;
                new_active_beams.insert(*beam_x - 1);
                new_active_beams.insert(*beam_x + 1);
            } else {
                new_active_beams.insert(*beam_x);
            }
        }
        active_beams = new_active_beams;
    }

    splits
}

pub fn part1_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![]
}

pub fn part2(diagram: Parsed) -> Output {
    let mut active_beams = BTreeMap::new();
    active_beams.insert(diagram.origin.0, 1);

    for y in 0..diagram.grid.height() {
        let mut new_active_beams = BTreeMap::new();
        for (beam_x, multiplicity) in &active_beams {
            if diagram.grid.get(*beam_x, y) {
                *new_active_beams.entry(*beam_x - 1).or_insert(0) += *multiplicity;
                *new_active_beams.entry(*beam_x + 1).or_insert(0) += *multiplicity;
            } else {
                *new_active_beams.entry(*beam_x).or_insert(0) += *multiplicity;
            }
        }
        active_beams = new_active_beams;
    }

    active_beams.values().sum()
}

pub fn part2_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![]
}

day_tests!("1499", "24743903847942");
