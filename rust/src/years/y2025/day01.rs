use std::str::FromStr;

use thiserror::Error;

use crate::day_tests;

pub const NAME: &str = "Secret Entrance";
pub const YEAR: u16 = 2025;
pub const DAY: u16 = 1;

pub type Parsed = Vec<Direction>;
pub type Output = u32;

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum Direction {
    Left(u32),
    Right(u32),
}

#[derive(Debug, Error)]
pub enum DirectionParseError {
    #[error("empty string")]
    Empty,
    #[error("invalid prefix char: {0}")]
    InvalidPrefix(char),
    #[error("invalid number of steps: {0}")]
    InvalidSteps(#[from] std::num::ParseIntError),
}

impl FromStr for Direction {
    type Err = DirectionParseError;
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s.chars().next() {
            Some('L') => {
                let steps = s[1..].parse().map_err(DirectionParseError::InvalidSteps)?;
                Ok(Direction::Left(steps))
            }
            Some('R') => {
                let steps = s[1..].parse().map_err(DirectionParseError::InvalidSteps)?;
                Ok(Direction::Right(steps))
            }
            Some(c) => Err(DirectionParseError::InvalidPrefix(c)),
            None => Err(DirectionParseError::Empty),
        }
    }
}

struct Dial {
    offset: u32,
    min: u32,
    max: u32,
}

impl Dial {
    fn apply(&mut self, direction: Direction) {
        let modulo = self.max - self.min + 1;
        match direction {
            Direction::Left(steps) => {
                self.offset += modulo - (steps % modulo);
            }
            Direction::Right(steps) => {
                self.offset += steps % modulo;
            }
        }
        if self.offset >= modulo {
            self.offset -= modulo;
        }
    }

    fn apply_count_wraps(&mut self, direction: Direction) -> u32 {
        let modulo = self.max - self.min + 1;
        let (steps, partial_wrap) = match direction {
            Direction::Left(steps) => (steps, steps % modulo > self.offset),
            Direction::Right(steps) => (
                steps,
                steps % modulo > self.max - self.current_position() + 1,
            ),
        };
        let wraps = if partial_wrap && self.offset > 0 {
            steps / modulo + 1
        } else {
            steps / modulo
        };
        self.apply(direction);
        wraps
    }

    const fn current_position(&self) -> u32 {
        self.min + self.offset
    }
}

struct NoModuloDial {
    offset: i32,
    min: u32,
    max: u32,
}

impl NoModuloDial {
    fn apply(&mut self, direction: Direction) {
        match direction {
            Direction::Left(steps) => {
                self.offset -= steps as i32;
            }
            Direction::Right(steps) => {
                self.offset += steps as i32;
            }
        }
    }

    const fn current_position(&self) -> u32 {
        let modulo = self.max - self.min + 1;
        self.min + self.offset.rem_euclid(modulo as i32) as u32
    }
}

pub fn parse(input: &str) -> Vec<Direction> {
    let mut directions = Vec::new();
    for line in input.lines() {
        directions.push(line.parse().unwrap());
    }
    directions
}

pub fn parse_streaming(line: &str) -> Direction {
    line.parse().unwrap()
}

pub fn part1(directions: Parsed) -> Output {
    part1_streaming(directions.into_iter())
}

pub fn part1_streaming(directions: impl Iterator<Item = Direction>) -> Output {
    let mut dial = Dial {
        offset: 50,
        min: 0,
        max: 99,
    };
    let mut count = 0;
    for dir in directions {
        dial.apply(dir);
        if dial.current_position() == 0 {
            count += 1;
        }
    }
    count
}

pub fn part1_no_modulo(directions: Parsed) -> Output {
    let mut dial = NoModuloDial {
        offset: 50,
        min: 0,
        max: 99,
    };
    let mut count = 0;
    for dir in directions {
        dial.apply(dir);
        if dial.current_position() == 0 {
            count += 1;
        }
    }
    count
}

pub fn part1_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![("no mod", part1_no_modulo)]
}

pub fn part2(directions: Parsed) -> Output {
    part2_streaming(directions.into_iter())
}

pub fn part2_streaming(directions: impl Iterator<Item = Direction>) -> Output {
    let mut dial = Dial {
        offset: 50,
        min: 0,
        max: 99,
    };
    let mut count = 0;
    for dir in directions {
        count += dial.apply_count_wraps(dir);
        if dial.current_position() == 0 {
            count += 1;
        }
    }
    count
}

pub fn part2_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![]
}

day_tests!("995", "5847");
