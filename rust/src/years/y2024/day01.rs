//! Historian Hysteria

use aoc::core::*;

pub fn parse(input: &str) -> (Vec<i32>, Vec<i32>) {
    let mut list1 = Vec::new();
    let mut list2 = Vec::new();
    for line in input.lines() {
        if line.is_empty() {
            continue;
        }
        let mut parts = line.split_ascii_whitespace();
        let num1 = parts.next().unwrap().parse::<i32>().unwrap();
        list1.push(num1);
        let num2 = parts.next().unwrap().parse::<i32>().unwrap();
        list2.push(num2);
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
        let mut parts = line.split_ascii_whitespace();
        let num1 = parts.next().unwrap().parse::<i32>().unwrap();
        list1.push(num1);
        let num2 = parts.next().unwrap().parse::<i32>().unwrap();
        list2.push(num2);
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
