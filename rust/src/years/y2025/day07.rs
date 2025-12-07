use std::collections::BTreeMap;
use std::collections::BTreeSet;

use bitvec::order::Lsb0;
use bitvec::vec::BitVec;
use rustc_hash::FxHashMap;
use rustc_hash::FxHashSet;

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
    let height = input.lines().count() / 2;
    let first_line = input.lines().next().unwrap();
    let width = first_line.len();

    // read origin position
    let origin_x = first_line.chars().position(|c| c == 'S').unwrap();
    let origin = (origin_x, 0);

    // read grid
    let mut grid = binary_grid::BinaryGrid::new(width, height, false);
    for (y, line) in input.lines().enumerate() {
        // ignore empty rows
        if y % 2 == 1 {
            continue;
        }
        for (x, c) in line.chars().enumerate() {
            if c == '^' {
                grid.set(x, y / 2, true);
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

pub fn part1_fxhash(diagram: Parsed) -> Output {
    let mut splits = 0;
    let mut active_beams = FxHashSet::default();
    active_beams.insert(diagram.origin.0);

    for y in 0..diagram.grid.height() {
        let mut new_active_beams = FxHashSet::default();
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

pub fn part1_alloc_once(diagram: Parsed) -> Output {
    let mut splits = 0;
    let mut active_beams = FxHashSet::default();
    active_beams.insert(diagram.origin.0);
    let mut new_active_beams = FxHashSet::default();

    for y in 0..diagram.grid.height() {
        new_active_beams.clear();
        for beam_x in &active_beams {
            if diagram.grid.get(*beam_x, y) {
                splits += 1;
                new_active_beams.insert(*beam_x - 1);
                new_active_beams.insert(*beam_x + 1);
            } else {
                new_active_beams.insert(*beam_x);
            }
        }
        std::mem::swap(&mut active_beams, &mut new_active_beams);
    }

    splits
}

pub fn part1_bitvec(diagram: Parsed) -> Output {
    let mut splits = 0;
    let mut active_beams = BitVec::<usize, Lsb0>::repeat(false, diagram.grid.width());
    active_beams.set(diagram.origin.0, true);
    let mut new_active_beams = BitVec::repeat(false, diagram.grid.width());

    for y in 0..diagram.grid.height() {
        new_active_beams.fill(false);
        for beam_x in active_beams.iter_ones() {
            if diagram.grid.get(beam_x, y) {
                splits += 1;
                new_active_beams.set(beam_x - 1, true);
                new_active_beams.set(beam_x + 1, true);
            } else {
                new_active_beams.set(beam_x, true);
            }
        }
        std::mem::swap(&mut active_beams, &mut new_active_beams);
    }

    splits
}

pub fn part1_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![
        ("fxhash", part1_fxhash),
        ("alloc once", part1_alloc_once),
        ("bitvec", part1_bitvec),
    ]
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

pub fn part2_fxhash(diagram: Parsed) -> Output {
    let mut active_beams = FxHashMap::default();
    active_beams.insert(diagram.origin.0, 1);

    for y in 0..diagram.grid.height() {
        let mut new_active_beams = FxHashMap::default();
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

pub fn part2_alloc_once(diagram: Parsed) -> Output {
    let mut active_beams = FxHashMap::default();
    active_beams.insert(diagram.origin.0, 1);
    let mut new_active_beams = FxHashMap::default();

    for y in 0..diagram.grid.height() {
        new_active_beams.clear();
        for (beam_x, multiplicity) in &active_beams {
            if diagram.grid.get(*beam_x, y) {
                *new_active_beams.entry(*beam_x - 1).or_insert(0) += *multiplicity;
                *new_active_beams.entry(*beam_x + 1).or_insert(0) += *multiplicity;
            } else {
                *new_active_beams.entry(*beam_x).or_insert(0) += *multiplicity;
            }
        }
        std::mem::swap(&mut active_beams, &mut new_active_beams);
    }

    active_beams.values().sum()
}

pub fn part2_vec(diagram: Parsed) -> Output {
    let mut active_beams = vec![0u64; diagram.grid.width()];
    active_beams[diagram.origin.0] = 1;
    let mut new_active_beams = vec![0u64; diagram.grid.width()];

    for y in 0..diagram.grid.height() {
        new_active_beams.fill(0);
        for (beam_x, multiplicity) in active_beams.iter().enumerate() {
            if diagram.grid.get(beam_x, y) {
                new_active_beams[beam_x - 1] += *multiplicity;
                new_active_beams[beam_x + 1] += *multiplicity;
            } else {
                new_active_beams[beam_x] += *multiplicity;
            }
        }
        std::mem::swap(&mut active_beams, &mut new_active_beams);
    }

    active_beams.iter().sum()
}

pub fn part2_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![
        ("fxhash", part2_fxhash),
        ("alloc once", part2_alloc_once),
        ("vec", part2_vec),
    ]
}

day_tests!("1499", "24743903847942");
