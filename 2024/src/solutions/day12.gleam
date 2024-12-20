import aoc_solution.{type Solution, type SolutionResult, IntResult, Solution}
import gleam/io
import gleam/list
import gleam/option.{type Option, None, Some}
import gleam/result
import gleam/set.{type Set}
import gleam/string
import glearray.{type Array}

type Grid =
  Array(Array(Int))

type Score {
  Score(perimeter: Int, area: Int)
}

type Coords {
  Coords(x: Int, y: Int)
}

fn ascii(str: String) -> Int {
  str
  |> string.to_utf_codepoints()
  |> list.first()
  |> result.map(fn(x) { string.utf_codepoint_to_int(x) })
  |> result.unwrap(0)
}

fn parse_input(input: String) -> Grid {
  input
  |> string.split("\n")
  |> list.filter(fn(x) { x != "" })
  |> list.map(fn(line) {
    line
    |> string.split("")
    |> list.filter(fn(x) { x != "" })
    |> list.map(fn(x) { ascii(x) })
    |> glearray.from_list()
  })
  |> glearray.from_list()
}

fn neighbours(pos: Coords) -> List(Coords) {
  let Coords(x, y) = pos
  [
    Coords(x: x - 1, y: y),
    Coords(x: x + 1, y: y),
    Coords(x: x, y: y - 1),
    Coords(x: x, y: y + 1),
  ]
}

fn value(grid: Grid, pos: Coords) -> Int {
  glearray.get(grid, pos.y)
  |> result.try(fn(row) { glearray.get(row, pos.x) })
  |> result.unwrap(0)
}

