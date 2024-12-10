import argv
import gleam/int
import gleam/io
import gleam/list
import gleam/option.{Some}
import gleam/result
import runner

fn parse_int(args: List(String)) -> Result(#(Int, List(String)), String) {
  let first = list.first(args) |> result.try(fn(s) { int.parse(s) })
  let rest = list.rest(args)

  case first, rest {
    Ok(i), Ok(rest) -> Ok(#(i, rest))
    _, _ -> Error("Invalid arguments")
  }
}

pub fn main() {
  let args = argv.load().arguments
  let default_input_str = "input.txt"
  let default_input = Some(default_input_str)

  case list.length(args) {
    0 -> {
      let _ = runner.run_all_solutions(default_input)
      Nil
    }
    1 -> {
      let day = parse_int(args)
      case day {
        Ok(#(day, _)) -> {
          let _ = runner.run_solution(day, 1, default_input)
          let _ = runner.run_solution(day, 2, default_input)
          Nil
        }
        Error(_) -> io.println("Invalid arguments")
      }

      Nil
    }
    2 -> {
      let day = parse_int(args)
      let rest = result.unwrap(day, #(-1, list.new())).1
      let part = parse_int(rest)
      case day, part {
        Ok(#(day, _)), Ok(#(part, _)) -> {
          let _ = runner.run_solution(day, part, default_input)
          Nil
        }
        _, _ -> io.println("Invalid arguments")
      }
      Nil
    }
    3 -> {
      let day = parse_int(args)
      let rest = result.unwrap(day, #(-1, list.new())).1
      let part = parse_int(rest)
      let input = list.last(rest) |> result.unwrap(default_input_str)
      case day, part {
        Ok(#(day, _)), Ok(#(part, _)) -> {
          let _ = runner.run_solution(day, part, Some(input))
          Nil
        }
        _, _ -> io.println("Invalid arguments")
      }
      Nil
    }
    _ -> {
      io.println("Invalid arguments")
      Nil
    }
  }
}
