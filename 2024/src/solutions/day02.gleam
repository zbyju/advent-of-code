import aoc_solution.{
  type Solution, type SolutionResult, IntResult, Solution, StringResult,
}
import gleam/int
import gleam/list
import gleam/result
import gleam/string

type Order {
  Increasing
  Decreasing
  Constant
  Unknown
}

fn is_safe_adjacent(
  prev: Int,
  curr: Int,
  previous_order: Order,
) -> #(Bool, Order) {
  let diff = curr - prev
  let current_order = case diff {
    _ if diff == 0 -> Constant
    _ if diff > 0 -> Increasing
    _ -> Decreasing
  }
  let abs_diff = int.absolute_value(diff)
  case previous_order, current_order {
    _, _ if abs_diff < 1 || abs_diff > 3 -> #(False, Unknown)
    Constant, _ -> #(False, Constant)
    _, Constant -> #(False, Constant)
    Increasing, Decreasing -> #(False, Increasing)
    Decreasing, Increasing -> #(False, Decreasing)

    Unknown, _ -> #(True, current_order)
    _, _ -> #(True, current_order)
  }
}

fn is_safe(nums: List(Int), order: Order, acc: Bool) -> Bool {
  case nums {
    [] -> acc
    [_] -> acc
    [x, y, ..rest] -> {
      let is_safe_xy = is_safe_adjacent(x, y, order)
      case is_safe_xy.0 {
        False -> False
        True -> is_safe([y, ..rest], is_safe_xy.1, acc)
      }
    }
  }
}

pub fn day02_solution() -> Solution {
  Solution(
    part1: fn(input: String) -> SolutionResult {
      let lines =
        string.split(input, "\n")
        |> list.filter(fn(s) { s != "" })

      let safe_lines =
        list.map(lines, fn(line) {
          let split = string.split(line, " ")
          let nums =
            list.map(split, fn(s) { int.parse(s) })
            |> result.all()

          result.try(nums, fn(nums) { Ok(is_safe(nums, Unknown, True)) })
        })
        |> result.all()

      case safe_lines {
        Ok(safe_lines) ->
          IntResult(list.length(list.filter(safe_lines, fn(x) { x })))
        Error(_) -> StringResult("Error parsing input")
      }
    },
    part2: fn(input: String) -> SolutionResult {
      let lines =
        string.split(input, "\n")
        |> list.filter(fn(s) { s != "" })

      let safe_lines =
        list.map(lines, fn(line) {
          let split = string.split(line, " ")
          let nums =
            list.map(split, fn(s) { int.parse(s) })
            |> result.all()

          result.try(nums, fn(nums) {
            let combinations = list.combinations(nums, list.length(nums) - 1)

            let safe =
              list.fold(combinations, False, fn(acc, combination) {
                case acc {
                  True -> True
                  False -> is_safe(combination, Unknown, True)
                }
              })

            Ok(safe)
          })
        })
        |> result.all()

      case safe_lines {
        Ok(safe_lines) ->
          IntResult(list.length(list.filter(safe_lines, fn(x) { x })))
        Error(_) -> StringResult("Error parsing input")
      }
    },
  )
}
