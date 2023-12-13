use crate::days::{AdventDay, SolutionOutput};

pub struct Day13;

type Grid = Vec<Vec<char>>;

fn count_differences(vec1: &[char], vec2: &[char]) -> usize {
    vec1.iter()
        .zip(vec2.iter())
        .filter(|&(a, b)| a != b)
        .count()
}

fn is_reflection_point(grid: &Grid, ri: usize, max_changes: isize) -> bool {
    let mut max_changes = max_changes;
    for offset in 0..=ri {
        let i = ri - offset;
        let j = ri + offset + 1;

        if j >= grid.len() {
            return max_changes == 0;
        }

        max_changes -= count_differences(&grid[i], &grid[j]) as isize;

        if max_changes < 0 {
            return false;
        }
    }
    max_changes == 0
}

fn find_reflection_point(grid: &Grid, max_changes: isize) -> Option<usize> {
    let row_len = grid.len();

    (0..row_len - 1).find(|&i| is_reflection_point(grid, i, max_changes))
}

fn transpose(matrix: &Vec<Vec<char>>) -> Vec<Vec<char>> {
    let row_len = matrix.len();
    let col_len = matrix[0].len();

    let mut transposed = vec![vec![' '; row_len]; col_len];

    (0..row_len).for_each(|i| {
        (0..col_len).for_each(|j| {
            transposed[j][i] = matrix[i][j];
        });
    });

    transposed
}

impl AdventDay for Day13 {
    fn input_base_path(&self) -> String {
        "src/days/day13".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        let grids: Vec<Grid> = input
            .split("\n\n")
            .map(|b| b.lines().map(|l| l.chars().collect()).collect())
            .collect();

        let points: usize = grids
            .iter()
            .map(|grid| {
                if let Some(row_reflection) = find_reflection_point(grid, 0) {
                    return 100 * (row_reflection + 1);
                }
                let transposed = transpose(grid);
                if let Some(col_reflection) = find_reflection_point(&transposed, 0) {
                    return col_reflection + 1;
                }
                0
            })
            .sum();

        SolutionOutput::Int(points as i64)
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        let grids: Vec<Grid> = input
            .split("\n\n")
            .map(|b| b.lines().map(|l| l.chars().collect()).collect())
            .collect();

        let points: usize = grids
            .iter()
            .map(|grid| {
                if let Some(row_reflection) = find_reflection_point(grid, 1) {
                    return 100 * (row_reflection + 1);
                }
                let transposed = transpose(grid);
                if let Some(col_reflection) = find_reflection_point(&transposed, 1) {
                    return col_reflection + 1;
                }
                0
            })
            .sum();

        SolutionOutput::Int(points as i64)
    }
}
