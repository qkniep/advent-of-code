mod day_run_macro;
mod day_tests_macro;
mod years;

use aoc::core::*;
use clap::Parser;
use colored::Colorize;

use self::years::*;

const FIRST_YEAR: u16 = 2015;
const LAST_YEAR: u16 = 2025;

#[derive(Parser)]
#[clap(author, version, about)]
struct Args {
    /// Year to run
    #[clap(short, long)]
    year: u16,

    /// Day to run
    #[clap(short, long)]
    day: u16,

    /// Part of the day to run (1 or 2)
    #[clap(short, long)]
    part: Option<u8>,

    /// Input file override
    #[clap(short, long)]
    input: Option<String>,

    /// Generate ASCII/graph visualization
    #[clap(long)]
    plot: bool,
}

fn main() {
    // parse CLI arguments
    let args = Args::parse();
    let Args {
        year, day, part, ..
    } = args;

    let title = format!("⁘⁙⁘⁙⁘ Advent of Code {} ⁘⁙⁘⁙⁘", year);
    println!("{}", title.bold().green());

    match year {
        2024 => match day {
            1 => day_run!(y2024::day01, part),
            2 => day_run!(y2024::day02, part),
            3 => day_run!(y2024::day03, part),
            4 => day_run!(y2024::day04, part),
            5 => day_run!(y2024::day05, part),
            d if d > 0 && d <= days_per_year(2024) => todo!(),
            d => panic!("invalid day {}", d),
        },
        FIRST_YEAR..=LAST_YEAR => todo!(),
        y => panic!("invalid year {}", y),
    }
}

/// Returns number of days in `year`.
pub const fn days_per_year(year: u16) -> u16 {
    match year {
        FIRST_YEAR..=2024 => 25,
        2025 => 12,
        _ => 0,
    }
}
