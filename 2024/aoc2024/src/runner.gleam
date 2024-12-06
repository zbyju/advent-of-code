import aoc_solution.{
  type Solution, FloatResult, IntResult, Solution, StringResult,
}
import gleam/float
import gleam/int
import gleam/list
import gleam/option.{type Option}
import gleam/result
import gleam/string
import simplifile

// Import all day solutions
import solutions/day00.{day00_solution}

// Add more day solutions as they are created

pub type DaySolution {
  DaySolution(day: Int, solution: Solution)
}

pub fn get_solutions() -> List(DaySolution) {
  [DaySolution(0, day00_solution())]
}

pub fn run_solution(
  day: Int,
  part: Int,
  filename: Option(String),
) -> Result(String, String) {
  // Find the corresponding solution
  let solutions = get_solutions()

  let maybe_day_solution = list.find(solutions, fn(ds) { ds.day == day })

  // Validate day and part
  use day_solution <- result.try(
    maybe_day_solution
    |> result.map_error(fn(_) { "Invalid day. Must be between 1 and 24." }),
  )

  use part_fn <- result.try(case part {
    1 -> Ok(day_solution.solution.part1)
    2 -> Ok(day_solution.solution.part2)
    _ -> Error("Invalid part. Must be 1 or 2.")
  })

  // Construct input file path
  let default_filename = "input.txt"
  let filepath =
    string.concat([
      "/inputs/day",
      string.pad_start(int.to_string(day), 2, "0"),
      "/",
      option.unwrap(filename, default_filename),
    ])

  use current_path <- result.try(
    simplifile.current_directory()
    |> result.map_error(fn(_) { "Failed to get current directory." }),
  )
  let full_path = current_path <> filepath

  // Read input file
  use input <- result.try(
    simplifile.read(full_path)
    |> result.map_error(fn(e) { simplifile.describe_error(e) }),
  )

  // Run the solution
  let res = part_fn(input)

  Ok(case res {
    IntResult(i) -> int.to_string(i)
    FloatResult(f) -> float.to_string(f)
    StringResult(s) -> s
  })
}
