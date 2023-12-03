use std::collections::HashSet;

use crate::days::{AdventDay, SolutionOutput};

pub struct Day03;

impl AdventDay for Day03 {
    fn input_base_path(&self) -> String {
        "src/days/day03".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        let grid: Vec<Vec<char>> = input.split('\n').map(|row| row.chars().collect()).collect();
        let mut cs = HashSet::new();

        for (r, row) in grid.iter().enumerate() {
            for (c, ch) in row.iter().enumerate() {
                if ch.is_numeric() || *ch == '.' {
                    continue;
                }
                for dr in [r - 1, r, r + 1] {
                    for mut dc in [c - 1, c, c + 1] {
                        let target_ch = grid.get(dr).and_then(|row| row.get(dc)).unwrap_or(&'.');
                        if !target_ch.is_numeric() {
                            continue;
                        }
                        while grid[dr]
                            .iter()
                            .nth(dc.wrapping_sub(1))
                            .unwrap_or(&'.')
                            .is_numeric()
                        {
                            dc -= 1;
                        }
                        cs.insert((dr, dc));
                    }
                }
            }
        }

        let mut ns = Vec::new();
        for (nr, mut nc) in cs {
            let mut s = String::new();
            while nc < grid[nr].len() && grid[nr][nc].is_numeric() {
                s.push(grid[nr][nc]);
                nc += 1;
            }
            ns.push(s.parse::<i64>().unwrap());
        }

        let sum = ns.iter().sum();

        SolutionOutput::Int(sum)
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        let grid: Vec<Vec<char>> = input.split('\n').map(|row| row.chars().collect()).collect();
        let mut total: i64 = 0;

        for (r, row) in grid.iter().enumerate() {
            for (c, ch) in row.iter().enumerate() {
                if ch.is_numeric() || *ch == '.' {
                    continue;
                }
                let mut cs = HashSet::new();
                for dr in [r - 1, r, r + 1] {
                    for mut dc in [c - 1, c, c + 1] {
                        let target_ch = grid.get(dr).and_then(|row| row.get(dc)).unwrap_or(&'.');
                        if !target_ch.is_numeric() {
                            continue;
                        }
                        while grid[dr]
                            .iter()
                            .nth(dc.wrapping_sub(1))
                            .unwrap_or(&'.')
                            .is_numeric()
                        {
                            dc -= 1;
                        }

                        cs.insert((dr, dc));
                    }
                }
                if cs.len() != 2 {
                    continue;
                }
                let mut ns = Vec::new();
                for (nr, mut nc) in cs {
                    let mut s = String::new();
                    while nc < grid[nr].len() && grid[nr][nc].is_numeric() {
                        s.push(grid[nr][nc]);
                        nc += 1;
                    }
                    ns.push(s.parse::<i64>().unwrap());
                }
                total += ns[0] * ns[1];
            }
        }

        SolutionOutput::Int(total)
    }
}
