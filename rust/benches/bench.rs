use aoc::core::*;
use aoc::years;

fn main() {
    divan::main();
}

#[divan::bench_group(sample_size = 100, sample_count = 1000)]
mod y2025 {
    use super::*;

    #[divan::bench]
    fn day01(bencher: divan::Bencher) {
        let input = input::load_input(2025, 1);
        let parsed = years::y2025::day01::parse(&input);

        bencher.with_inputs(|| parsed.clone()).bench_values(|p| {
            let _ = years::y2025::day01::part1(p);
        });
    }

    #[divan::bench]
    fn day02(bencher: divan::Bencher) {
        let input = input::load_input(2025, 2);
        let parsed = years::y2025::day02::parse(&input);

        bencher.with_inputs(|| parsed.clone()).bench_values(|p| {
            let _ = years::y2025::day02::part1_iter(p);
        });
    }

    #[divan::bench_group]
    mod day03 {
        use super::*;

        #[divan::bench]
        fn part1(bencher: divan::Bencher) {
            let input = input::load_input(2025, 3);
            let parsed = years::y2025::day03::parse(&input);

            bencher.with_inputs(|| parsed.clone()).bench_values(|p| {
                let _ = years::y2025::day03::part1_custom_max(p);
            });
        }

        #[divan::bench]
        fn part2(bencher: divan::Bencher) {
            let input = input::load_input(2025, 3);
            let parsed = years::y2025::day03::parse(&input);

            bencher.with_inputs(|| parsed.clone()).bench_values(|p| {
                let _ = years::y2025::day03::part2_custom_max(p);
            });
        }
    }
}
