import aoc_solution.{type Solution, type SolutionResult, IntResult, Solution}

pub fn day00_solution() -> Solution {
  Solution(
    part1: fn(_: String) -> SolutionResult { IntResult(42) },
    part2: fn(_: String) -> SolutionResult { IntResult(84) },
  )
}
