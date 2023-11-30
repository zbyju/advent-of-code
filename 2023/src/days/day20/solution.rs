use crate::days::{AdventDay, SolutionOutput};

pub struct Day20;

impl AdventDay for Day20 {
    fn input_base_path(&self) -> String {
        "src/days/day20".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        println!("Input for Part 1:\n{}", input);
        SolutionOutput::String("Part 1 output displayed".to_string())
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        println!("Input for Part 2:\n{}", input);
        SolutionOutput::String("Part 2 output displayed".to_string())
    }
}
