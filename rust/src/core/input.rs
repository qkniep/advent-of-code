//! Input parsing utilities.
//!
//!

use std::fmt::Debug;
use std::str::FromStr;

/// Loads input file for given `year` and `day` into a string.
pub fn load_input(year: u16, day: u16) -> String {
    let path = format!("../data/{}/inputs/day{:02}.txt", year, day);
    std::fs::read_to_string(path).unwrap()
}

/// Splits `input` into lines (no allocation).
pub fn lines(input: &str) -> impl Iterator<Item = &str> {
    input.lines()
}

/// Splits `input` into lines and parses each line.
pub fn lines_as_numbers<T: FromStr>(input: &str) -> impl Iterator<Item = T>
where
    <T as FromStr>::Err: Debug,
{
    input
        .lines()
        .map(|line| line.parse::<T>().expect("failed to parse number"))
}

/// Splits a string by ASCII whitespace and parses each such "word".
pub fn parse_words<T: FromStr>(s: &str) -> Vec<T>
where
    <T as FromStr>::Err: Debug,
{
    s.split_ascii_whitespace()
        .map(|x| x.parse::<T>().expect("failed to parse number"))
        .collect()
}

/// Splits a string by a character and parses to numbers.
pub fn split_to_numbers<T: FromStr>(s: &str, sep: char) -> Vec<T>
where
    <T as FromStr>::Err: Debug,
{
    s.split(sep)
        .map(|x| x.trim().parse::<T>().expect("failed to parse number"))
        .collect()
}

/// Splits string by multiple separators.
pub fn split_multi<T: Debug + FromStr>(s: &str, seps: &[char]) -> Vec<T>
where
    <T as FromStr>::Err: Debug,
{
    s.split(|c| seps.contains(&c))
        .map(|x| x.parse::<T>().expect("failed to parse number"))
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn basic() {}
}
