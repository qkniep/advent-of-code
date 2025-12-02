use std::str::FromStr;

use aoc::core::*;
use thiserror::Error;

use crate::day_tests;

pub const NAME: &str = "Gift Shop";
pub const YEAR: u16 = 2025;
pub const DAY: u16 = 2;

pub type Parsed = Vec<ProductIdRange>;
pub type Output = u64;

#[derive(Clone, Copy, Debug)]
pub struct ProductIdRange(u64, u64);

impl ProductIdRange {
    /// Returns all IDs in the range that are NOT valid.
    fn invalid_ids_in_range(&self) -> Vec<u64> {
        let mut invalid_ids = Vec::new();
        for id in self.0..=self.1 {
            if !Self::is_valid_id(id) {
                invalid_ids.push(id);
            }
        }
        invalid_ids
    }

    /// Returns all IDs in the range that are NOT valid.
    ///
    /// Instead of iterating through all numbers, it uses the following observation:
    /// After any invalid ID there is at least a `10^(num_digits / 2) + 1` gap before the next one.
    /// After encountering an invalid ID, we thus jump by that amount.
    fn invalid_ids_in_range_jump(&self) -> Vec<u64> {
        let mut invalid_ids = Vec::new();
        let mut id = self.0;
        while id <= self.1 {
            if !Self::is_valid_id(id) {
                invalid_ids.push(id);
                let num_digits = id.ilog10() + 1;
                id += 10u64.pow(num_digits / 2) + 1;
            } else {
                id += 1;
            }
        }
        invalid_ids
    }

    /// Returns the sum of all invalid IDs in the range.
    ///
    /// Uses the same jumping trick as [`Self::invalid_ids_in_range_jump`].
    fn sum_invalid_ids_in_range(&self) -> u64 {
        let mut sum = 0;
        let mut id = self.0;
        while id <= self.1 {
            if !Self::is_valid_id(id) {
                sum += id;
                let num_digits = id.ilog10() + 1;
                id += 10u64.pow(num_digits / 2) + 1;
            } else {
                id += 1;
            }
        }
        sum
    }

    /// Returns the sum of all invalid IDs in the range.
    ///
    /// Iterates exactly through all the invalid IDs, and only those.
    /// For this it relies on [`Self::next_invalid_id`].
    fn sum_invalid_ids_in_range_iter(&self) -> u64 {
        let mut sum = 0;
        let mut id = Self::next_invalid_id(self.0, false);
        while id <= self.1 {
            sum += id;
            id = Self::next_invalid_id(id, true);
        }
        sum
    }

    /// Returns whether the ID is valid.
    ///
    /// If the number does not consist of two equal parts, it is valid.
    fn is_valid_id(id: u64) -> bool {
        let num_digits = id.ilog10() + 1;
        if !num_digits.is_multiple_of(2) {
            return true;
        }
        if id / (10u64.pow(num_digits / 2)) == id % (10u64.pow(num_digits / 2)) {
            return false;
        }
        true
    }

    /// Returns the next invalid ID.
    ///
    /// If `id_is_invalid`, we use the fact that `id` is invalid and also skip it.
    /// Otherwise, this works for any `id` and outputs `id` if it is itself invalid.
    fn next_invalid_id(id: u64, id_is_invalid: bool) -> u64 {
        Self::next_invalid_id_for_prime::<2>(id, id_is_invalid)
    }

    /// Returns all IDs in the range that are NOT strictly valid.
    fn invalid_ids_in_range_strict(&self) -> Vec<u64> {
        let mut invalid_ids = Vec::new();
        for id in self.0..=self.1 {
            if !Self::is_valid_id_strict(id) {
                invalid_ids.push(id);
            }
        }
        invalid_ids
    }

    /// Returns the sum of all NOT strictly valid IDs in the range.
    fn sum_invalid_ids_in_range_strict(&self) -> u64 {
        let mut sum = 0;
        let mut id = Self::next_invalid_id_strict(self.0, false);
        while id <= self.1 {
            sum += id;
            id = Self::next_invalid_id_strict(id + 1, false);
        }
        sum
    }

    /// Returns whether the ID is strictly valid.
    ///
    /// If the number does not consist of `k` equal parts for any `k`, it is valid.
    fn is_valid_id_strict(id: u64) -> bool {
        const PRIMES: [u32; 4] = [2, 3, 5, 7];
        let num_digits = id.ilog10() + 1;
        for num_parts in PRIMES {
            if !num_digits.is_multiple_of(num_parts) {
                continue;
            }
            let len = num_digits / num_parts;
            let last_part = id % (10u64.pow(len));
            if (1..num_parts)
                .map(|part| id / (10u64.pow(part * len)) % (10u64.pow(len)))
                .all(|part| part == last_part)
            {
                return false;
            }
        }
        true
    }

