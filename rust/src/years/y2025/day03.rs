use std::str::FromStr;

use crate::day_tests;

pub const NAME: &str = "Lobby";
pub const YEAR: u16 = 2025;
pub const DAY: u16 = 3;

pub type Parsed = Vec<Bank>;
pub type Output = u64;

#[derive(Clone, Debug)]
pub struct Bank(Vec<u8>);

impl Bank {
    /// Calculates the maximum joltage of the bank when turning on `num_batteries` batteries.
    ///
    /// That is, it finds the largest `num_batteries`-digit number appearing in `self`.
    pub fn max_joltage(&self, num_batteries: usize) -> u64 {
        let mut batteries_to_take = num_batteries;
        let mut offset = 0;
        let mut joltage = 0;
        while batteries_to_take > 0 {
            let (i, &max_digit) = self.0[offset..self.0.len() - (batteries_to_take - 1)]
                .iter()
                .enumerate()
                .rev()
                .max_by_key(|(_, d)| *d)
                .unwrap();
            batteries_to_take -= 1;
            offset += i + 1;
            joltage = joltage * 10 + max_digit as u64;
        }
        joltage
    }

    /// Calculates the maximum joltage of the bank when turning on `num_batteries` batteries.
    ///
    /// That is, it finds the largest `num_batteries`-digit number appearing in `self`.
    pub fn max_joltage_custom_max(&self, num_batteries: usize) -> u64 {
        let mut batteries_to_take = num_batteries;
        let mut offset = 0;
        let mut joltage = 0;
        while batteries_to_take > 0 {
            let (i, max_digit) =
                Bank::idx_max_early_abort(&self.0[offset..self.0.len() - (batteries_to_take - 1)]);
            batteries_to_take -= 1;
            offset += i + 1;
            joltage = joltage * 10 + max_digit as u64;
        }
        joltage
    }

    /// Calculates the maximum joltage of the bank when turning on `B` batteries.
    pub fn idx_max_early_abort(digits: &[u8]) -> (usize, u8) {
        let mut max = (usize::MAX, 0);
        for (i, &d) in digits.iter().enumerate() {
            if d == 9 {
                return (i, d);
            }
            if d > max.1 {
                max = (i, d);
            }
        }
        max
    }
}

impl FromStr for Bank {
    type Err = ();

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        Ok(Self(s.as_bytes().iter().map(|b| b - b'0').collect()))
    }
}

pub fn parse(input: &str) -> Parsed {
    input.lines().map(|l| l.parse().unwrap()).collect()
}

pub fn parse_streaming(line: &str) -> Bank {
    line.parse().unwrap()
}

pub fn part1(banks: Parsed) -> Output {
    part1_streaming(banks.into_iter())
}

pub fn part1_streaming(banks: impl Iterator<Item = Bank>) -> Output {
    let mut total_joltage = 0;
    for bank in banks {
        total_joltage += bank.max_joltage(2);
    }
    total_joltage
}

pub fn part1_single_pass(banks: Parsed) -> Output {
    let mut total_joltage = 0;
    for bank in banks {
        total_joltage += bank.max_joltage_custom_max(2);
    }
    total_joltage
}

pub fn part1_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![("custom max", part1_single_pass)]
}

pub fn part2(banks: Parsed) -> Output {
    part2_streaming(banks.into_iter())
}

pub fn part2_streaming(banks: impl Iterator<Item = Bank>) -> Output {
    let mut total_joltage = 0;
    for bank in banks {
        total_joltage += bank.max_joltage(12);
    }
    total_joltage
}

pub fn part2_single_pass(banks: Parsed) -> Output {
    let mut total_joltage = 0;
    for bank in banks {
        total_joltage += bank.max_joltage_custom_max(12);
    }
    total_joltage
}

pub fn part2_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![("custom max", part2_single_pass)]
}

day_tests!("16842", "167523425665348");
