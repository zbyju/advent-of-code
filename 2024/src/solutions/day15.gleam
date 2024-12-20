import aoc_solution.{type Solution, type SolutionResult, IntResult, Solution}
import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/result
import gleam/string

type Coords {
  Coords(x: Int, y: Int)
}

type Tile {
  Wall
  Box
  Empty
}

type Tile2 {
  Wall2
  LeftBox
  RightBox
  Empty2
}

type Move {
  Up
  Down
  Left
  Right
}

type Parsed {
  Parsed(
    tiles: Dict(Coords, Tile),
    robot: Coords,
    moves: List(Move),
    grid_size: Coords,
  )
}

type Parsed2 {
  Parsed2(
    tiles: Dict(Coords, Tile2),
    robot: Coords,
    moves: List(Move),
    grid_size: Coords,
  )
}

fn is_inside(pos: Coords, grid_size: Coords) -> Bool {
  pos.x > 0 && pos.x < grid_size.x - 1 && pos.y > 0 && pos.y < grid_size.y - 1
}

fn is_inside2(pos: Coords, grid_size: Coords) -> Bool {
  pos.x > 1 && pos.x < grid_size.x - 2 && pos.y > 0 && pos.y < grid_size.y - 1
}

fn parse_grid(grid: String) -> #(Dict(Coords, Tile), Coords, Coords) {
  let lines =
    string.split(grid, "\n") |> list.filter(fn(line) { !string.is_empty(line) })

  let height = list.length(lines)
  let width = list.first(lines) |> result.unwrap("") |> string.length()
  let grid_size = Coords(width, height)

  let grid =
    list.index_fold(lines, #(dict.new(), Coords(0, 0)), fn(acc, line, y) {
      list.index_fold(string.split(line, ""), acc, fn(acc, char, x) {
        let coords = Coords(x, y)
        case char {
          "#" ->
            case is_inside(coords, grid_size) {
              True -> #(dict.insert(acc.0, coords, Wall), acc.1)
              False -> acc
            }
          "O" -> #(dict.insert(acc.0, coords, Box), acc.1)
          "@" -> #(acc.0, coords)
          _ -> acc
        }
      })
    })

  #(grid.0, grid.1, Coords(width, height))
}

fn parse_grid2(grid: String) -> #(Dict(Coords, Tile2), Coords, Coords) {
  let lines =
    string.split(grid, "\n") |> list.filter(fn(line) { !string.is_empty(line) })

  let height = list.length(lines)
  let width = list.first(lines) |> result.unwrap("") |> string.length()
  let grid_size = Coords(width, height)

  let grid =
    list.index_fold(lines, #(dict.new(), Coords(0, 0)), fn(acc, line, y) {
      list.index_fold(string.split(line, ""), acc, fn(acc, char, x) {
        let coords = Coords(x, y)
        case char {
          "#" ->
            case is_inside2(coords, grid_size) {
              True -> #(dict.insert(acc.0, coords, Wall2), acc.1)
              False -> acc
            }
          "[" -> #(dict.insert(acc.0, coords, LeftBox), acc.1)
          "]" -> #(dict.insert(acc.0, coords, RightBox), acc.1)
          "@" -> #(acc.0, coords)
          _ -> acc
        }
      })
    })

  #(grid.0, grid.1, Coords(width, height))
}

fn parse_moves(moves: List(String)) -> List(Move) {
  case moves {
    [">", ..rest] -> [Right, ..parse_moves(rest)]
    ["<", ..rest] -> [Left, ..parse_moves(rest)]
    ["^", ..rest] -> [Up, ..parse_moves(rest)]
    ["v", ..rest] -> [Down, ..parse_moves(rest)]
    _ -> []
  }
}

fn parse_input(input: String) -> Parsed {
  let sections =
    string.split(input, "\n\n")
    |> list.filter(fn(line) { !string.is_empty(line) })

  case sections {
    [grid, moves] -> {
      let #(tiles, robot, grid_size) = parse_grid(grid)
      let moves =
        parse_moves(
          string.split(moves, "")
          |> list.filter(fn(move) { move != "\n" }),
        )
      Parsed(tiles, robot, moves, grid_size)
    }
    _ -> Parsed(dict.new(), Coords(0, 0), [], Coords(0, 0))
  }
}

fn parse_input2(input: String) -> Parsed2 {
  let sections =
    string.split(input, "\n\n")
    |> list.filter(fn(line) { !string.is_empty(line) })

  case sections {
    [grid, moves] -> {
      let #(tiles, robot, grid_size) = parse_grid2(grid)
      let moves =
        parse_moves(
          string.split(moves, "")
          |> list.filter(fn(move) { move != "\n" }),
        )
      Parsed2(tiles, robot, moves, grid_size)
    }
    _ -> Parsed2(dict.new(), Coords(0, 0), [], Coords(0, 0))
  }
}

