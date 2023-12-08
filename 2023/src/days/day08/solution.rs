use std::collections::HashMap;

use crate::days::{AdventDay, SolutionOutput};

pub struct Day08;

fn gcd(mut n: i64, mut m: i64) -> i64 {
    while m != 0 {
        if m < n {
            std::mem::swap(&mut m, &mut n);
        }
        m %= n;
    }
    n
}

fn parse(input: &str) -> (Vec<char>, HashMap<String, (String, String)>) {
    let mut split = input.split("\n\n");
    let instructions = split.nth(0).unwrap().chars().collect();

    let map = split
        .last()
        .unwrap()
        .lines()
        .fold(HashMap::new(), |mut acc, l| {
            let mut s1 = l.split(" = ");
            let id = s1.next().unwrap().to_string();
            let mut s2 = s1.next().unwrap().split(", ");
            let left = s2.next().unwrap().chars().skip(1).collect();
            let mut right: String = s2.next().unwrap().chars().collect();
            right.pop();
            acc.insert(id, (left, right));
            acc
        });

    (instructions, map)
}

fn parse2(input: &str) -> (Vec<char>, HashMap<String, (String, String)>, Vec<String>) {
    let mut split = input.split("\n\n");
    let instructions = split.nth(0).unwrap().chars().collect();

    let (map, start_positions) =
        split
            .last()
            .unwrap()
            .lines()
            .fold((HashMap::new(), Vec::new()), |mut acc, l| {
                let mut s1 = l.split(" = ");
                let id = s1.next().unwrap().to_string();
                let mut s2 = s1.next().unwrap().split(", ");
                let left = s2.next().unwrap().chars().skip(1).collect();
                let mut right: String = s2.next().unwrap().chars().collect();
                right.pop();
                if id.ends_with('A') {
                    acc.1.push(id.clone());
                }
                acc.0.insert(id, (left, right));
                acc
            });

    (instructions, map, start_positions)
}

impl AdventDay for Day08 {
    fn input_base_path(&self) -> String {
        "src/days/day08".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        let (instructions, map) = parse(input);
        let instructions_len = instructions.len();
        let mut current_node = &"AAA".to_string();
        let mut steps = 0 as usize;

        while current_node != "ZZZ" {
            let node = map.get(current_node).unwrap();
            let index = steps % instructions_len;
            let instruction = instructions.get(index).unwrap();
            match instruction {
                'L' => current_node = &node.0,
                'R' => current_node = &node.1,
                _ => unreachable!(),
            }
            steps += 1
        }

        SolutionOutput::Int(steps as i64)
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        let (instructions, map, start_positions) = parse2(input);
        let instructions_len = instructions.len();

        let cycle_lengths: Vec<i64> = start_positions
            .iter()
            .map(|pos| {
                let mut steps = 0 as usize;
                let mut current_node = pos;

                while !current_node.ends_with('Z') {
                    let node = map.get(current_node).unwrap();
                    let index = steps % instructions_len;
                    let instruction = instructions.get(index).unwrap();
                    match instruction {
                        'L' => current_node = &node.0,
                        'R' => current_node = &node.1,
                        _ => unreachable!(),
                    }
                    steps += 1;
                }
                steps as i64
            })
            .collect();

        let res = cycle_lengths
            .into_iter()
            .fold(1, |acc, x| acc * x / gcd(acc, x));

        SolutionOutput::Int(res)
    }
}
