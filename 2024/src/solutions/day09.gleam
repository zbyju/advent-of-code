import aoc_solution.{type Solution, type SolutionResult, IntResult, Solution}
import gleam/int
import gleam/list
import gleam/string

type Block {
  FileSpace(id: Int, size: Int)
  FreeSpace(size: Int)
}

fn parse_input(input: String) -> List(Block) {
  string.split(input, "")
  |> list.filter(fn(block) { block != "\n" })
  |> list.index_map(fn(block, index) {
    let size = int.parse(block)
    let file_id = index / 2
    case size, index % 2 == 0 {
      Ok(size), True -> FileSpace(id: file_id, size: size)
      Ok(size), False -> FreeSpace(size: size)
      _, _ -> FreeSpace(size: 0)
    }
  })
}

fn rearrange(
  blocks: List(Block),
  free_space: Int,
  acc: List(Block),
) -> List(Block) {
  case blocks, free_space > 0 {
    [], _ -> acc
    [FileSpace(_, _) as file, ..rest], False ->
      rearrange(rest, 0, list.append(acc, [file]))
    [FreeSpace(size), ..rest], False -> rearrange(rest, size, acc)
    blocks, True ->
      case list.reverse(blocks) {
        [FileSpace(_, size) as file, ..rest] if size <= free_space ->
          rearrange(
            list.reverse(rest),
            free_space - size,
            list.append(acc, [file]),
          )
        [FileSpace(id, size), ..rest] ->
          rearrange(
            list.append(list.reverse(rest), [FileSpace(id, size - free_space)]),
            0,
            list.append(acc, [FileSpace(id, size: free_space)]),
          )
        [_, ..rest] -> rearrange(list.reverse(rest), free_space, acc)
        [] -> acc
      }
  }
}

fn inner_rearrange(
  blocks: List(Block),
  right_id: Int,
  right_space: Int,
  has_moved: Bool,
  acc: List(Block),
) -> List(Block) {
  case blocks {
    [] -> acc
    [FileSpace(id, space), ..rest] if id == right_id && has_moved ->
      inner_rearrange(
        rest,
        right_id,
        right_space,
        has_moved,
        list.append(acc, [FreeSpace(space)]),
      )
    [FileSpace(id, _) as block, ..rest] if id == right_id && !has_moved ->
      inner_rearrange(
        rest,
        right_id,
        right_space,
        True,
        list.append(acc, [block]),
      )
    [FreeSpace(size), ..rest] if size >= right_space && !has_moved -> {
      let replacement = case size == right_space {
        True -> [FileSpace(id: right_id, size: right_space)]
        False -> [
          FileSpace(id: right_id, size: right_space),
          FreeSpace(size: size - right_space),
        ]
      }
      inner_rearrange(
        rest,
        right_id,
        right_space,
        True,
        list.append(acc, replacement),
      )
    }
    [block, ..rest] ->
      inner_rearrange(
        rest,
        right_id,
        right_space,
        has_moved,
        list.append(acc, [block]),
      )
  }
}

fn rearrange2(rev_blocks: List(Block), acc: List(Block)) -> List(Block) {
  case rev_blocks {
    [] -> acc
    [FreeSpace(_), ..rest] -> rearrange2(rest, acc)
    [FileSpace(right_id, right_space), ..rest] -> {
      let next_blocks = inner_rearrange(acc, right_id, right_space, False, [])
      rearrange2(rest, next_blocks)
    }
  }
}

pub fn day09_solution() -> Solution {
  Solution(
    part1: fn(input: String) -> SolutionResult {
      let blocks = parse_input(input)

      let rearranged = rearrange(blocks, 0, [])

      let result =
        list.flat_map(rearranged, fn(block) {
          case block {
            FileSpace(id, size) -> list.repeat(id, size)
            FreeSpace(_) -> list.repeat(0, 0)
          }
        })
        |> list.index_fold(0, fn(acc, id, index) { id * index + acc })
      IntResult(result)
    },
    part2: fn(input: String) -> SolutionResult {
      let blocks = parse_input(input)

      let rearranged = rearrange2(list.reverse(blocks), blocks)

      let mapped =
        list.flat_map(rearranged, fn(block) {
          case block {
            FileSpace(id, size) -> list.repeat(id, size)
            FreeSpace(size) -> list.repeat(0, size)
          }
        })

      let result =
        list.index_fold(mapped, 0, fn(acc, id, index) { id * index + acc })
      IntResult(result)
    },
  )
}