fn next_pos(pos: Coords, move: Move) -> Coords {
  let next_pos = case move {
    Up -> Coords(pos.x, pos.y - 1)
    Down -> Coords(pos.x, pos.y + 1)
    Left -> Coords(pos.x - 1, pos.y)
    Right -> Coords(pos.x + 1, pos.y)
  }
}

fn tile_at2(tiles: Dict(Coords, Tile2), pos: Coords, grid_size: Coords) -> Tile2 {
  case is_inside2(pos, grid_size) {
    True -> {
      case dict.get(tiles, pos) {
        Ok(tile) -> tile
        Error(_) -> Empty2
      }
    }
    False -> Wall2
  }
}

fn tile_at(tiles: Dict(Coords, Tile), pos: Coords, grid_size: Coords) -> Tile {
  case is_inside(pos, grid_size) {
    True -> {
      case dict.get(tiles, pos) {
        Ok(tile) -> tile
        Error(_) -> Empty
      }
    }
    False -> Wall
  }
}

fn make_move(
  tiles: Dict(Coords, Tile),
  pos: Coords,
  move: Move,
  grid_size: Coords,
) -> Result(Dict(Coords, Tile), Nil) {
  let next_pos = next_pos(pos, move)
  let next_tile = tile_at(tiles, next_pos, grid_size)
  case next_tile {
    Wall -> Error(Nil)
    Empty -> {
      let current_tile = tile_at(tiles, pos, grid_size)
      let next_tiles =
        tiles |> dict.delete(pos) |> dict.insert(next_pos, current_tile)
      Ok(next_tiles)
    }
    Box -> {
      let rec_tiles_res = make_move(tiles, next_pos, move, grid_size)
      case rec_tiles_res {
        Error(_) -> Error(Nil)
        Ok(rec_tiles) -> {
          let current_tile = tile_at(tiles, pos, grid_size)
          let next_tiles =
            rec_tiles |> dict.delete(pos) |> dict.insert(next_pos, current_tile)
          Ok(next_tiles)
        }
      }
    }
  }
}

fn simulate(
  tiles: Dict(Coords, Tile),
  robot: Coords,
  moves: List(Move),
  grid_size: Coords,
) -> Dict(Coords, Tile) {
  case moves {
    [] -> tiles
    [move, ..rest] -> {
      let new_tiles = make_move(tiles, robot, move, grid_size)
      case new_tiles {
        Error(_) -> {
          simulate(tiles, robot, rest, grid_size)
        }
        Ok(new_tiles) -> {
          let new_robot = next_pos(robot, move)
          simulate(new_tiles, new_robot, rest, grid_size)
        }
      }
    }
  }
}

fn make_move2(
  tiles: Dict(Coords, Tile2),
  pos: Coords,
  move: Move,
  grid_size: Coords,
) -> Result(Dict(Coords, Tile2), Nil) {
  let next_pos = next_pos(pos, move)
  let next_tile = tile_at2(tiles, next_pos, grid_size)
  case next_tile, move {
    Wall2, _ -> Error(Nil)
    Empty2, _ -> {
      let current_tile = tile_at2(tiles, pos, grid_size)
      let next_tiles =
        tiles |> dict.delete(pos) |> dict.insert(next_pos, current_tile)
      Ok(next_tiles)
    }
    LeftBox, Right | LeftBox, Left -> {
      let rec_tiles_res = make_move2(tiles, next_pos, move, grid_size)
      case rec_tiles_res {
        Error(_) -> Error(Nil)
        Ok(rec_tiles) -> {
          let current_tile = tile_at2(tiles, pos, grid_size)
          let next_tiles =
            rec_tiles |> dict.delete(pos) |> dict.insert(next_pos, current_tile)
          Ok(next_tiles)
        }
      }
    }
    LeftBox, _ -> {
      let rec_tiles_res =
        make_move2(tiles, next_pos, move, grid_size)
        |> result.try(fn(rec_tiles) {
          make_move2(rec_tiles, Coords(pos.x + 1, pos.y), move, grid_size)
        })
      case rec_tiles_res {
        Error(_) -> Error(Nil)
        Ok(rec_tiles) -> {
          let current_tile = tile_at2(tiles, pos, grid_size)
          let next_tiles =
            rec_tiles |> dict.delete(pos) |> dict.insert(next_pos, current_tile)
          Ok(next_tiles)
        }
      }
    }
    RightBox, Left | RightBox, Right -> {
      let rec_tiles_res = make_move2(tiles, next_pos, move, grid_size)
      case rec_tiles_res {
        Error(_) -> Error(Nil)
        Ok(rec_tiles) -> {
          let current_tile = tile_at2(tiles, pos, grid_size)
          let next_tiles =
            rec_tiles |> dict.delete(pos) |> dict.insert(next_pos, current_tile)
          Ok(next_tiles)
        }
      }
    }
    RightBox, _ -> {
      let rec_tiles_res =
        make_move2(tiles, next_pos, move, grid_size)
        |> result.try(fn(rec_tiles) {
          make_move2(rec_tiles, Coords(pos.x - 1, pos.y), move, grid_size)
        })
      case rec_tiles_res {
        Error(_) -> Error(Nil)
        Ok(rec_tiles) -> {
          let current_tile = tile_at2(tiles, pos, grid_size)
          let next_tiles =
            rec_tiles |> dict.delete(pos) |> dict.insert(next_pos, current_tile)
          Ok(next_tiles)
        }
      }
    }
  }
}

