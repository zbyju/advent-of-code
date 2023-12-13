use itertools::Itertools;
use std::collections::HashMap;

use crate::days::{AdventDay, SolutionOutput};

pub struct Day12;

#[derive(Hash, Eq, PartialEq, Debug, Clone)]
struct Key(String, Vec<usize>);

fn count(cfgStr: &String, nums: &[usize], cache: &mut HashMap<Key, i64>) -> i64 {
    let cfg: Vec<char> = cfgStr.chars().collect();
    if cfg.is_empty() {
        return if nums.is_empty() { 1 } else { 0 };
    }

    if nums.is_empty() {
        return if cfg.contains(&'#') { 0 } else { 1 };
    }

    let key = Key(cfg.iter().collect(), nums.to_vec());

    if let Some(&cached_result) = cache.get(&key) {
        return cached_result;
    }

    let mut result = 0;

    if cfg[0] == '.' || cfg[0] == '?' {
        result += count(&cfg[1..].iter().collect(), nums, cache);
    }

    if cfg[0] == '#' || cfg[0] == '?' {
        if nums[0] <= cfg.len()
            && !cfg[..nums[0]].contains(&'.')
            && (nums[0] == cfg.len() || cfg[nums[0]] != '#')
        {
            let i = nums[0] + 1;
            result += count(
                &cfg.get(i..).unwrap_or(&Vec::new()).iter().collect(),
                &nums[1..],
                cache,
            );
        }
    }

    cache.insert(key, result);
    result
}

fn parse_line(line: &String, repetitions: usize) -> (Vec<char>, Vec<usize>) {
    let (schematic_str, groups_str) = line.split_whitespace().collect_tuple().unwrap();
    let schematic = schematic_str.chars().collect_vec();
    let groups = groups_str
        .split(",")
        .map(|s| s.parse::<usize>().unwrap())
        .collect_vec();

    let repeated_schematic = (0..repetitions)
        .into_iter()
        .flat_map(|i| {
            schematic
                .clone()
                .into_iter()
                .chain(if i < repetitions - 1 { Some('?') } else { None })
        })
        .collect_vec();
    let repeated_groups = groups.repeat(repetitions);
    (repeated_schematic, repeated_groups)
}

impl AdventDay for Day12 {
    fn input_base_path(&self) -> String {
        "src/days/day12".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        let mut total: i64 = 0;

        for line in input.lines() {
            let parts: Vec<&str> = line.split_whitespace().collect();

            if parts.len() != 2 {
                continue;
            }

            let (cfg, nums_str) = (parts[0], parts[1]);
            let nums: Vec<usize> = nums_str.split(',').filter_map(|s| s.parse().ok()).collect();

            total += count(&cfg.chars().collect(), &nums, &mut HashMap::new());
        }
        SolutionOutput::Int(total)
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        let mut cache = HashMap::new();
        let mut total = 0;

        for line in input.lines() {
            let (cfg, nums) = parse_line(&line.to_string(), 5);

            total += count(&cfg.iter().collect(), &nums, &mut cache);
        }
        SolutionOutput::Int(total)
    }
}
