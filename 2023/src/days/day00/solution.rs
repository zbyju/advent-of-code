use crate::days::{AdventDay, SolutionOutput};

pub struct Day00;

impl AdventDay for Day00 {
    fn input_base_path(&self) -> String {
        "src/days/day00".to_string()
    }

    fn part1(&self, _: &str) -> SolutionOutput {
        SolutionOutput::String("Hello, Advent of Code!".to_string())
    }

    fn part2(&self, _: &str) -> SolutionOutput {
        SolutionOutput::Float(2.0)
    }
}
