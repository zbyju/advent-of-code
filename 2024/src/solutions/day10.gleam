import aoc_solution.{type Solution, type SolutionResult, IntResult, Solution}
import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/iterator
import gleam/list
import gleam/result
import gleam/set.{type Set}
import gleam/string

import glearray.{type Array}

type Tile {
  Tile(height: Int)
}

type Coords {
  Coords(x: Int, y: Int)
}

fn parse_input(input: String) -> #(Array(Array(Tile)), List(Coords)) {
  let grid =
    string.split(input, "\n")
    |> list.filter(fn(line) { line != "" })
    |> glearray.from_list()
    |> glearray.iterate()
    |> iterator.map(fn(line) {
      string.split(line, "")
      |> glearray.from_list()
      |> glearray.iterate()
      |> iterator.map(fn(char) {
        let int = int.parse(char)
        case int {
          Ok(height) -> Tile(height)
          Error(_) -> Tile(height: 9999)
        }
      })
      |> iterator.to_list()
      |> glearray.from_list()
    })
    |> iterator.to_list()
    |> glearray.from_list()

  let heads =
    glearray.iterate(grid)
    |> iterator.index()
    |> iterator.fold(set.new(), fn(acc, line_tuple) {
      let #(line, y) = line_tuple
      let heads_line =
        glearray.iterate(line)
        |> iterator.index()
        |> iterator.fold(set.new(), fn(acc, tile_tuple) {
          let #(tile, x) = tile_tuple
          case tile {
            Tile(height: height) if height == 0 ->
              set.insert(acc, Coords(x: x, y: y))
            _ -> acc
          }
        })

      set.union(acc, heads_line)
    })
    |> set.to_list()

  #(grid, heads)
}

fn get_position(grid: Array(Array(Tile)), pos: Coords) -> Tile {
  let Coords(x, y) = pos
  let tile =
    glearray.get(grid, y)
    |> result.try(fn(row) { glearray.get(row, x) })

  case tile {
    Ok(tile) -> tile
    Error(_) -> Tile(height: 9999)
  }
}

fn can_move(from: Tile, to: Tile) -> Bool {
  case from, to {
    Tile(height: from_height), Tile(height: to_height) ->
      from_height < to_height
  }
}

fn get_neighbors(pos: Coords) -> List(Coords) {
  let Coords(x, y) = pos
  [
    Coords(x: x - 1, y: y),
    Coords(x: x + 1, y: y),
    Coords(x: x, y: y - 1),
    Coords(x: x, y: y + 1),
  ]
}

fn find_peaks(
  grid: Array(Array(Tile)),
  queue: List(Coords),
  visited: Set(Coords),
  count: Int,
) {
  case queue {
    [] -> count
    [head, ..rest] -> {
      let is_visited = set.contains(visited, head)

      case is_visited {
        True -> find_peaks(grid, rest, visited, count)
        False -> {
          let next_visited = set.insert(visited, head)
          let current_height = get_position(grid, head).height
          let neighbors =
            get_neighbors(head)
            |> list.filter(fn(neighbor) {
              let neighbor_height = get_position(grid, neighbor).height
              neighbor_height - current_height == 1
            })
          case current_height {
            9 ->
              find_peaks(
                grid,
                list.append(rest, neighbors),
                next_visited,
                count + 1,
              )
            _ ->
              find_peaks(
                grid,
                list.append(rest, neighbors),
                next_visited,
                count,
              )
          }
        }
      }
    }
  }
}

fn find_ratings(
  grid: Array(Array(Tile)),
  queue: List(Coords),
  count: Int,
) -> Int {
  case queue {
    [] -> count
    [head, ..rest] -> {
      let current_height = get_position(grid, head).height
      let neighbors =
        get_neighbors(head)
        |> list.filter(fn(neighbor) {
          let neighbor_height = get_position(grid, neighbor).height
          neighbor_height - current_height == 1
        })
      case current_height {
        9 -> find_ratings(grid, list.append(rest, neighbors), count + 1)
        _ -> find_ratings(grid, list.append(rest, neighbors), count)
      }
    }
  }
}

pub fn day10_solution() -> Solution {
  Solution(
    part1: fn(input: String) -> SolutionResult {
      let #(grid, starts) = parse_input(input)

      let counts =
        list.fold(starts, 0, fn(acc, start) {
          let count = find_peaks(grid, [start], set.new(), 0)
          acc + count
        })

      IntResult(counts)
    },
    part2: fn(input: String) -> SolutionResult {
      let #(grid, starts) = parse_input(input)

      let ratings =
        list.fold(starts, 0, fn(acc, start) {
          let rating = find_ratings(grid, [start], 0)
          acc + rating
        })

      IntResult(ratings)
    },
  )
}
