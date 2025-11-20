use aoc::core::*;

use crate::day_tests;

pub const NAME: &str = "Mull It Over";
pub const YEAR: u16 = 2024;
pub const DAY: u16 = 3;

enum State {
    Empty,
    M,
    U,
    L,
    LParen,
    A1,
    A2,
    A3,
    Comma,
    B1,
    B2,
    B3,
    RParen,
}

pub fn parse(input: &str) -> &[u8] {
    input.as_bytes()
}

pub fn part1(input: &[u8]) -> u32 {
    let mut sum = 0;
    let mut i = 0;

    while i + 8 <= input.len() {
        // look for "mul("
        if &input[i..i + 4] != b"mul(" {
            i += 1;
            continue;
        }
        i += 4;

        // parse X (1–3 digits)
        let (a, consumed) = parse_num(&input[i..], 3);
        if consumed == 0 {
            continue;
        }
        i += consumed;

        // must have comma
        if i >= input.len() || input[i] != b',' {
            continue;
        }
        i += 1;

        // parse Y (1–3 digits)
        let (b, consumed) = parse_num(&input[i..], 3);
        if consumed == 0 {
            continue;
        }
        i += consumed;

        // must have closing paren
        if i >= input.len() || input[i] != b')' {
            continue;
        }
        i += 1;

        sum += a * b;
    }

    sum
}

pub fn part1_alternatives() -> Vec<(&'static str, fn(&[u8]) -> u32)> {
    vec![]
}

pub fn part2(input: &[u8]) -> u32 {
    let mut enabled = true;
    let mut sum = 0;
    let mut i = 0;

    while i + 8 <= input.len() {
        // look for "mul(", "do()", or "don't()"
        if enabled && &input[i..i + 4] == b"mul(" {
            i += 4;
        } else if !enabled && &input[i..i + 4] == b"do()" {
            i += 4;
            enabled = true;
            continue;
        } else if enabled && &input[i..i + 7] == b"don't()" {
            i += 7;
            enabled = false;
            continue;
        } else {
            i += 1;
            continue;
        }

        // parse X (1–3 digits)
        let (a, consumed) = parse_num(&input[i..], 3);
        if consumed == 0 {
            continue;
        }
        i += consumed;

        // must have comma
        if i >= input.len() || input[i] != b',' {
            continue;
        }
        i += 1;

        // parse Y (1–3 digits)
        let (b, consumed) = parse_num(&input[i..], 3);
        if consumed == 0 {
            continue;
        }
        i += consumed;

        // must have closing paren
        if i >= input.len() || input[i] != b')' {
            continue;
        }
        i += 1;

        sum += a * b;
    }

    sum
}

pub fn part2_alternatives() -> Vec<(&'static str, fn(&[u8]) -> u32)> {
    vec![]
}

/// Parses a number with at most `max_digits` digits.
///
/// Returns `(value, number_of_digits)`.
fn parse_num(input: &[u8], max_digits: usize) -> (u32, usize) {
    let mut val = 0;
    let mut count = 0;

    for c in input.iter().take(max_digits) {
        if !c.is_ascii_digit() {
            break;
        }
        val = val * 10 + (c - b'0') as u32;
        count += 1;
    }

    (val, count)
}

day_tests!("162813399", "53783319");
