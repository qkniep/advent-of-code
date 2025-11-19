/// Runs a day with pretty output and timing.
///
/// # Example
/// ```no_run
/// use aoc::day_run;
/// use super::years::*;
/// // run both parts
/// day_run!(y2024::day01, None);
/// // run only part 1
/// day_run!(y2024::day01, Some(1));
/// ```
#[macro_export]
macro_rules! day_run {
    ($day:path, $part:expr) => {{
        use aoc::core::format_duration;
        use colored::Colorize;
        use std::time::Instant;

        use $day as day;
        let subtitle = format!("Day {}: {}", day::DAY, day::NAME);
        println!("{}", subtitle.bold());

        // load correct input file
        let start = Instant::now();
        let input = input::load_input(day::YEAR, day::DAY);
        println!(
            " -> Read input file {}",
            format_duration(start.elapsed()).bright_black()
        );

        // parse input
        let start = Instant::now();
        let parsed = day::parse(&input);
        println!(
            " -> Parse input {}",
            format_duration(start.elapsed()).bright_black()
        );

        if $part.is_none() || $part.unwrap() == 1 {
            let parsed = parsed.clone();
            let start = Instant::now();
            let output = day::part1(parsed);
            println!(
                " |> Part 1 {}\t{}",
                format_duration(start.elapsed()).bright_black(),
                output,
            );
        }
        if $part.is_none() || $part.unwrap() == 2 {
            let start = Instant::now();
            let output = day::part2(parsed);
            println!(
                " |> Part 2 {}\t{}",
                format_duration(start.elapsed()).bright_black(),
                output,
            );
        }
    }};
}
