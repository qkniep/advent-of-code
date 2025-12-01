use aoc::core::*;

use crate::day_tests;

pub const NAME: &str = "Red-Nosed Reports";
pub const YEAR: u16 = 2024;
pub const DAY: u16 = 2;

pub type Parsed = Vec<Report>;
pub type Report = Vec<i32>;

pub fn parse(input: &str) -> Parsed {
    let mut reports = Vec::with_capacity(input.lines().count());
    for line in input.lines() {
        if line.is_empty() {
            continue;
        }
        let levels: Vec<i32> = input::parse_words(line);
        reports.push(levels);
    }
    reports
}

pub fn parse_streaming(line: &str) -> Report {
    input::parse_words(line)
}

pub fn part1(reports: Parsed) -> String {
    count_safe_reports(reports, report_is_safe).to_string()
}

pub fn part1_alternatives() -> Vec<(&'static str, fn(Parsed) -> String)> {
    vec![]
}

pub fn part1_streaming(reports: impl Iterator<Item = Report>) -> String {
    count_safe_reports(reports, report_is_safe).to_string()
}

pub fn part2(reports: Parsed) -> String {
    part2_no_clone(reports)
}

pub fn part2_alternatives() -> Vec<(&'static str, fn(Parsed) -> String)> {
    vec![("naive", part2_naive)]
}

pub fn part2_naive(reports: Parsed) -> String {
    count_safe_reports(reports, report_is_safe_tolerance_naive).to_string()
}

pub fn part2_no_clone(reports: Parsed) -> String {
    count_safe_reports(reports, report_is_safe_tolerance_no_clone).to_string()
}

/// Counts how many reports are safe, according to the given `is_safe` predicate.
fn count_safe_reports(
    reports: impl IntoIterator<Item = Report>,
    is_safe: fn(&Report) -> bool,
) -> i32 {
    let mut count_safe = 0;
    for report in reports {
        if is_safe(&report) {
            count_safe += 1;
        }
    }
    count_safe
}

/// Checks whether the report is safe.
///
/// Safety checks:
/// * Levels are strictly increasing or strictly decreasing.
/// * No two consecutive levels differ by more than 3.
fn report_is_safe(report: &Report) -> bool {
    let increasing = report[1] > report[0];
    for window in report.as_slice().windows(2) {
        if (increasing && window[1] < window[0]) || (!increasing && window[1] > window[0]) {
            return false;
        }
        if (window[0] - window[1]).abs() == 0 || (window[0] - window[1]).abs() > 3 {
            return false;
        }
    }
    true
}

/// Checks whether the report is safe, when ignoring the level at `index`.
///
/// Safety checks are the same as `report_is_safe`.
fn report_is_safe_without(report: &Report, index: usize) -> bool {
    let mut left = 0;
    let mut right = 1;
    if left == index {
        left += 1;
    }
    if right == index || right == left {
        right += 1;
    }
    let increasing = report[right] > report[left];
    while right < report.len() {
        if left == index {
            left += 1;
        }
        if right == index || right == left {
            right += 1;
        }
        if right == report.len() {
            break;
        }
        if (increasing && report[right] < report[left])
            || (!increasing && report[right] > report[left])
            || report[right] == report[left]
            || (report[right] - report[left]).abs() > 3
        {
            return false;
        }
        left += 1;
        right += 1;
    }
    true
}

/// Checks whether the report is safe if we allow removal of any level.
fn report_is_safe_tolerance_naive(report: &Report) -> bool {
    if report_is_safe(report) {
        return true;
    }
    for i in 0..report.len() {
        let mut report = report.clone();
        report.remove(i);
        if report_is_safe(&report) {
            return true;
        }
    }
    false
}

/// Checks whether the report is safe if we allow removal of any level.
fn report_is_safe_tolerance_no_clone(report: &Report) -> bool {
    if report_is_safe(report) {
        return true;
    }
    for i in 0..report.len() {
        if report_is_safe_without(report, i) {
            return true;
        }
    }
    false
}

day_tests!("463", "514");