fn simulate2(
  tiles: Dict(Coords, Tile2),
  robot: Coords,
  moves: List(Move),
  grid_size: Coords,
) -> Dict(Coords, Tile2) {
  io.println(grid_to_str2(tiles, robot, grid_size))
  case moves {
    [] -> tiles
    [move, ..rest] -> {
      io.debug(move)
      let new_tiles = make_move2(tiles, robot, move, grid_size)
      case new_tiles {
        Error(_) -> {
          io.debug("Error")
          simulate2(tiles, robot, rest, grid_size)
        }
        Ok(new_tiles) -> {
          io.debug("Moved")
          let new_robot = next_pos(robot, move)
          simulate2(new_tiles, new_robot, rest, grid_size)
        }
      }
    }
  }
}

fn grid_to_str2(
  grid: Dict(Coords, Tile2),
  pos: Coords,
  grid_size: Coords,
) -> String {
  list.range(0, grid_size.y - 1)
  |> list.map(fn(y) {
    list.range(0, grid_size.x - 1)
    |> list.map(fn(x) {
      let current_pos = Coords(x, y)
      case dict.get(grid, current_pos) {
        Ok(tile) ->
          case tile {
            Wall2 -> "#"
            LeftBox -> "["
            RightBox -> "]"
            Empty2 -> "@"
          }
        Error(_) -> {
          case current_pos == pos {
            True -> "@"
            False -> "."
          }
        }
      }
    })
    |> string.join("")
  })
  |> string.join("\n")
  |> string.append("\n\n")
}

fn grid_to_str(
  grid: Dict(Coords, Tile),
  pos: Coords,
  grid_size: Coords,
) -> String {
  list.range(0, grid_size.y - 1)
  |> list.map(fn(y) {
    list.range(0, grid_size.x - 1)
    |> list.map(fn(x) {
      let current_pos = Coords(x, y)
      case dict.get(grid, current_pos) {
        Ok(tile) ->
          case tile {
            Wall -> "#"
            Box -> "O"
            Empty -> "@"
          }
        Error(_) -> {
          case current_pos == pos {
            True -> "@"
            False -> "."
          }
        }
      }
    })
    |> string.join("")
  })
  |> string.join("\n")
  |> string.append("\n\n")
}

fn expand_grid(input: String) -> String {
  input
  |> string.replace("#", "##")
  |> string.replace("O", "[]")
  |> string.replace(".", "..")
  |> string.replace("@.", "@.")
}

pub fn day15_solution() -> Solution {
  Solution(
    part1: fn(input: String) -> SolutionResult {
      let Parsed(tiles, robot, moves, grid_size) = parse_input(input)
      let new_tiles = simulate(tiles, robot, moves, grid_size)
      let score =
        new_tiles
        |> dict.filter(fn(_, tile) { tile == Box })
        |> dict.keys()
        |> list.map(fn(pos) { 100 * pos.y + pos.x })
        |> list.fold(0, fn(acc, score) { acc + score })

      IntResult(score)
    },
    part2: fn(input: String) -> SolutionResult {
      let Parsed2(tiles, robot, moves, grid_size) =
        parse_input2(expand_grid(input))
      let new_tiles = simulate2(tiles, robot, moves, grid_size)
      let score =
        new_tiles
        |> dict.filter(fn(_, tile) { tile == LeftBox })
        |> dict.keys()
        |> list.map(fn(pos) { 100 * pos.y + pos.x })
        |> list.fold(0, fn(acc, score) { acc + score })

      IntResult(score)
    },
  )
}
