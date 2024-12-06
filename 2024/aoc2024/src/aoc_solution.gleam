pub type Solution {
  Solution(
    part1: fn(String) -> SolutionResult,
    part2: fn(String) -> SolutionResult,
  )
}

pub type SolutionResult {
  IntResult(Int)
  FloatResult(Float)
  StringResult(String)
}
