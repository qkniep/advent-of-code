use std::fmt::Debug;
use std::str::FromStr;

/// Split input into lines (no allocation).
pub fn lines(input: &str) -> impl Iterator<Item = &str> {
    input.lines()
}

/// Split and parse each line to numbers.
pub fn lines_as_numbers<T: FromStr>(input: &str) -> impl Iterator<Item = T>
where
    <T as FromStr>::Err: Debug,
{
    input
        .lines()
        .map(|line| line.parse::<T>().expect("failed to parse number"))
}

/// Split a string by a character and parse to numbers.
pub fn split_to_numbers<T: FromStr>(s: &str, sep: char) -> Vec<T>
where
    <T as FromStr>::Err: Debug,
{
    s.split(sep)
        .map(|x| x.parse::<T>().expect("failed to parse number"))
        .collect()
}

/// Split string by multiple separators (like regex-free CSV).
pub fn split_multi<T: Debug + FromStr>(s: &str, seps: &[char]) -> Vec<T>
where
    <T as FromStr>::Err: Debug,
{
    s.split(|c| seps.contains(&c))
        .map(|x| x.parse::<T>().expect("failed to parse number"))
        .collect()
}
