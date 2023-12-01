use crate::days::{AdventDay, SolutionOutput};

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
                let first_num = l.chars().nth(0).and_then(|x| x.to_digit(10));
                let last_num = l.chars().nth_back(0).and_then(|x| x.to_digit(10));
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
        let digits = [
            "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "0", "1", "2",
            "3", "4", "5", "6", "7", "8", "9",
        ];
        let res = input
            .lines()
            .map(|l| {
                let first_num = find_number(l, digits);
                let last_num = find_number_rev(l, digits);
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

fn find_number(line: &str, digits: [&str; 19]) -> Option<usize> {
    let res = digits
        .iter()
        .enumerate()
        .filter_map(|(i, d)| line.find(d).map(|x| (i, x)))
        .min_by_key(|t| t.1)
        .map(|(i, _)| if i <= 8 { i + 1 } else { i - 9 });

    res
}

fn find_number_rev(line: &str, digits: [&str; 19]) -> Option<usize> {
    let res = digits
        .iter()
        .enumerate()
        .filter_map(|(i, d)| line.rfind(d).map(|x| (i, x)))
        .max_by_key(|t| t.1)
        .map(|(i, _)| if i <= 8 { i + 1 } else { i - 9 });

    res
}
