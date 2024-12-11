import aoc_solution.{type Solution, type SolutionResult, IntResult, Solution}
import gleam/int
import gleam/list
import gleam/option.{None, Some}
import gleam/regex
import gleam/result

type MulExpression {
  MulExpression(first: Int, second: Int)
}

fn parse_mul_expressions(input: String) -> List(MulExpression) {
  let assert Ok(pattern) = regex.from_string("mul\\((\\d+),(\\d+)\\)")

  pattern
  |> regex.scan(input)
  |> list.map(fn(match) {
    case match.submatches {
      [Some(first), Some(second)] ->
        Some(MulExpression(
          result.unwrap(int.parse(first), 0),
          result.unwrap(int.parse(second), 0),
        ))
      _ -> None
    }
  })
  |> option.values()
}

type Expression {
  MulExpression2(first: Int, second: Int)
  DoOperation
  DontOperation
}

fn parse_expressions(input: String) -> List(Expression) {
  let pattern = regex.from_string("mul\\((\\d+),(\\d+)\\)|do\\(\\)|don't\\(\\)")

  case pattern {
    Ok(p) ->
      p
      |> regex.scan(input)
      |> list.map(fn(match) {
        case match.content {
          "do()" -> DoOperation
          "don't()" -> DontOperation
          mul_match -> {
            case match.submatches {
              [Some(first), Some(second)] ->
                MulExpression2(
                  result.unwrap(int.parse(first), 0),
                  result.unwrap(int.parse(second), 0),
                )
              _ -> panic as "Invalid mul expression"
            }
          }
        }
      })
    Error(_) -> []
  }
}

fn run2(expressions: List(Expression), do: Bool, acc: Int) -> Int {
  case expressions {
    [] -> acc
    [e, ..rest] ->
      case e {
        DoOperation -> run2(rest, True, acc)
        DontOperation -> run2(rest, False, acc)
        MulExpression2(a, b) ->
          case do {
            True -> run2(rest, do, acc + a * b)
            False -> run2(rest, do, acc)
          }
      }
  }
}

pub fn day03_solution() -> Solution {
  Solution(
    part1: fn(input: String) -> SolutionResult {
      let mul_expressions = parse_mul_expressions(input)

      let res =
        list.fold(mul_expressions, 0, fn(acc, mul) {
          acc + mul.first * mul.second
        })

      IntResult(res)
    },
    part2: fn(input: String) -> SolutionResult {
      let expressions = parse_expressions(input)

      let res = run2(expressions, True, 0)

      IntResult(res)
    },
  )
}
