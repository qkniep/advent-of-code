mod years;
use years::*;

fn main() {
    let args: Vec<String> = std::env::args().collect();
    let year = args.get(1).expect("year provided");
    let year: u16 = year.parse().expect("year is number");
    let day = args.get(2).expect("day provided");
    let day: u16 = day.parse().expect("day is number");
    let input = std::fs::read_to_string(format!("inputs/y{}/day{}.txt", year, day)).unwrap();
    match year {
        2024 => match day {
            1 => println!("Part 1: {}", y2024::day01::part1(&input)),
            d if d <= 25 => unimplemented!(),
            d => panic!("invalid day {}", d),
        },
        y if (2015..2025).contains(&y) => unimplemented!(),
        y => panic!("invalid year {}", y),
    }
}
