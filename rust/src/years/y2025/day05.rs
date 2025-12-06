use std::cmp::Ordering;
use std::str::FromStr;

use crate::core::*;
use crate::day_tests;

pub const NAME: &str = "Cafeteria";
pub const YEAR: u16 = 2025;
pub const DAY: u16 = 5;

pub type Parsed = Ingredients;
pub type Output = u64;

#[derive(Clone, Debug)]
pub struct Ingredients {
    fresh_id_ranges: Vec<FreshIdRange>,
    available_ids: Vec<u64>,
}

/// Inclusive range of fresh ingredient IDs.
#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub struct FreshIdRange(u64, u64);

impl FromStr for FreshIdRange {
    type Err = ();
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let (start, end) = s.split_once('-').unwrap();
        Ok(FreshIdRange(start.parse().unwrap(), end.parse().unwrap()))
    }
}

pub fn parse(input: &str) -> Parsed {
    let mut fresh = Vec::new();
    let mut iter = input.lines();
    for line in iter.by_ref() {
        if line.is_empty() {
            break;
        }
        let range: FreshIdRange = line.parse().unwrap();
        fresh.push(range);
    }

    let mut available = Vec::new();
    for line in iter {
        available.push(line.trim().parse().unwrap());
    }

    Ingredients {
        fresh_id_ranges: fresh,
        available_ids: available,
    }
}

/// Counts how many of the available ingredients are fresh.
pub fn part1(ingredients: Ingredients) -> Output {
    let mut count = 0;
    for id in ingredients.available_ids {
        for range in &ingredients.fresh_id_ranges {
            if id >= range.0 && id <= range.1 {
                count += 1;
                break;
            }
        }
    }
    count
}

pub fn part1_merge(ingredients: Ingredients) -> Output {
    let merged_ranges = merge_ranges(ingredients.fresh_id_ranges);
    let mut count = 0;
    for id in ingredients.available_ids {
        let cmp = |range: &FreshIdRange| {
            if id < range.0 {
                Ordering::Greater
            } else if id > range.1 {
                Ordering::Less
            } else {
                Ordering::Equal
            }
        };
        if merged_ranges.binary_search_by(cmp).is_ok() {
            count += 1;
        }
    }
    count
}

pub fn part1_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![("merge", part1_merge)]
}

/// Counts how many fresh ingredients exist overall.
pub fn part2(mut ingredients: Parsed) -> Output {
    // sort ranges
    ingredients
        .fresh_id_ranges
        .sort_unstable_by_key(|range| range.0);

    // count number of IDs covered by all ranges
    let mut fresh = 0;
    let mut prev_end = 0;
    for range in ingredients.fresh_id_ranges {
        if range.0 > prev_end {
            fresh += range.1 - range.0 + 1;
        } else {
            fresh += range.1.saturating_sub(prev_end);
        }
        prev_end = prev_end.max(range.1);
    }
    fresh
}

pub fn part2_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![]
}

fn merge_ranges(mut ranges: Vec<FreshIdRange>) -> Vec<FreshIdRange> {
    ranges.sort_unstable_by_key(|range| range.0);
    let mut merged = Vec::new();
    let mut current_range = ranges[0];
    for range in ranges.into_iter().skip(1) {
        if range.0 > current_range.1 {
            merged.push(current_range);
            current_range = range;
        } else {
            current_range.1 = current_range.1.max(range.1);
        }
    }
    merged
}

day_tests!("513", "339668510830757");
