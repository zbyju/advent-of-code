pub mod day00;
pub mod day01;
pub mod day02;
pub mod day03;
pub mod day04;
pub mod day05;
pub mod day06;
pub mod day07;
pub mod day08;
pub mod day09;
pub mod day10;
pub mod day11;
pub mod day12;
pub mod day13;
pub mod day14;
pub mod day15;
pub mod day16;
pub mod day17;
pub mod day18;
pub mod day19;
pub mod day20;
pub mod day21;
pub mod day22;
pub mod day23;
pub mod day24;
pub mod day25;

pub trait AdventDay {
    fn input_base_path(&self) -> String;
    fn part1(&self, input: &str) -> SolutionOutput;
    fn part2(&self, input: &str) -> SolutionOutput;

    fn run(&self, test_case: Option<String>, timeit: bool) {
        use std::time::Instant;

        let base_path = self.input_base_path();
        let input_file = match test_case {
            Some(test_file) => format!("{}/inputs/{}", base_path, test_file),
            None => format!("{}/inputs/input.txt", base_path),
        };

        let input = std::fs::read_to_string(input_file).expect("Failed to read input file");

        let now1 = Instant::now();
        let solution_part1 = self.part1(&input);
        let elapsed1 = now1.elapsed();

        let now2 = Instant::now();
        let solution_part2 = self.part2(&input);
        let elapsed2 = now2.elapsed();

        let output_part1 = match solution_part1 {
            SolutionOutput::Int(value) => value.to_string(),
            SolutionOutput::Float(value) => value.to_string(),
            SolutionOutput::String(value) => value,
        };

        let output_part2 = match solution_part2 {
            SolutionOutput::Int(value) => value.to_string(),
            SolutionOutput::Float(value) => value.to_string(),
            SolutionOutput::String(value) => value,
        };

        println!("Part 1: {}", output_part1);
        println!("Part 2: {}", output_part2);

        if timeit {
            println!("Time to solve part1: {:.2?}", elapsed1);
            println!("Time to solve part2: {:.2?}", elapsed2);
            println!("Time to solve all:   {:.2?}", elapsed1 + elapsed2);
        }
    }
}

pub enum SolutionOutput {
    Int(i64),
    Float(f64),
    String(String),
}
