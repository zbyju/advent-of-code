import aoc_solution.{type Solution, type SolutionResult, IntResult, Solution}
import gleam/dict.{type Dict}
import gleam/int
import gleam/list
import gleam/option.{type Option, None, Some}
import gleam/order
import gleam/result
import gleam/set.{type Set}
import gleam/string

fn parse_input(input: String) -> #(Dict(Int, Set(Int)), List(List(Int))) {
  let parts = string.split(input, "\n\n")

  case parts {
    [first, second] -> {
      let rules =
        string.split(first, "\n") |> list.filter(fn(line) { line != "" })

      let print_rules =
        list.fold(rules, dict.new(), fn(acc, rule) {
          let parts = string.split(rule, "|")
          case parts {
            [before, after] -> {
              let b = result.unwrap(int.parse(before), 0)
              let a = result.unwrap(int.parse(after), 0)
              dict.upsert(acc, b, fn(ma) {
                case ma {
                  None -> set.from_list([a])
                  Some(s) -> set.insert(s, a)
                }
              })
            }
            _ -> panic as "Invalid print rule"
          }
        })

      let numbers =
        string.split(second, "\n")
        |> list.filter(fn(line) { line != "" })
        |> list.map(fn(line) {
          list.map(string.split(line, ","), fn(num) {
            result.unwrap(int.parse(num), 0)
          })
        })

      #(print_rules, numbers)
    }
    _ -> panic as "Invalid input"
  }
}

fn check_rest(
  line: List(Int),
  print_rules: Dict(Int, Set(Int)),
  before: Int,
) -> Bool {
  list.fold(line, True, fn(acc, after) {
    let rules =
      dict.get(print_rules, after)
      |> result.unwrap(set.new())
    let is_before_in_after_rules = set.contains(rules, before)
    case is_before_in_after_rules {
      True -> False
      False -> acc
    }
  })
}

fn check_line(line: List(Int), print_rules: Dict(Int, Set(Int))) -> Bool {
  case line {
    [] -> True
    [_] -> True
    [num, ..rest] ->
      check_rest(rest, print_rules, num) && check_line(rest, print_rules)
  }
}

fn get_middle(list: List(Int)) -> Option(Int) {
  let length = list.length(list)

  case length % 2 == 1 {
    True -> {
      let middle_index = length / 2
      list
      |> list.drop(middle_index)
      |> list.first
      |> option.from_result
    }
    False -> None
  }
}

pub fn day05_solution() -> Solution {
  Solution(
    part1: fn(input: String) -> SolutionResult {
      let #(print_rules, numbers) = parse_input(input)

      let correct_numbers =
        list.map(numbers, fn(line) {
          let is_correct = check_line(line, print_rules)
          let middle = get_middle(line)
          #(is_correct, middle)
        })

      let result =
        list.filter(correct_numbers, fn(b) { b.0 && option.is_some(b.1) })
        |> list.fold(0, fn(acc, b) {
          let #(_, middle) = b
          case middle {
            Some(m) -> acc + m
            None -> acc
          }
        })
      IntResult(result)
    },
    part2: fn(input: String) -> SolutionResult {
      let #(print_rules, numbers) = parse_input(input)

      let correct_numbers =
        list.filter(numbers, fn(line) { !check_line(line, print_rules) })
        |> list.map(fn(line) {
          list.sort(line, fn(a, b) {
            let a_set =
              dict.get(print_rules, a)
              |> result.unwrap(set.new())
            let b_set =
              dict.get(print_rules, b)
              |> result.unwrap(set.new())

            case set.contains(a_set, b), set.contains(b_set, a) {
              True, False -> order.Lt
              False, True -> order.Gt
              _, _ -> order.Eq
            }
          })
        })

      let result =
        list.fold(correct_numbers, 0, fn(acc, line) {
          let middle = get_middle(line)

          case middle {
            Some(m) -> acc + m
            None -> acc
          }
        })
      IntResult(result)
    },
  )
}