fn get_next(
  q1: List(Coords),
  q2: List(Coords),
) -> Option(#(Coords, List(Coords), List(Coords))) {
  case q1 {
    [] ->
      case q2 {
        [] -> None
        [x, ..r2] -> Some(#(x, q1, r2))
      }
    [x, ..r1] -> Some(#(x, r1, q2))
  }
}

fn split_neighbours(
  grid: Grid,
  current: Coords,
  neighbours: List(Coords),
) -> #(List(Coords), List(Coords)) {
  list.partition(neighbours, fn(x) { value(grid, x) == value(grid, current) })
}

type FloodComponentReturn {
  // @return score (of the component), queue (of the next component points), visited (TODO: only the perimeter points)
  FloodComponentReturn(score: Score, queue: List(Coords), visited: Set(Coords))
}

fn flood_component(
  grid: Grid,
  queue: List(Coords),
  visited: Set(Coords),
  acc: FloodComponentReturn,
) -> FloodComponentReturn {
  case queue {
    [] -> acc
    [current, ..rest] -> {
      let is_visited = set.contains(visited, current)
      let current_value = value(grid, current)
      case is_visited, current_value == 0 {
        True, _ -> flood_component(grid, rest, visited, acc)
        _, True -> flood_component(grid, rest, visited, acc)
        _, _ -> {
          let neighbours = neighbours(current)
          let #(same_neighbours, other_neighbours) =
            split_neighbours(grid, current, neighbours)

          let next_queue = list.append(rest, same_neighbours)
          let next_visited = set.insert(visited, current)
          let next_score =
            Score(
              perimeter: acc.score.perimeter + list.length(other_neighbours),
              area: acc.score.area + 1,
            )
          let next_queue_acc = list.append(acc.queue, other_neighbours)

          flood_component(
            grid,
            next_queue,
            next_visited,
            FloodComponentReturn(next_score, next_queue_acc, next_visited),
          )
        }
      }
    }
  }
}

fn flood(
  grid: Grid,
  queue: List(Coords),
  visited: Set(Coords),
  acc: List(Score),
) -> List(Score) {
  case queue {
    [] -> acc
    [current, ..rest] -> {
      let is_visited = set.contains(visited, current)
      case is_visited {
        True -> flood(grid, rest, visited, acc)
        False -> {
          let result =
            flood_component(
              grid,
              [current],
              visited,
              FloodComponentReturn(
                Score(perimeter: 0, area: 0),
                list.new(),
                visited,
              ),
            )
          let next_score = list.append(acc, [result.score])
          let next_queue = list.append(rest, result.queue)
          flood(grid, next_queue, result.visited, next_score)
        }
      }
    }
  }
}

fn is_north_side(grid: Grid, pos: Coords) -> Bool {
  let val = value(grid, pos)

  let top = value(grid, Coords(x: pos.x, y: pos.y - 1))
  let right = value(grid, Coords(x: pos.x + 1, y: pos.y))
  let top_right = value(grid, Coords(x: pos.x + 1, y: pos.y - 1))

  case val == top, val == right, val == top_right {
    True, _, _ -> False
    _, False, _ -> True
    _, True, True -> True
    _, True, False -> False
  }
}

fn is_south_side(grid: Grid, pos: Coords) -> Bool {
  let val = value(grid, pos)

  let bot = value(grid, Coords(x: pos.x, y: pos.y + 1))
  let right = value(grid, Coords(x: pos.x + 1, y: pos.y))
  let bot_right = value(grid, Coords(x: pos.x + 1, y: pos.y + 1))

  case val == bot, val == right, val == bot_right {
    True, _, _ -> False
    _, False, _ -> True
    _, True, True -> True
    _, True, False -> False
  }
}

fn is_west_side(grid: Grid, pos: Coords) -> Bool {
  let val = value(grid, pos)

  let left = value(grid, Coords(x: pos.x - 1, y: pos.y))
  let top = value(grid, Coords(x: pos.x, y: pos.y - 1))
  let top_left = value(grid, Coords(x: pos.x - 1, y: pos.y - 1))

  case val == left, val == top, val == top_left {
    True, _, _ -> False
    _, False, _ -> True
    _, True, True -> True
    _, True, False -> False
  }
}

fn is_east_side(grid: Grid, pos: Coords) -> Bool {
  let val = value(grid, pos)

  let right = value(grid, Coords(x: pos.x + 1, y: pos.y))
  let top = value(grid, Coords(x: pos.x, y: pos.y - 1))
  let top_right = value(grid, Coords(x: pos.x + 1, y: pos.y - 1))

  case val == right, val == top, val == top_right {
    True, _, _ -> False
    _, False, _ -> True
    _, True, True -> True
    _, True, False -> False
  }
}

type Score2 {
  Score2(sides: Int, area: Int)
}

type FloodComponent2Return {
  FloodComponent2Return(
    score: Score2,
    queue: List(Coords),
    visited: Set(Coords),
  )
}

fn flood_component2(
  grid: Grid,
  queue: List(Coords),
  visited: Set(Coords),
  acc: FloodComponent2Return,
) -> FloodComponent2Return {
  case queue {
    [] -> acc
    [current, ..rest] -> {
      let is_visited = set.contains(visited, current)
      let current_value = value(grid, current)
      case is_visited, current_value == 0 {
        True, _ -> flood_component2(grid, rest, visited, acc)
        _, True -> flood_component2(grid, rest, visited, acc)
        _, _ -> {
          let neighbours = neighbours(current)
          let #(same_neighbours, other_neighbours) =
            split_neighbours(grid, current, neighbours)
          let sides =
            [
              is_north_side(grid, current),
              is_south_side(grid, current),
              is_west_side(grid, current),
              is_east_side(grid, current),
            ]
            |> list.filter(fn(x) { x })

          let next_queue = list.append(rest, same_neighbours)
          let next_visited = set.insert(visited, current)
          let next_score =
            Score2(
              sides: acc.score.sides + list.length(sides),
              area: acc.score.area + 1,
            )
          let next_queue_acc = list.append(acc.queue, other_neighbours)

          flood_component2(
            grid,
            next_queue,
            next_visited,
            FloodComponent2Return(next_score, next_queue_acc, next_visited),
          )
        }
      }
    }
  }
}

fn flood2(
  grid: Grid,
  queue: List(Coords),
  visited: Set(Coords),
  acc: List(Score2),
) -> List(Score2) {
  case queue {
    [] -> acc
    [current, ..rest] -> {
      let is_visited = set.contains(visited, current)
      case is_visited {
        True -> flood2(grid, rest, visited, acc)
        False -> {
          let result =
            flood_component2(
              grid,
              [current],
              visited,
              FloodComponent2Return(
                Score2(sides: 0, area: 0),
                list.new(),
                visited,
              ),
            )
          let next_score = list.append(acc, [result.score])
          let next_queue = list.append(rest, result.queue)
          flood2(grid, next_queue, result.visited, next_score)
        }
      }
    }
  }
}

pub fn day12_solution() -> Solution {
  Solution(
    part1: fn(input: String) -> SolutionResult {
      let grid = parse_input(input)
      let scores = flood(grid, [Coords(x: 0, y: 0)], set.new(), list.new())
      let result =
        list.fold(scores, 0, fn(acc, x) { acc + { x.area * x.perimeter } })
      IntResult(result)
    },
    part2: fn(input: String) -> SolutionResult {
      let grid = parse_input(input)
      let scores = flood2(grid, [Coords(x: 0, y: 0)], set.new(), list.new())
      let result =
        list.fold(scores, 0, fn(acc, x) { acc + { x.area * x.sides } })
      IntResult(result)
    },
  )
}
