use std::collections::BTreeSet;

use crate::days::{AdventDay, SolutionOutput};

pub struct Day11;

fn parse(input: &str) -> (Vec<(usize, usize)>, Vec<usize>, Vec<usize>) {
    let lines: Vec<_> = input.lines().collect();
    let ly = lines.len();
    let lx = lines.first().map_or(0, |line| line.len());
    let mut galaxies = vec![];
    let mut empty_rows: BTreeSet<usize> = (0..ly).collect();
    let mut empty_cols: BTreeSet<usize> = (0..lx).collect();
    for (r, row) in lines.iter().enumerate() {
        for (c, ch) in row.chars().enumerate() {
            if ch == '#' {
                galaxies.push((c, r));
                empty_rows.remove(&r);
                empty_cols.remove(&c);
            }
        }
    }
    (
        galaxies,
        empty_rows.into_iter().collect(),
        empty_cols.into_iter().collect(),
    )
}

fn minmax(x1: usize, x2: usize) -> (usize, usize) {
    if x1 < x2 {
        (x1, x2)
    } else {
        (x2, x1)
    }
}

fn count_range(sorted_vec: &[usize], min_val: usize, max_val: usize) -> usize {
    sorted_vec
        .iter()
        .filter(|&&x| x > min_val && x < max_val)
        .count()
}

impl AdventDay for Day11 {
    fn input_base_path(&self) -> String {
        "src/days/day11".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        let (galaxies, empty_rows, empty_cols) = parse(input);
        let mut total = 0;

        for (i, &(c1, r1)) in galaxies.iter().enumerate() {
            for &(c2, r2) in &galaxies[i + 1..] {
                let (min_r, max_r) = minmax(r1, r2);
                let (min_c, max_c) = minmax(c1, c2);
                let d = c1.abs_diff(c2) + r1.abs_diff(r2);

                // Efficiently count empty rows and columns within the range
                let expanded_rows = count_range(&empty_rows, min_r, max_r);
                let expanded_cols = count_range(&empty_cols, min_c, max_c);

                total += d + expanded_rows + expanded_cols;
            }
        }

        SolutionOutput::Int(total as i64)
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        let (galaxies, empty_rows, empty_cols) = parse(input);
        const EXPANSION: usize = 1_000_000 - 1;
        let mut total = 0;

        let mut sorted_empty_rows: Vec<_> = empty_rows.into_iter().collect();
        sorted_empty_rows.sort_unstable();
        let mut sorted_empty_cols: Vec<_> = empty_cols.into_iter().collect();
        sorted_empty_cols.sort_unstable();

        for (i, &(c1, r1)) in galaxies.iter().enumerate() {
            for &(c2, r2) in &galaxies[i + 1..] {
                let (min_r, max_r) = minmax(r1, r2);
                let (min_c, max_c) = minmax(c1, c2);
                let d = c1.abs_diff(c2) + r1.abs_diff(r2);

                // Efficiently count empty rows and columns within the range
                let expanded_rows = count_range(&sorted_empty_rows, min_r, max_r);
                let expanded_cols = count_range(&sorted_empty_cols, min_c, max_c);

                total += d + expanded_rows * EXPANSION + expanded_cols * EXPANSION;
            }
        }

        SolutionOutput::Int(total as i64)
    }
}
