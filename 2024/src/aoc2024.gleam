import gleam/io
import gleam/option.{Some}
import gleam/result
import runner

pub fn main() {
  // Run day 0, part 1 with default input

  let result = runner.run_solution(6, 2, Some("input.txt"))

  let str = result.unwrap_both(result)

  // Print the result
  io.println(str)
}
