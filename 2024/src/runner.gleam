import aoc_solution.{
  type Solution, FloatResult, IntResult, Solution, StringResult,
}
import birl
import birl/duration
import gleam/float
import gleam/int
import gleam/io
import gleam/list
import gleam/option.{type Option}
import gleam/result
import gleam/string
import simplifile

// Import all day solutions
import solutions/day00.{day00_solution}
import solutions/day01.{day01_solution}
import solutions/day02.{day02_solution}
import solutions/day03.{day03_solution}
import solutions/day04.{day04_solution}
import solutions/day05.{day05_solution}
import solutions/day06.{day06_solution}
import solutions/day07.{day07_solution}
import solutions/day08.{day08_solution}
import solutions/day09.{day09_solution}
import solutions/day10.{day10_solution}
import solutions/day11.{day11_solution}

type DaySolution {
  DaySolution(solution: Solution)
}

fn get_solutions() -> List(DaySolution) {
  [
    DaySolution(day00_solution()),
    DaySolution(day01_solution()),
    DaySolution(day02_solution()),
    DaySolution(day03_solution()),
    DaySolution(day04_solution()),
    DaySolution(day05_solution()),
    DaySolution(day06_solution()),
    DaySolution(day07_solution()),
    DaySolution(day08_solution()),
    DaySolution(day09_solution()),
    DaySolution(day10_solution()),
    DaySolution(day11_solution()),
  ]
}

fn get_solution(
  solutions: List(DaySolution),
  day: Int,
) -> Result(DaySolution, String) {
  case solutions, day {
    [], _ ->
      Error(
        "Invalid day. Must be between 0 and "
        <> list.length(solutions) + 1 |> int.to_string
        <> ".",
      )
    [solution, ..], 0 -> Ok(solution)
    [_, ..rest], _ -> get_solution(rest, day - 1)
  }
}

fn get_input(day: Int, filename: String) -> Result(String, String) {
  let filepath =
    string.concat([
      "/inputs/day",
      string.pad_start(int.to_string(day), 2, "0"),
      "/",
      filename,
    ])

  use current_path <- result.try(
    simplifile.current_directory()
    |> result.map_error(fn(_) { "Failed to get current directory." }),
  )
  let full_path = current_path <> filepath

  // Read input file
  simplifile.read(full_path)
  |> result.map_error(fn(e) { simplifile.describe_error(e) })
}

fn time_str(time: Int, divider: Float) -> String {
  let f = int.to_float(time)

  float.divide(f, divider)
  |> result.unwrap(0.0)
  |> float.to_precision(2)
  |> float.to_string
}

fn elapsed_to_str(elapsed: duration.Duration) -> String {
  let time = duration.blur_to(elapsed, duration.MicroSecond)

  let unit_symbol = case time {
    _ if time < 1000 -> "µs"
    _ if time < 1_000_000 -> "ms"
    _ if time < 1_000_000_000 -> "s"
    _ -> "m"
  }

  let time = case unit_symbol {
    "µs" -> time_str(time, 1.0)
    "ms" -> time_str(time, 1000.0)
    "s" -> time_str(time, 1_000_000.0)
    _ -> time_str(time, 60_000_000.0)
  }

  time <> unit_symbol
}

pub fn run_solution(
  day: Int,
  part: Int,
  filename: Option(String),
) -> Result(duration.Duration, String) {
  let solutions = get_solutions()
  let result_day_solution = get_solution(solutions, day)

  use day_solution <- result.try(result_day_solution)

  use part_fn <- result.try(case part {
    1 -> Ok(day_solution.solution.part1)
    2 -> Ok(day_solution.solution.part2)
    _ -> Error("Invalid part. Must be 1 or 2.")
  })

  use input <- result.try(get_input(day, option.unwrap(filename, "input.txt")))

  let start = birl.now()
  let res = part_fn(input)
  let end = birl.now()
  let elapsed = birl.difference(end, start)

  io.println("=====================================")
  io.println(
    "Day "
    <> int.to_string(day)
    <> ", Part "
    <> int.to_string(part)
    <> " (Done in: "
    <> elapsed_to_str(elapsed)
    <> "):",
  )
  case res {
    StringResult(result) -> io.println(result)
    IntResult(result) -> io.println(int.to_string(result))
    FloatResult(result) -> io.println(float.to_string(result))
  }
  io.println("=====================================")

  Ok(elapsed)
}

pub fn run_all_solutions(
  filename: Option(String),
) -> Result(List(duration.Duration), String) {
  let len = list.length(get_solutions())
  list.range(0, int.subtract(len, 1))
  |> list.map(fn(day) {
    [run_solution(day, 1, filename), run_solution(day, 2, filename)]
  })
  |> list.flatten()
  |> result.all()
}
