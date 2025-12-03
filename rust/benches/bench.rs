fn main() {
    divan::main();
}

#[divan::bench_group]
mod y2025 {
    use aoc::core::*;
    use aoc::years::*;

    #[divan::bench]
    fn day01(bencher: divan::Bencher) {
        let input = input::load_input(2025, 1);
        let parsed = y2025::day01::parse(&input);

        bencher.with_inputs(|| parsed.clone()).bench_values(|p| {
            let _ = y2025::day01::part1(p);
        });
    }

    #[divan::bench]
    fn day02(bencher: divan::Bencher) {
        let input = input::load_input(2025, 2);
        let parsed = y2025::day02::parse(&input);

        bencher.with_inputs(|| parsed.clone()).bench_values(|p| {
            let _ = y2025::day02::part1_iter(p);
        });
    }

    #[divan::bench]
    fn day03(bencher: divan::Bencher) {
        let input = input::load_input(2025, 3);
        let parsed = y2025::day03::parse(&input);

        bencher.with_inputs(|| parsed.clone()).bench_values(|p| {
            let _ = y2025::day03::part1(p);
        });
    }
}
