/// Generates tests for a day module.
///
/// # Example
/// ```no_run
/// use aoc::day_tests;
/// day_tests!(2024, 1, "123", "456");
/// ```
#[macro_export]
macro_rules! day_tests {
    ($year:literal, $day:literal, $expected1:expr, $expected2:expr) => {
        #[cfg(test)]
        mod tests {
            use aoc::core::input::load_input;

            use super::*;

            #[test]
            fn part1() {
                let input = load_input($year, $day);
                let parsed = parse(&input);
                let result = super::part1(parsed).to_string();
                assert_eq!(result, $expected1, "Part1 failed for {}", stringify!($day));
            }

            #[test]
            fn part2() {
                let input = load_input($year, $day);
                let parsed = parse(&input);
                let result = super::part2(parsed).to_string();
                assert_eq!(result, $expected2, "Part2 failed for {}", stringify!($day));
            }
        }
    };
}
