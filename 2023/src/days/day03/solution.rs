use std::collections::HashSet;

use crate::days::{AdventDay, SolutionOutput};

pub struct Day03;

fn find_leftest(grid: &Vec<Vec<char>>, (r, c): (i64, i64)) -> (i64, i64) {
    let ch = grid
        .get(r as usize)
        .and_then(|y| y.get((c as usize).wrapping_sub(1)));
    if ch.is_some() && ch.unwrap().is_numeric() {
        find_leftest(grid, (r, c - 1))
    } else {
        (r, c)
    }
}

impl AdventDay for Day03 {
    fn input_base_path(&self) -> String {
        "src/days/day03".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        let grid: Vec<Vec<char>> = input.split('\n').map(|row| row.chars().collect()).collect();
        let coords_of_symbols: Vec<(i64, i64)> = grid
            .iter()
            .enumerate()
            .flat_map(move |(r, l)| {
                l.iter().enumerate().flat_map(move |(c, ch)| {
                    if ch != &'.' && !ch.is_numeric() {
                        Some((r as i64, c as i64))
                    } else {
                        None
                    }
                })
            })
            .collect();

        let symbol_neighbors: Vec<HashSet<(i64, i64)>> = coords_of_symbols
            .iter()
            .map(|&(x, y)| {
                (-1..=1)
                    .flat_map(move |dx| (-1..=1).map(move |dy| (x + dx, y + dy)))
                    .filter(|(r, c)| {
                        let ch = grid.get(*r as usize).and_then(|y| y.get(*c as usize));
                        ch.is_some() && ch.unwrap().is_numeric()
                    })
                    .map(|(r, c)| find_leftest(&grid, (r, c)))
                    .collect()
            })
            .collect();

        let numbers: Vec<i64> = symbol_neighbors
            .iter()
            .flat_map(|sc| {
                sc.iter().map(|(r, c)| {
                    let num_str: String = grid
                        .get(*r as usize)
                        .unwrap()
                        .iter()
                        .skip(*c as usize)
                        .take_while(|ch| ch.is_numeric())
                        .collect();
                    num_str.parse::<i64>().unwrap()
                })
            })
            .collect();

        SolutionOutput::Int(numbers.iter().sum())
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        let grid: Vec<Vec<char>> = input.split('\n').map(|row| row.chars().collect()).collect();
        let coords_of_symbols: Vec<(i64, i64)> = grid
            .iter()
            .enumerate()
            .flat_map(move |(r, l)| {
                l.iter().enumerate().flat_map(move |(c, ch)| {
                    if ch != &'.' && !ch.is_numeric() {
                        Some((r as i64, c as i64))
                    } else {
                        None
                    }
                })
            })
            .collect();

        let symbol_neighbors: Vec<HashSet<(i64, i64)>> = coords_of_symbols
            .iter()
            .map(|&(x, y)| {
                (-1..=1)
                    .flat_map(move |dx| (-1..=1).map(move |dy| (x + dx, y + dy)))
                    .filter(|(r, c)| {
                        let ch = grid.get(*r as usize).and_then(|y| y.get(*c as usize));
                        ch.is_some() && ch.unwrap().is_numeric()
                    })
                    .map(|(r, c)| find_leftest(&grid, (r, c)))
                    .collect()
            })
            .collect();

        let numbers: Vec<Vec<i64>> = symbol_neighbors
            .iter()
            .map(|sc| {
                sc.iter()
                    .map(|(r, c)| {
                        let num_str: String = grid
                            .get(*r as usize)
                            .unwrap()
                            .iter()
                            .skip(*c as usize)
                            .take_while(|ch| ch.is_numeric())
                            .collect();
                        num_str.parse::<i64>().unwrap()
                    })
                    .collect()
            })
            .collect();

        let gear_ratios = numbers
            .iter()
            .filter(|sc| sc.len() == 2)
            .map(|sc| sc.first().unwrap() * sc.get(1).unwrap());

        SolutionOutput::Int(gear_ratios.sum())
    }
}
