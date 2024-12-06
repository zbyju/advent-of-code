import aoc_solution.{type Solution, type SolutionResult, IntResult, Solution}
import gleam/io
import gleam/iterator
import gleam/list
import gleam/option.{type Option, None, Some}
import gleam/string
import glearray.{type Array}

type Letter {
  X
  M
  A
  S
  E
}

fn parse_input(input: String) -> Array(Array(Letter)) {
  let lists =
    string.split(input, "\n")
    |> list.filter(fn(line) { string.length(line) > 0 })
    |> list.map(fn(line) {
      string.split(line, "")
      |> list.map(fn(letter) {
        case letter {
          "X" -> X
          "M" -> M
          "A" -> A
          "S" -> S
          _ -> E
        }
      })
    })

  glearray.from_list(list.map(lists, fn(l) { glearray.from_list(l) }))
}

fn get_position(grid: Array(Array(Letter)), x: Int, y: Int) -> Option(Letter) {
  let maybe_row = glearray.get(grid, y)
  case maybe_row {
    Error(_) -> None
    Ok(row) -> {
      let maybe_letter = glearray.get(row, x)
      case maybe_letter {
        Error(_) -> None
        Ok(letter) -> Some(letter)
      }
    }
  }
}

fn check_mas(grid: Array(Array(Letter)), x: Int, y: Int) -> Int {
  let top_right = get_position(grid, x + 1, y + 1)
  let top_left = get_position(grid, x - 1, y + 1)
  let bottom_right = get_position(grid, x + 1, y - 1)
  let bottom_left = get_position(grid, x - 1, y - 1)

  case top_right, top_left, bottom_right, bottom_left {
    Some(M), Some(M), Some(S), Some(S) -> 1
    Some(S), Some(S), Some(M), Some(M) -> 1
    Some(S), Some(M), Some(S), Some(M) -> 1
    Some(M), Some(S), Some(M), Some(S) -> 1
    _, _, _, _ -> 0
  }
}

fn check_xmas_in_dir(
  grid: Array(Array(Letter)),
  looking_for: Letter,
  x: Int,
  y: Int,
  dx: Int,
  dy: Int,
) -> Int {
  let cx = x + dx
  let cy = y + dy
  let looking_at = get_position(grid, cx, cy)

  case looking_at {
    Some(looking_at) if looking_at == looking_for -> {
      case looking_for {
        X -> check_xmas_in_dir(grid, M, cx, cy, dx, dy)
        M -> check_xmas_in_dir(grid, A, cx, cy, dx, dy)
        A -> check_xmas_in_dir(grid, S, cx, cy, dx, dy)
        S -> 1
        _ -> 0
      }
    }
    _ -> 0
  }
}

fn count(grid: Array(Array(Letter))) -> Int {
  glearray.iterate(grid)
  |> iterator.index()
  |> iterator.fold(0, fn(acc, r) {
    let #(row, y) = r
    glearray.iterate(row)
    |> iterator.index()
    |> iterator.fold(acc, fn(acc, l) {
      let #(letter, x) = l
      case letter {
        X -> {
          let up = check_xmas_in_dir(grid, M, x, y, 0, -1)
          let down = check_xmas_in_dir(grid, M, x, y, 0, 1)
          let right = check_xmas_in_dir(grid, M, x, y, 1, 0)
          let left = check_xmas_in_dir(grid, M, x, y, -1, 0)

          let up_right = check_xmas_in_dir(grid, M, x, y, 1, -1)
          let down_right = check_xmas_in_dir(grid, M, x, y, 1, 1)
          let down_left = check_xmas_in_dir(grid, M, x, y, -1, 1)
          let up_left = check_xmas_in_dir(grid, M, x, y, -1, -1)

          io.debug(x)
          io.debug(y)
          io.debug(
            up
            + down
            + right
            + left
            + up_right
            + down_right
            + down_left
            + up_left,
          )

          acc
          + up
          + down
          + right
          + left
          + up_right
          + down_right
          + down_left
          + up_left
        }

        _ -> acc
      }
    })
  })
}

fn count2(grid: Array(Array(Letter))) -> Int {
  glearray.iterate(grid)
  |> iterator.index()
  |> iterator.fold(0, fn(acc, r) {
    let #(row, y) = r
    glearray.iterate(row)
    |> iterator.index()
    |> iterator.fold(acc, fn(acc, l) {
      let #(letter, x) = l
      case letter {
        A -> {
          acc + check_mas(grid, x, y)
        }

        _ -> acc
      }
    })
  })
}

pub fn day04_solution() -> Solution {
  Solution(
    part1: fn(input: String) -> SolutionResult {
      let grid = parse_input(input)
      let found = count(grid)

      IntResult(found)
    },
    part2: fn(input: String) -> SolutionResult {
      let grid = parse_input(input)
      let found = count2(grid)

      IntResult(found)
    },
  )
}
