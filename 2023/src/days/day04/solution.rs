use std::collections::HashMap;

use crate::days::{AdventDay, SolutionOutput};

pub struct Day04;

impl AdventDay for Day04 {
    fn input_base_path(&self) -> String {
        "src/days/day04".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        let cards: Vec<(Vec<i64>, Vec<i64>)> = input
            .lines()
            .map(|l| {
                let split = l.split(':').last().unwrap().trim();
                let mut num_split = split.split('|');

                (
                    num_split
                        .next()
                        .unwrap()
                        .trim()
                        .split(' ')
                        .filter(|x| !x.is_empty())
                        .map(|x| x.parse::<i64>().unwrap())
                        .collect(),
                    num_split
                        .next()
                        .unwrap()
                        .trim()
                        .split(' ')
                        .filter(|x| !x.is_empty())
                        .map(|x| x.parse::<i64>().unwrap())
                        .collect(),
                )
            })
            .collect();

        let points = cards
            .iter()
            .map(|(winning_numbers, numbers)| {
                numbers
                    .iter()
                    .filter(|x| winning_numbers.contains(x))
                    .count()
            })
            .map(|c| if c > 0 { 2_i64.pow(c as u32 - 1) } else { 0 });

        SolutionOutput::Int(points.sum())
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        let cards: Vec<(Vec<i64>, Vec<i64>)> = input
            .lines()
            .map(|l| {
                let split = l.split(':').last().unwrap().trim();
                let mut num_split = split.split('|');

                (
                    num_split
                        .next()
                        .unwrap()
                        .trim()
                        .split(' ')
                        .filter(|x| !x.is_empty())
                        .map(|x| x.parse::<i64>().unwrap())
                        .collect(),
                    num_split
                        .next()
                        .unwrap()
                        .trim()
                        .split(' ')
                        .filter(|x| !x.is_empty())
                        .map(|x| x.parse::<i64>().unwrap())
                        .collect(),
                )
            })
            .collect();

        let copies = cards
            .iter()
            .map(|(winning_numbers, numbers)| {
                numbers
                    .iter()
                    .filter(|x| winning_numbers.contains(x))
                    .count()
            })
            .enumerate()
            .fold(HashMap::new(), |mut acc, (i, cnt)| {
                for n in i + 1..=i + cnt {
                    acc.insert(n, acc.get(&n).unwrap_or(&1) + acc.get(&i).unwrap_or(&1));
                }
                acc
            });

        SolutionOutput::Int((0..cards.len()).map(|i| copies.get(&i).unwrap_or(&1)).sum())
    }
}
