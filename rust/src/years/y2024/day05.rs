use std::collections::{BTreeMap, BTreeSet};
use std::str::FromStr;

use rustc_hash::{FxHashMap, FxHashSet};

use crate::core::*;
use crate::day_tests;

pub const NAME: &str = "Print Queue";
pub const YEAR: u16 = 2024;
pub const DAY: u16 = 5;

pub type Parsed = PrintQueue;
pub type Output = u32;

#[derive(Clone, Debug)]
pub struct PrintQueue {
    ordering_rules_hash: FxHashMap<u32, FxHashSet<u32>>,
    ordering_rules_btree: BTreeMap<u32, BTreeSet<u32>>,
    ordering_rules_vec: Vec<PageOrderingRule>,
    page_updates: Vec<PageUpdate>,
}

#[derive(Clone, Debug)]
pub struct PageOrderingRule(u32, u32);

impl FromStr for PageOrderingRule {
    type Err = ();
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let (part1, part2) = s.split_once('|').ok_or(())?;
        let page1 = part1.parse().unwrap();
        let page2 = part2.parse().unwrap();
        Ok(PageOrderingRule(page1, page2))
    }
}

#[derive(Clone, Debug)]
pub struct PageUpdate(Vec<u32>);

impl PageUpdate {
    pub fn is_correctly_ordered_hash(&self, rules: &FxHashMap<u32, FxHashSet<u32>>) -> bool {
        let mut correct = true;
        let mut forbidden: Vec<&FxHashSet<u32>> = Vec::with_capacity(self.0.len());
        for page in &self.0 {
            for f in &forbidden {
                if f.contains(page) {
                    correct = false;
                }
            }
            if let Some(forbidden_pages) = rules.get(page) {
                forbidden.push(forbidden_pages);
            }
        }
        correct
    }

    pub fn is_correctly_ordered_btree(&self, rules: &BTreeMap<u32, BTreeSet<u32>>) -> bool {
        let mut correct = true;
        let mut forbidden: Vec<&BTreeSet<u32>> = Vec::with_capacity(self.0.len());
        for page in &self.0 {
            for f in &forbidden {
                if f.contains(page) {
                    correct = false;
                }
            }
            if let Some(forbidden_pages) = rules.get(page) {
                forbidden.push(forbidden_pages);
            }
        }
        correct
    }

    pub fn is_correctly_ordered_vec(&self, rules: &[PageOrderingRule]) -> bool {
        let mut correct = true;
        for rule in rules {
            let pos1 = self.0.iter().position(|p| *p == rule.0);
            let pos2 = self.0.iter().position(|p| *p == rule.1);
            if let (Some(pos1), Some(pos2)) = (pos1, pos2)
                && pos1 > pos2
            {
                correct = false;
            }
        }
        correct
    }

    pub fn middle_page(&self) -> u32 {
        assert!(self.0.len() % 2 == 1, "middle page not well defined");
        self.0[self.0.len() / 2]
    }
}

impl FromStr for PageUpdate {
    type Err = ();
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let pages = input::split_to_numbers(s, ',');
        Ok(PageUpdate(pages))
    }
}

pub fn parse(input: &str) -> Parsed {
    let mut ordering_rules_hash: FxHashMap<u32, FxHashSet<u32>> = FxHashMap::default();
    let mut ordering_rules_btree: BTreeMap<u32, BTreeSet<u32>> = BTreeMap::new();
    let mut ordering_rules_vec: Vec<PageOrderingRule> = Vec::new();
    let mut page_updates: Vec<PageUpdate> = Vec::new();
    let mut reading_updates = false;

    for line in input.lines() {
        if line.is_empty() {
            reading_updates = true;
            continue;
        }
        if !reading_updates {
            let ordering_rule: PageOrderingRule = line.parse().unwrap();
            ordering_rules_hash
                .entry(ordering_rule.1)
                .or_default()
                .insert(ordering_rule.0);
            ordering_rules_btree
                .entry(ordering_rule.1)
                .or_default()
                .insert(ordering_rule.0);
            ordering_rules_vec.push(ordering_rule);
        } else {
            page_updates.push(line.parse().unwrap());
        }
    }

    PrintQueue {
        ordering_rules_hash,
        ordering_rules_btree,
        ordering_rules_vec,
        page_updates,
    }
}

pub fn part1(queue: Parsed) -> Output {
    let mut sum = 0;
    for update in &queue.page_updates {
        if update.is_correctly_ordered_hash(&queue.ordering_rules_hash) {
            sum += update.middle_page();
        }
    }
    sum
}

pub fn part1_btree(queue: Parsed) -> Output {
    let mut sum = 0;
    for update in &queue.page_updates {
        if update.is_correctly_ordered_btree(&queue.ordering_rules_btree) {
            sum += update.middle_page();
        }
    }
    sum
}

pub fn part1_vec(queue: Parsed) -> Output {
    let mut sum = 0;
    for update in &queue.page_updates {
        if update.is_correctly_ordered_vec(&queue.ordering_rules_vec) {
            sum += update.middle_page();
        }
    }
    sum
}

pub fn part1_streaming(queue: Parsed) -> Output {
    0
}

pub fn part1_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![("btree", part1_btree), ("vec", part1_vec)]
}

pub fn part2(mut queue: Parsed) -> Output {
    let mut sum = 0;
    for update in &mut queue.page_updates {
        if update.is_correctly_ordered_hash(&queue.ordering_rules_hash) {
            continue;
        }
        loop {
            let mut correct = true;
            for rule in &queue.ordering_rules_vec {
                let pos1 = update.0.iter().position(|p| *p == rule.0);
                let pos2 = update.0.iter().position(|p| *p == rule.1);
                if let (Some(pos1), Some(pos2)) = (pos1, pos2)
                    && pos1 > pos2
                {
                    update.0.swap(pos1, pos2);
                    correct = false;
                }
            }
            if correct {
                break;
            }
        }
        sum += update.middle_page();
    }
    sum
}

pub fn part2_alternatives() -> Vec<(&'static str, fn(Parsed) -> Output)> {
    vec![]
}

day_tests!("7024", "4151");
