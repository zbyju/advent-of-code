import aoc_solution.{type Solution, type SolutionResult, IntResult, Solution}
import gleam/io
import gleam/list
import gleam/result
import gleam/set
import gleam/string

import gleam/dict.{type Dict}
import gleam/option.{type Option, None, Some}

type Antenna {
  Antenna(x: Int, y: Int)
}

type Coords {
  Coords(x: Int, y: Int)
}

fn is_alphanumeric(x: Int) -> Bool {
  { x >= 48 && x <= 57 } || { x >= 65 && x <= 90 } || { x >= 97 && x <= 122 }
}

fn ascii_of_char(c: String) -> Option(Int) {
  let ascii =
    string.to_utf_codepoints(c)
    |> list.first()
    |> result.map(fn(x) { string.utf_codepoint_to_int(x) })

  case ascii {
    Ok(ascii) -> {
      let is_alphanumeric = is_alphanumeric(ascii)
      case is_alphanumeric {
        True -> Some(ascii)
        False -> None
      }
    }
    _ -> None
  }
}

fn parse_input(input: String) -> #(Dict(Int, List(Antenna)), Coords) {
  let lines =
    string.split(input, "\n")
    |> list.filter(fn(line) { !string.is_empty(line) })

  let y_max = list.length(lines)
  let x_max = list.first(lines) |> result.unwrap("") |> string.length()
  let antennas =
    list.index_fold(lines, dict.new(), fn(acc, line, y) {
      let chars = string.split(line, "")
      list.index_fold(chars, acc, fn(acc, char, x) {
        let value = ascii_of_char(char)

        case value {
          Some(value) ->
            dict.upsert(acc, value, fn(l) {
              case l {
                None -> [Antenna(x: x, y: y)]
                Some(l) -> list.append(l, [Antenna(x: x, y: y)])
              }
            })
          None -> acc
        }
      })
    })

  #(antennas, Coords(x: x_max, y: y_max))
}

fn is_within_bounds(coords: Coords, antenna: Antenna) -> Bool {
  let Coords(x: x, y: y) = coords
  let Antenna(x: ax, y: ay) = antenna
  ax >= 0 && ax < x && ay >= 0 && ay < y
}

fn create_antenna(a1: Antenna, a2: Antenna) -> Antenna {
  let dx = a1.x - a2.x
  let dy = a1.y - a2.y

  Antenna(x: a1.x + dx, y: a1.y + dy)
}

fn get_difference(a1: Antenna, a2: Antenna) -> Coords {
  Coords(x: a1.x - a2.x, y: a1.y - a2.y)
}

fn create_antenna2(a1: Antenna, diff: Coords) -> Antenna {
  Antenna(x: a1.x + diff.x, y: a1.y + diff.y)
}

fn create_antennas(
  max_coords: Coords,
  antenna: Antenna,
  difference: Coords,
  acc: List(Antenna),
) -> List(Antenna) {
  let created = create_antenna2(antenna, difference)

  io.debug(#(antenna, difference, created))

  case is_within_bounds(max_coords, created) {
    True -> {
      let next_acc = list.append(acc, [created])
      create_antennas(max_coords, created, difference, next_acc)
    }
    False -> acc
  }
}

pub fn day08_solution() -> Solution {
  Solution(
    part1: fn(input: String) -> SolutionResult {
      let #(antennas, max_coords) = parse_input(input)

      let result =
        dict.fold(antennas, set.new(), fn(acc, freq, antennas) {
          let combinations = list.combination_pairs(antennas)

          list.fold(combinations, acc, fn(acc, pair) {
            let #(a1, a2) = pair

            let n1 = create_antenna(a1, a2)
            let n2 = create_antenna(a2, a1)

            case
              is_within_bounds(max_coords, n1),
              is_within_bounds(max_coords, n2)
            {
              True, True -> {
                set.insert(acc, n1) |> set.insert(n2)
              }
              True, False -> set.insert(acc, n1)
              False, True -> set.insert(acc, n2)
              False, False -> acc
            }
          })
        })

      IntResult(set.size(result))
    },
    part2: fn(input: String) -> SolutionResult {
      let #(antennas, max_coords) = parse_input(input)

      let result =
        dict.fold(antennas, set.new(), fn(acc, freq, antennas) {
          let combinations = list.combination_pairs(antennas)

          list.fold(combinations, acc, fn(acc, pair) {
            let #(a1, a2) = pair

            let n1 = create_antennas(max_coords, a1, get_difference(a1, a2), [])
            let n2 = create_antennas(max_coords, a2, get_difference(a2, a1), [])
            let new_nodes = list.append(n1, n2)

            io.debug(#(freq, a1, a2, n1, n2, new_nodes))

            set.from_list(new_nodes)
            |> set.union(acc)
            |> set.insert(a1)
            |> set.insert(a2)
          })
        })

      IntResult(set.size(result))
    },
  )
}
