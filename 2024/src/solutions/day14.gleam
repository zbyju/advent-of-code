import aoc_solution.{type Solution, type SolutionResult, IntResult, Solution}
import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/set.{type Set}
import gleam/string
import simplifile

type Robot {
  Robot(x: Int, y: Int, vx: Int, vy: Int)
}

type GridSize {
  GridSize(width: Int, height: Int)
}

type Position {
  Position(x: Int, y: Int)
}

fn floored_mod(a: Int, b: Int) -> Int {
  { { a % b } + b } % b
}

fn parse_num(str: String) -> #(Int, Int) {
  let nums =
    string.drop_start(str, 2)
    |> string.split(",")
    |> list.map(fn(x) { int.parse(x) |> result.unwrap(0) })
  case nums {
    [a, b] -> #(a, b)
    _ -> #(0, 0)
  }
}

fn parse_robots(section: String) -> List(Robot) {
  string.split(section, "\n")
  |> list.filter(fn(x) { x != "" })
  |> list.map(fn(line) {
    let parts = string.split(line, " ")
    case parts {
      [pos, vel] -> {
        let #(x, y) = parse_num(pos)
        let #(vx, vy) = parse_num(vel)
        Robot(x, y, vx, vy)
      }
      _ -> Robot(0, 0, 0, 0)
    }
  })
}

fn parse_size(str: String) -> #(Int, Int) {
  let nums =
    string.split(str, " ")
    |> list.filter(fn(x) { x != "" })
    |> list.map(fn(x) { int.parse(x) |> result.unwrap(0) })

  case nums {
    [a, b] -> #(a, b)
    _ -> #(0, 0)
  }
}

fn parse_input(input: String) -> #(List(Robot), GridSize) {
  let sections = string.split(input, "\n\n") |> list.filter(fn(x) { x != "" })
  case sections {
    [grid_size, robots] -> {
      let robots = parse_robots(robots)
      let #(width, height) = parse_size(grid_size)
      #(robots, GridSize(width, height))
    }
    _ -> #(list.new(), GridSize(0, 0))
  }
}

fn simulate_robots(robots: List(Robot), grid_size: GridSize) -> List(Robot) {
  list.map(robots, fn(r) {
    let px = floored_mod(r.x + r.vx, grid_size.width)
    let py = floored_mod(r.y + r.vy, grid_size.height)
    Robot(px, py, r.vx, r.vy)
  })
}

fn check_christmas_tree(robots: List(Robot)) -> Bool {
  let set =
    list.fold_until(robots, set.new(), fn(acc, r) {
      let position = Position(r.x, r.y)
      let length = set.size(acc)
      let new_set = set.insert(acc, position)
      case length == set.size(new_set) {
        True -> list.Stop(new_set)
        False -> list.Continue(new_set)
      }
    })
  set.size(set) == list.length(robots)
}

fn seconds_until_christmas_tree(
  robots: List(Robot),
  grid_size: GridSize,
  acc: Int,
) -> Int {
  case check_christmas_tree(robots) {
    True -> acc
    False ->
      seconds_until_christmas_tree(
        simulate_robots(robots, grid_size),
        grid_size,
        acc + 1,
      )
  }
}

pub fn day14_solution() -> Solution {
  Solution(
    part1: fn(input: String) -> SolutionResult {
      let seconds = 100

      let #(robots, grid_size) = parse_input(input)
      let half_width = grid_size.width / 2
      let half_height = grid_size.height / 2

      let simulated_robots =
        list.map(robots, fn(r) {
          let px = floored_mod(r.x + { r.vx * seconds }, grid_size.width)
          let py = floored_mod(r.y + { r.vy * seconds }, grid_size.height)
          Robot(px, py, r.vx, r.vy)
        })
        |> list.fold(#(0, 0, 0, 0), fn(acc, r) {
          case r {
            _ if r.x < half_width && r.y < half_height -> #(
              acc.0 + 1,
              acc.1,
              acc.2,
              acc.3,
            )
            _ if r.x > half_width && r.y < half_height -> #(
              acc.0,
              acc.1 + 1,
              acc.2,
              acc.3,
            )
            _ if r.x < half_width && r.y > half_height -> #(
              acc.0,
              acc.1,
              acc.2 + 1,
              acc.3,
            )
            _ if r.x > half_width && r.y > half_height -> #(
              acc.0,
              acc.1,
              acc.2,
              acc.3 + 1,
            )
            _ -> acc
          }
        })

      let result =
        simulated_robots.0
        * simulated_robots.1
        * simulated_robots.2
        * simulated_robots.3

      IntResult(result)
    },
    part2: fn(input: String) -> SolutionResult {
      let #(robots, grid_size) = parse_input(input)

      let seconds = seconds_until_christmas_tree(robots, grid_size, 0)

      IntResult(seconds)
    },
  )
}
