mod years;

use clap::Parser;

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
    #[clap(short, long, default_value_t = 0)]
    part: u8,

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

    // load correct input file
    let path = format!("../data/{}/inputs/day{:02}.txt", year, day);
    let input = std::fs::read_to_string(path).unwrap();

    match year {
        2024 => match day {
            1 => println!("Part 1: {}", y2024::day01::part1(&input)),
            d if d <= 25 => todo!(),
            d => panic!("invalid day {}", d),
        },
        y if (2015..2025).contains(&y) => todo!(),
        y => panic!("invalid year {}", y),
    }
}
