import aoc_solution.{type Solution, type SolutionResult, IntResult, Solution}
import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string

type Equation {
  Equation(result: Int, nums: List(Int))
}

fn parse_input(input: String) -> List(Equation) {
  let lines =
    string.split(input, "\n")
    |> list.filter(fn(line) { string.length(line) > 0 })

  list.map(lines, fn(line) {
    let parts = string.split(line, ": ")
    let result =
      list.first(parts)
      |> result.try(fn(r) { int.parse(r) })
      |> result.unwrap(0)

    let nums =
      list.last(parts)
      |> result.map(fn(n) { string.split(n, " ") })
      |> result.unwrap(list.new())
      |> list.map(fn(n) { result.unwrap(int.parse(n), 0) })

    Equation(result, nums)
  })
}

fn check_equation(equation: Equation, acc: Int) -> Bool {
  case equation.nums {
    [] -> equation.result == acc
    [num, ..rest] -> {
      check_equation(Equation(equation.result, rest), acc + num)
      || check_equation(Equation(equation.result, rest), acc * num)
    }
  }
}

fn combine(x: Int, y: Int) -> Int {
  let x_str = int.to_string(x)
  let y_str = int.to_string(y)
  int.parse(x_str <> y_str) |> result.unwrap(0)
}

fn check_equation2(equation: Equation, acc: Int) -> Bool {
  case equation.nums {
    [] -> equation.result == acc
    [num, ..rest] -> {
      check_equation2(Equation(equation.result, rest), acc + num)
      || check_equation2(Equation(equation.result, rest), acc * num)
      || check_equation2(Equation(equation.result, rest), combine(acc, num))
    }
  }
}

pub fn day07_solution() -> Solution {
  Solution(
    part1: fn(input: String) -> SolutionResult {
      let equations = parse_input(input)
      let correct_equations =
        list.filter(equations, fn(e) {
          let without_first = list.rest(e.nums) |> result.unwrap(list.new())
          let first = e.nums |> list.first() |> result.unwrap(0)
          check_equation(Equation(e.result, without_first), first)
        })

      let sum = list.fold(correct_equations, 0, fn(acc, e) { acc + e.result })

      IntResult(sum)
    },
    part2: fn(input: String) -> SolutionResult {
      let equations = parse_input(input)
      let correct_equations =
        list.filter(equations, fn(e) {
          let without_first = list.rest(e.nums) |> result.unwrap(list.new())
          let first = e.nums |> list.first() |> result.unwrap(0)
          check_equation2(Equation(e.result, without_first), first)
        })

      let sum = list.fold(correct_equations, 0, fn(acc, e) { acc + e.result })

      IntResult(sum)
    },
  )
}