    /// Returns the next ID that is NOT strictly valid.
    ///
    /// If `id_is_invalid`, we use the fact that `id` is invalid and also skip it.
    /// Otherwise, this works for any `id` and outputs `id` if it is itself invalid.
    fn next_invalid_id_strict(id: u64, id_is_invalid: bool) -> u64 {
        let next_ids = [
            Self::next_invalid_id_for_prime::<2>(id, id_is_invalid),
            Self::next_invalid_id_for_prime::<3>(id, id_is_invalid),
            Self::next_invalid_id_for_prime::<5>(id, id_is_invalid),
            Self::next_invalid_id_for_prime::<7>(id, id_is_invalid),
        ];
        *next_ids.iter().min().unwrap()
    }

    /// Returns the next ID that is NOT `P`-valid.
    fn next_invalid_id_for_prime<const P: u32>(mut id: u64, id_is_invalid: bool) -> u64 {
        let num_digits = id.ilog10() + 1;
        if !num_digits.is_multiple_of(P) {
            let mut x = 0;
            for i in 0..P {
                x += 10u64.pow(num_digits / P) * 10u64.pow(num_digits / P + 1).pow(i);
            }
            return x;
        }
        if id_is_invalid {
            if id == 10u64.pow(num_digits) - 1 {
                let mut x = 0;
                for i in 0..P {
                    x += 10u64.pow(num_digits / P) * 10u64.pow(num_digits / P + 1).pow(i);
                }
                return x;
            }
            for i in 0..P {
                id += 10u64.pow(num_digits / P).pow(i);
            }
            id
        } else {
            let part_digits = 10u64.pow(num_digits / P);
            let mut need_to_inc = false;
            let first_part = id / part_digits.pow(P - 1) % part_digits;
            for i in 1..P {
                let later_part = id / part_digits.pow(P - 1 - i) % part_digits;
                if later_part > first_part {
                    need_to_inc = true;
                } else if later_part < first_part {
                    break;
                }
            }
            let mut x = 0;
            for i in 0..P {
                if need_to_inc {
                    x += (first_part + 1) * part_digits.pow(i);
                } else {
                    x += first_part * part_digits.pow(i);
                }
            }
            x
        }
    }
}

#[derive(Debug, Error)]
pub enum DirectionParseError {
    #[error("separator '-' not found")]
    NoSeparator,
    #[error("invalid ProductID: {0}")]
    InvalidProductId(#[from] std::num::ParseIntError),
}

impl FromStr for ProductIdRange {
    type Err = DirectionParseError;
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let (min, max) = s.split_once('-').ok_or(DirectionParseError::NoSeparator)?;
        let min: u64 = min.parse().map_err(DirectionParseError::InvalidProductId)?;
        let max: u64 = max.parse().map_err(DirectionParseError::InvalidProductId)?;
        Ok(ProductIdRange(min, max))
    }
}

pub fn parse(list: &str) -> Parsed {
    input::split_to_numbers(list, ',')
}

pub fn part1(ranges: Parsed) -> Output {
    ranges
        .into_iter()
        .map(|range| range.invalid_ids_in_range().into_iter().sum::<u64>())
        .sum()
}

pub fn part1_jump(ranges: Parsed) -> Output {
    ranges
        .into_iter()
        .map(|range| range.invalid_ids_in_range_jump().into_iter().sum::<u64>())
        .sum()
}

pub fn part1_cum_sum(ranges: Parsed) -> Output {
    ranges
        .into_iter()
        .map(|range| range.sum_invalid_ids_in_range())
        .sum()
}

pub fn part1_iter(ranges: Parsed) -> Output {
    ranges
        .into_iter()
        .map(|range| range.sum_invalid_ids_in_range_iter())
        .sum()
}

pub fn part1_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![
        ("jumps", part1_jump),
        ("cum sum", part1_cum_sum),
        ("iter", part1_iter),
    ]
}

pub fn part2(ranges: Parsed) -> Output {
    ranges
        .into_iter()
        .map(|range| range.invalid_ids_in_range_strict().into_iter().sum::<u64>())
        .sum()
}

pub fn part2_iter(ranges: Parsed) -> Output {
    ranges
        .into_iter()
        .map(|range| range.sum_invalid_ids_in_range_strict())
        .sum()
}

pub fn part2_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![("iter", part2_iter)]
}

day_tests!("23701357374", "34284458938");
