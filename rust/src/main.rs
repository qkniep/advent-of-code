mod day_run_macro;
mod day_tests_macro;
mod years;

use aoc::core::*;
use clap::Parser;
use colored::Colorize;

use self::years::*;

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
            d if year >= 2025 && d <= 12 => todo!(),
            d if d <= 25 => todo!(),
            d => panic!("invalid day {}", d),
        },
        y if (2015..2025).contains(&y) => todo!(),
        y => panic!("invalid year {}", y),
    }
}
