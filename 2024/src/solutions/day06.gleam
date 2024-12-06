import aoc_solution.{type Solution, type SolutionResult, IntResult, Solution}
import gleam/io
import gleam/iterator
import gleam/list
import gleam/result
import gleam/set.{type Set}
import gleam/string
import glearray.{type Array}

type Tile {
  Wall
  Floor
  Out
}

type Direction {
  Up
  Down
  Left
  Right
}

type Coords {
  Coords(x: Int, y: Int)
}

type Guard {
  Guard(coords: Coords, facing: Direction)
}

fn parse_input(input: String) -> #(Array(Array(Tile)), Coords) {
  let lines =
    string.split(input, "\n")
    |> list.filter(fn(line) { string.length(line) > 0 })

  let tile_lists =
    list.map(lines, fn(line) {
      let line_list = string.split(line, "")
      list.map(line_list, fn(tile) {
        case tile {
          "#" -> Wall
          _ -> Floor
        }
      })
    })

  let tiles =
    glearray.from_list(list.map(tile_lists, fn(l) { glearray.from_list(l) }))

  let player_coords =
    list.index_fold(lines, Coords(-1, -1), fn(acc, line, y) {
      case acc {
        Coords(-1, -1) -> {
          let x =
            string.split(line, "")
            |> list.index_fold(-1, fn(acc, tile, x) {
              case tile {
                "^" -> x
                _ -> acc
              }
            })
          case x {
            -1 -> acc
            _ -> Coords(x, y)
          }
        }
        _ -> acc
      }
    })

  #(tiles, player_coords)
}

fn next_direction(facing: Direction) -> Direction {
  case facing {
    Up -> Right
    Right -> Down
    Down -> Left
    Left -> Up
  }
}

fn next_coords(coords: Coords, facing: Direction) -> Coords {
  let #(dx, dy) = case facing {
    Up -> #(0, -1)
    Down -> #(0, 1)
    Left -> #(-1, 0)
    Right -> #(1, 0)
  }
  Coords(coords.x + dx, coords.y + dy)
}

fn get_tile(tiles: Array(Array(Tile)), coords: Coords) -> Tile {
  glearray.get(tiles, coords.y)
  |> result.unwrap(glearray.from_list([]))
  |> glearray.get(coords.x)
  |> result.unwrap(Out)
}

fn get_tile2(
  tiles: Array(Array(Tile)),
  coords: Coords,
  obstacle: Coords,
) -> Tile {
  case coords == obstacle {
    True -> Wall
    False -> get_tile(tiles, coords)
  }
}

fn simulate_guard(
  tiles: Array(Array(Tile)),
  guard: Coords,
  facing: Direction,
  visited: Set(Coords),
  acc: Int,
) -> Int {
  let is_out = get_tile(tiles, guard)

  case is_out {
    Out -> set.size(visited)
    _ -> {
      let next_visited = set.insert(visited, guard)

      let next_dir =
        list.fold_until([0, 1, 2, 3], facing, fn(acc, _) {
          let forward = next_coords(guard, acc)
          let tile_ahead = get_tile(tiles, forward)
          case tile_ahead {
            Wall -> list.Continue(next_direction(acc))
            Floor | Out -> list.Stop(acc)
          }
        })
      let next_guard = next_coords(guard, next_dir)
      simulate_guard(tiles, next_guard, next_dir, next_visited, acc + 1)
    }
  }
}

fn is_loop(
  tiles: Array(Array(Tile)),
  guard: Coords,
  facing: Direction,
  obstacle: Coords,
  visited: Set(Guard),
) -> Bool {
  let is_out = get_tile(tiles, guard)
  let g = Guard(guard, facing)

  case is_out {
    Out -> False
    _ -> {
      let has_visited = set.contains(visited, g)
      case has_visited {
        True -> True
        False -> {
          let g = Guard(guard, facing)
          let next_visited = set.insert(visited, g)

          let next_dir =
            list.fold_until([0, 1, 2, 3], facing, fn(acc, _) {
              let forward = next_coords(guard, acc)
              let tile_ahead = get_tile2(tiles, forward, obstacle)
              case tile_ahead {
                Wall -> list.Continue(next_direction(acc))
                Floor | Out -> list.Stop(acc)
              }
            })
          let next_guard = next_coords(guard, next_dir)
          is_loop(tiles, next_guard, next_dir, obstacle, next_visited)
        }
      }
    }
  }
}

pub fn day06_solution() -> Solution {
  Solution(
    part1: fn(input: String) -> SolutionResult {
      let #(tiles, player_coords) = parse_input(input)
      let steps = simulate_guard(tiles, player_coords, Up, set.new(), 0)

      IntResult(steps)
    },
    part2: fn(input: String) -> SolutionResult {
      let #(tiles, player_coords) = parse_input(input)
      let loops =
        glearray.iterate(tiles)
        |> iterator.index()
        |> iterator.fold(0, fn(acc, row_entry) {
          let #(row, y) = row_entry
          glearray.iterate(row)
          |> iterator.index()
          |> iterator.fold(acc, fn(acc, tile_entry) {
            let #(tile, x) = tile_entry
            case tile {
              Wall | Out -> acc
              Floor -> {
                let loop =
                  is_loop(tiles, player_coords, Up, Coords(x, y), set.new())

                case loop {
                  True -> acc + 1
                  False -> acc
                }
              }
            }
          })
        })

      IntResult(loops)
    },
  )
}
