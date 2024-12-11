import aoc_solution.{type Solution, type SolutionResult, IntResult, Solution}
import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string

fn parse_input(input: String) -> List(Int) {
  input
  |> string.trim()
  |> string.split(" ")
  |> list.filter(fn(x) { x != "" })
  |> list.map(int.parse)
  |> result.values()
}

fn next_stones(stone: Int) -> List(Int) {
  case stone {
    0 -> [1]
    _ -> {
      let str = int.to_string(stone)
      let len = string.length(str)
      case len % 2 == 0 {
        True -> {
          let half = len / 2
          let left = string.slice(str, 0, half)
          let right = string.slice(str, half, string.length(str))
          let l = int.parse(left) |> result.unwrap(0)
          let r = int.parse(right) |> result.unwrap(0)
          [l, r]
        }
        False -> [stone * 2024]
      }
    }
  }
}

type CacheKey =
  #(Int, Int)

fn rec_multiple(
  stones: List(Int),
  steps_remaining: Int,
  cache: Dict(CacheKey, Int),
) -> #(Dict(CacheKey, Int), Int) {
  list.fold(stones, #(cache, 0), fn(tuple, stone) {
    let #(cache, sum) = tuple
    let res = rec(stone, steps_remaining, cache)
    #(res.0, sum + res.1)
  })
}

fn rec(
  stone: Int,
  steps_remaining: Int,
  cache: Dict(CacheKey, Int),
) -> #(Dict(CacheKey, Int), Int) {
  case steps_remaining {
    0 -> #(cache, 1)
    _ -> {
      let key = #(stone, steps_remaining)
      case dict.get(cache, key) {
        Ok(value) -> #(cache, value)
        Error(_) -> {
          let next_stones = next_stones(stone)
          let #(next_cache, next_count) =
            rec_multiple(next_stones, steps_remaining - 1, cache)
          #(
            dict.insert(next_cache, #(stone, steps_remaining), next_count),
            next_count,
          )
        }
      }
    }
  }
}

pub fn day11_solution() -> Solution {
  Solution(
    part1: fn(input: String) -> SolutionResult {
      let stones = parse_input(input)

      let new_stones = rec_multiple(stones, 25, dict.new()).1

      IntResult(new_stones)
    },
    part2: fn(input: String) -> SolutionResult {
      let stones = parse_input(input)

      let new_stones = rec_multiple(stones, 75, dict.new()).1

      IntResult(new_stones)
    },
  )
}
