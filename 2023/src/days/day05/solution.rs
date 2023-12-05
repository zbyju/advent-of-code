use crate::days::{AdventDay, SolutionOutput};

pub struct Day05;

impl AdventDay for Day05 {
    fn input_base_path(&self) -> String {
        "src/days/day05".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        let blocks: Vec<&str> = input.split("\n\n").collect();

        let seeds: Vec<i64> = blocks
            .iter()
            .nth(0)
            .unwrap()
            .split(":")
            .nth(1)
            .unwrap()
            .split_whitespace()
            .map(|x| x.parse().unwrap())
            .collect();

        let block_ranges: Vec<Vec<(i64, i64, i64)>> = blocks
            .iter()
            .skip(1)
            .map(|b| {
                b.lines()
                    .skip(1)
                    .map(|l| {
                        let range: Vec<i64> =
                            l.split_whitespace().map(|x| x.parse().unwrap()).collect();
                        (range[0], range[1], range[2])
                    })
                    .collect()
            })
            .collect();

        let res: Vec<i64> = seeds
            .iter()
            .map(|seed| {
                block_ranges.iter().fold(*seed, |acc, ranges| {
                    let range = ranges
                        .iter()
                        .filter(|(_, b, c)| b <= &acc && &acc < &(b + c))
                        .nth(0);
                    match range {
                        None => acc,
                        Some((a, b, _)) => acc - b + a,
                    }
                })
            })
            .collect();

        SolutionOutput::Int(*res.iter().min().unwrap())
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        let blocks: Vec<&str> = input.split("\n\n").collect();

        let numbers: Vec<i64> = blocks
            .iter()
            .nth(0)
            .unwrap()
            .split(":")
            .nth(1)
            .unwrap()
            .split_whitespace()
            .map(|x| x.parse().unwrap())
            .collect();

        let mut seeds: Vec<(i64, i64)> = numbers
            .chunks_exact(2)
            .map(|chunk| (chunk[0], chunk[0] + chunk[1]))
            .collect();

        let block_ranges: Vec<Vec<(i64, i64, i64)>> = blocks
            .iter()
            .skip(1)
            .map(|b| {
                b.lines()
                    .skip(1)
                    .map(|l| {
                        let range: Vec<i64> =
                            l.split_whitespace().map(|x| x.parse().unwrap()).collect();
                        (range[0], range[1], range[2])
                    })
                    .collect()
            })
            .collect();

        for ranges in block_ranges {
            let mut new = Vec::new();
            while let Some((s, e)) = seeds.pop() {
                let mut added = false;
                for range in &ranges {
                    let (a, b, c) = (range.0, range.1, range.2);
                    let os = s.max(b);
                    let oe = e.min(b + c);
                    if os < oe {
                        added = true;
                        new.push((os - b + a, oe - b + a));
                        if os > s {
                            seeds.push((s, os));
                        }
                        if e > oe {
                            seeds.push((oe, e));
                        }
                        break;
                    }
                }
                if !added {
                    new.push((s, e));
                }
            }
            seeds = new;
        }

        SolutionOutput::Int(seeds.iter().map(|x| x.0).min().unwrap())
    }
}
