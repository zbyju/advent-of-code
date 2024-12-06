import aoc_solution.{
  type Solution, type SolutionResult, IntResult, Solution, StringResult,
}
import gleam/dict.{type Dict}
import gleam/int
import gleam/list
import gleam/option.{None, Some}
import gleam/result
import gleam/string

fn parse_lines(lines: String) -> Result(#(List(Int), List(Int)), String) {
  let lines = list.filter(string.split(lines, "\n"), fn(s) { s != "" })

  // Initialize empty lists for left and right numbers
  let initial = Ok(#([], []))

  // Fold over each line to parse the numbers
  list.fold(lines, initial, fn(acc, line) {
    let split = list.filter(string.split(line, " "), fn(s) { s != "" })

    case split, acc {
      [a, b], Ok(#(l, r)) ->
        case int.parse(a), int.parse(b) {
          Ok(a), Ok(b) -> Ok(#(list.append(l, [a]), list.append(r, [b])))
          _, _ -> Error("Invalid line format 2")
        }
      _, _ -> {
        Error("Invalid line format3")
      }
    }
  })
}

fn parse_lines2(lines: String) -> Result(#(List(Int), Dict(Int, Int)), String) {
  let lines = list.filter(string.split(lines, "\n"), fn(s) { s != "" })

  // Initialize empty lists for left and right numbers
  let initial = Ok(#([], dict.new()))

  // Fold over each line to parse the numbers
  list.fold(lines, initial, fn(acc, line) {
    let split = list.filter(string.split(line, " "), fn(s) { s != "" })

    case split, acc {
      [a, b], Ok(#(l, r)) ->
        case int.parse(a), int.parse(b) {
          Ok(a), Ok(b) ->
            Ok(#(
              list.append(l, [a]),
              dict.upsert(r, b, fn(x) {
                case x {
                  Some(x) -> x + 1
                  None -> 1
                }
              }),
            ))
          _, _ -> Error("Invalid line format 2")
        }
      _, _ -> {
        Error("Invalid line format3")
      }
    }
  })
}

pub fn day01_solution() -> Solution {
  Solution(
    part1: fn(input: String) -> SolutionResult {
      let parsed = parse_lines(input)

      let res =
        result.try(parsed, fn(parsed) {
          let #(left, right) = parsed
          let sl = list.sort(left, int.compare)
          let sr = list.sort(right, int.compare)

          let combined = list.zip(sl, sr)

          let x =
            list.fold(combined, 0, fn(acc, pair) {
              let #(a, b) = pair
              let c = int.absolute_value(a - b)
              acc + c
            })

          Ok(x)
        })

      case res {
        Ok(res) -> IntResult(res)
        Error(e) -> StringResult(e)
      }
    },
    part2: fn(input: String) -> SolutionResult {
      let parsed = parse_lines2(input)

      let res =
        result.try(parsed, fn(parsed) {
          let #(left, right) = parsed

          let x =
            list.fold(left, 0, fn(acc, l) {
              let count = result.unwrap(dict.get(right, l), 0)
              acc + count * l
            })

          Ok(x)
        })

      case res {
        Ok(res) -> IntResult(res)
        Error(e) -> StringResult(e)
      }
    },
  )
}
