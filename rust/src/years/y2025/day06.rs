use std::str::FromStr;

use smallvec::SmallVec;

use crate::core::*;
use crate::day_tests;

pub const NAME: &str = "Trash Compactor";
pub const YEAR: u16 = 2025;
pub const DAY: u16 = 6;

pub type Parsed = Vec<String>;
pub type Output = u64;

#[derive(Clone, Debug)]
pub enum Problem {
    Add(SmallVec<[u32; 8]>),
    Mul(SmallVec<[u32; 8]>),
}

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum ProblemType {
    Add,
    Mul,
}

impl Problem {
    fn from_lists(numbers: Vec<Vec<u32>>, problem_types: Vec<ProblemType>) -> Vec<Self> {
        for list in &numbers {
            assert!(list.len() == problem_types.len());
        }
        let mut problems = Vec::new();
        for i in 0..problem_types.len() {
            let mut nums = SmallVec::new();
            for num_list in &numbers {
                nums.push(num_list[i]);
            }
            let problem = match problem_types[i] {
                ProblemType::Add => Problem::Add(nums),
                ProblemType::Mul => Problem::Mul(nums),
            };
            problems.push(problem);
        }
        problems
    }

    fn solve(&self) -> u64 {
        match self {
            Problem::Add(v) => v.iter().map(|x| *x as u64).sum::<u64>(),
            Problem::Mul(v) => v.iter().map(|x| *x as u64).product::<u64>(),
        }
    }
}

impl FromStr for ProblemType {
    type Err = ();
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "+" => Ok(ProblemType::Add),
            "*" => Ok(ProblemType::Mul),
            _ => Err(()),
        }
    }
}

pub fn parse(input: &str) -> Parsed {
    input.lines().map(|x| x.to_string()).collect()
}

pub fn parse_part1(lines: impl IntoIterator<Item = String>) -> Vec<Problem> {
    // parse lines
    let mut numbers = Vec::new();
    let mut problem_types = Vec::new();
    for line in lines {
        let line = line.trim();
        if line.chars().next().unwrap().is_ascii_digit() {
            let list = input::parse_words::<u32>(line);
            numbers.push(list);
        } else {
            problem_types = input::parse_words::<ProblemType>(line);
        }
    }

    // turn lists into problems
    Problem::from_lists(numbers, problem_types)
}

pub fn parse_part2(lines: Vec<String>) -> Vec<Problem> {
    for line in &lines {
        assert!(line.len() == lines[0].len());
    }

    let mut problems = Vec::new();
    let mut nums = SmallVec::<[u32; 8]>::new();
    for i in (0..lines[0].len()).rev() {
        let mut num = 0u32;
        let mut valid = false;
        for line in lines[0..lines.len() - 1].iter() {
            let c = line.as_bytes()[i];
            if c.is_ascii_digit() {
                num *= 10;
                num += u32::from(c - b'0');
                valid = true;
            }
        }
        if valid {
            nums.push(num);
        }
        let ptype_char = lines.last().unwrap().as_bytes()[i];
        if ptype_char != b' ' {
            match ptype_char {
                b'+' => problems.push(Problem::Add(nums)),
                b'*' => problems.push(Problem::Mul(nums)),
                _ => unreachable!(),
            }
            nums = SmallVec::new();
        }
    }
    problems
}

/// Counts how many of the available ingredients are fresh.
pub fn part1(lines: Parsed) -> Output {
    let problems = parse_part1(lines);
    let mut total = 0;
    for problem in problems {
        total += problem.solve();
    }
    total
}

pub fn part1_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![]
}

/// Counts how many fresh ingredients exist overall.
pub fn part2(lines: Parsed) -> Output {
    let problems = parse_part2(lines);
    let mut total = 0;
    for problem in problems {
        total += problem.solve();
    }
    total
}

pub fn part2_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![]
}

day_tests!("513", "339668510830757");
