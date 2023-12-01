use crate::days::{AdventDay, SolutionOutput};

pub struct Day01;

impl AdventDay for Day01 {
    fn input_base_path(&self) -> String {
        "src/days/day01".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        let lines = input.lines();
        let res = lines
            .map(|l| {
                let r = l.chars().rev().collect::<String>();
                let first_num_index = l.find(char::is_numeric);
                let last_num_index = r.find(char::is_numeric);
                let first_num = first_num_index
                    .and_then(|x| l.chars().nth(x))
                    .map(|y| y.to_string());
                let last_num = last_num_index
                    .and_then(|x| r.chars().nth(x))
                    .map(|y| y.to_string());
                format!(
                    "{}{}",
                    first_num.unwrap_or("0".to_string()),
                    last_num.unwrap_or("0".to_string())
                )
            })
            .fold(0, |acc, x| acc + x.parse::<i64>().unwrap());

        SolutionOutput::Int(res)
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        let lines = input.lines();
        let res = lines
            .map(|l| {
                let first_num = find_number(l, false).unwrap();
                let last_num = find_number(l, true).unwrap();
                format!("{}{}", first_num, last_num)
            })
            .fold(0, |acc, x| acc + x.parse::<i64>().unwrap());

        SolutionOutput::Int(res)
    }
}
fn find_number(line: &str, reverse: bool) -> Option<i32> {
    let digits = [
        "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "0", "1", "2", "3",
        "4", "5", "6", "7", "8", "9",
    ];

    let res = digits
        .map(|d| (if reverse { line.rfind(d) } else { line.find(d) }, d))
        .iter()
        .filter(|t| t.0.is_some())
        .map(|t| (t.0.unwrap(), t.1))
        .reduce(|acc, t| {
            if reverse {
                if t.0 > acc.0 {
                    t
                } else {
                    acc
                }
            } else if t.0 < acc.0 {
                t
            } else {
                acc
            }
        })
        .map(|x| match x.1 {
            "one" | "1" => 1,
            "two" | "2" => 2,
            "three" | "3" => 3,
            "four" | "4" => 4,
            "five" | "5" => 5,
            "six" | "6" => 6,
            "seven" | "7" => 7,
            "eight" | "8" => 8,
            "nine" | "9" => 9,
            _ => 0,
        });

    res
}
