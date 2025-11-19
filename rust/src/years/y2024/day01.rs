use aoc::core::*;

use crate::day_tests;

pub const NAME: &str = "Historian Hysteria";
pub const YEAR: u16 = 2024;
pub const DAY: u16 = 1;

pub fn parse(input: &str) -> (Vec<i32>, Vec<i32>) {
    let mut list1 = Vec::new();
    let mut list2 = Vec::new();
    for line in input.lines() {
        if line.is_empty() {
            continue;
        }
        let nums = input::parse_words(line);
        list1.push(nums[0]);
        list2.push(nums[1]);
    }
    (list1, list2)
}

pub fn part1(lists: (Vec<i32>, Vec<i32>)) -> String {
    let (mut list1, mut list2) = lists;
    let mut sum = 0;
    assert_eq!(list1.len(), list2.len());
    list1.sort_unstable();
    list2.sort_unstable();
    for i in 0..list1.len() {
        sum += (list1[i] - list2[i]).abs();
    }
    sum.to_string()
}

pub fn part1_streaming(iter: impl Iterator<Item = String>) -> String {
    let mut list1 = Vec::new();
    let mut list2 = Vec::new();
    let mut sum = 0;
    for line in iter {
        if line.is_empty() {
            continue;
        }
        let nums: Vec<i32> = input::parse_words(&line);
        list1.push(nums[0]);
        list2.push(nums[1]);
    }
    assert_eq!(list1.len(), list2.len());
    list1.sort_unstable();
    list2.sort_unstable();
    for i in 0..list1.len() {
        sum += (list1[i] - list2[i]).abs();
    }
    sum.to_string()
}

pub fn part2(lists: (Vec<i32>, Vec<i32>)) -> String {
    let mut similarity = 0;
    for num in lists.0 {
        let appearances = lists.1.iter().filter(|n| **n == num).count();
        similarity += num as usize * appearances;
    }
    similarity.to_string()
}

pub fn part2_streaming(lists: (Vec<i32>, Vec<i32>)) -> String {
    todo!()
}

day_tests!("1882714", "19437052");
