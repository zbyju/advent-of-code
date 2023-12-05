extern crate regex;
use crate::days::{AdventDay, SolutionOutput};
use regex::Regex;

pub struct Day01;

impl AdventDay for Day01 {
    fn input_base_path(&self) -> String {
        "src/days/day01".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        let res = input
            .chars()
            .filter(|x| x.is_numeric() || x.is_control())
            .collect::<String>()
            .lines()
            .map(|l| {
                let first_num = l.chars().next().and_then(|x| x.to_digit(10));
                let last_num = l.chars().next_back().and_then(|x| x.to_digit(10));
                match (first_num, last_num) {
                    (Some(a), Some(b)) => Some(a * 10 + b),
                    _ => None,
                }
            })
            .filter_map(|x| x.map(|num| num as i64))
            .sum();

        SolutionOutput::Int(res)
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        let input = replace_numbers(input);

        let res = input
            .chars()
            .filter(|x| x.is_numeric() || x.is_control())
            .collect::<String>()
            .lines()
            .map(|l| {
                let first_num = l.chars().next().and_then(|x| x.to_digit(10));
                let last_num = l.chars().next_back().and_then(|x| x.to_digit(10));
                match (first_num, last_num) {
                    (Some(a), Some(b)) => Some(a * 10 + b),
                    _ => None,
                }
            })
            .filter_map(|x| x.map(|num| num as i64))
            .sum();

        SolutionOutput::Int(res)
    }
}

fn replace_numbers(input: &str) -> String {
    let replacements = [
        ("one", "o1e"),
        ("two", "t2o"),
        ("three", "t3e"),
        ("four", "f4r"),
        ("five", "f5e"),
        ("six", "s6x"),
        ("seven", "s7n"),
        ("eight", "e8t"),
        ("nine", "n9e"),
    ];

    let mut result = String::from(input);

    for &(pattern, replacement) in &replacements {
        let re = Regex::new(pattern).unwrap();
        result = re.replace_all(&result, replacement).into_owned();
    }
    result
}
