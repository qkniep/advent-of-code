mod years;

use std::time::Instant;

use aoc::core::*;
use aoc::measure;
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
    let Args { year, day, .. } = args;

    let title = format!("⁘⁙⁘⁙⁘ Advent of Code {} ⁘⁙⁘⁙⁘", year);
    println!("{}", title.bold().green());
    let subtitle = format!("Day {}: {}", day, "Historian Hysteria");
    println!("{}", subtitle.bold());

    // load correct input file
    let start = Instant::now();
    let input = input::load_input(year, day);
    println!(
        " -> Read input file {}",
        format_duration(start.elapsed()).bright_black()
    );

    match year {
        2024 => match day {
            1 => {
                let parsed = measure!("day1-parse", y2024::day01::parse(&input));
                println!(" -> Parse input {}", "106.5 µs".bright_black());
                if args.part.is_none() || args.part.unwrap() == 1 {
                    let parsed = parsed.clone();
                    println!(
                        " |> Part 1 {}\t{}",
                        "63.3 µs".bright_black(),
                        measure!("day1-part1", y2024::day01::part1(parsed)),
                    );
                }
                if args.part.is_none() || args.part.unwrap() == 2 {
                    println!(
                        " |> Part 2 {}\t{}",
                        "663.1 µs".bright_black(),
                        measure!("day1-part2", y2024::day01::part2(parsed)),
                    );
                }
            }
            d if year >= 2025 && d <= 12 => todo!(),
            d if d <= 25 => todo!(),
            d => panic!("invalid day {}", d),
        },
        y if (2015..2025).contains(&y) => todo!(),
        y => panic!("invalid year {}", y),
    }
}
