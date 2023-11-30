#!/bin/bash

# Change this to the number of days you want to create
NUM_DAYS=25

# Base directory for day modules
BASE_DIR="src/days"

# Create the base directory if it doesn't exist
mkdir -p $BASE_DIR

# Create mod.rs file
touch $BASE_DIR/mod.rs

# Function to create a day's directory, solution file, and input files
create_day() {
    local day=$1
    local day_formatted=$(printf "%02d" $day) # 1 -> 01, 2 -> 02, 10 -> 10, 15 -> 15
    local full_day_formatted=$(printf "day%02d" $day) 
    local day_dir="$BASE_DIR/$full_day_formatted"
    local inputs_dir="$day_dir/inputs"

    # Create day directory and inputs directory
    mkdir -p $day_dir $inputs_dir

    # Create solution.rs file
    cat > "$day_dir/solution.rs" << EOF
use crate::days::{AdventDay, SolutionOutput};

pub struct Day$day_formatted;

impl AdventDay for Day$day_formatted {
    fn input_base_path(&self) -> String {
        "src/days/day$day_formatted".to_string()
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
EOF

    # Create mod.rs
    echo "pub mod solution;" > "$day_dir/mod.rs" 

    # Create input.txt and test1.txt
    touch "$inputs_dir/input.txt"
    touch "$inputs_dir/test1.txt"

    # Append pub mod dayXX to mod.rs
    echo "pub mod $full_day_formatted;" >> $BASE_DIR/mod.rs
}

# Loop to create each day
for (( day=0; day<=NUM_DAYS; day++ ))
do
    echo "Creating $day structure..."
    create_day $day
done

echo "pub trait AdventDay {
    fn input_base_path(&self) -> String;
    fn part1(&self, input: &str) -> SolutionOutput;
    fn part2(&self, input: &str) -> SolutionOutput;

    fn run(&self, test_case: Option<String>) {
        let base_path = self.input_base_path();
        let input_file = match test_case {
            Some(test_file) => format!(\"{}/inputs/{}\", base_path, test_file),
            None => format!(\"{}/inputs/input.txt\", base_path),
        };

        let input = std::fs::read_to_string(input_file)
            .expect(\"Failed to read input file\");

        let output_part1 = match self.part1(&input) {
            SolutionOutput::Int(value) => value.to_string(),
            SolutionOutput::Float(value) => value.to_string(),
            SolutionOutput::String(value) => value,
        };

        let output_part2 = match self.part2(&input) {
            SolutionOutput::Int(value) => value.to_string(),
            SolutionOutput::Float(value) => value.to_string(),
            SolutionOutput::String(value) => value,
        };

        println!(\"Part 1: {}\", output_part1);
        println!(\"Part 2: {}\", output_part2);
    }
}

pub enum SolutionOutput {
    Int(i64),
    Float(f64),
    String(String),
}" >> $BASE_DIR/mod.rs

echo "Day setup complete."

