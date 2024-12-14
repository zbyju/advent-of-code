import aoc_solution.{type Solution, type SolutionResult, FloatResult, Solution}
import gleam/float
import gleam/int
import gleam/list
import gleam/result
import gleam/string

type Vector2 {
  Vector2(x: Float, y: Float)
}

type Machine {
  Machine(a: Vector2, b: Vector2, prize: Vector2)
}

fn parse_number(part: String) -> Float {
  string.drop_start(part, 2)
  |> int.parse()
  |> result.unwrap(0)
  |> int.to_float()
}

fn parse_line(line: String) -> Vector2 {
  let parts = string.split(line, ": ")
  let numbers = parts |> list.last() |> result.unwrap("") |> string.split(", ")
  case numbers {
    [x, y] -> Vector2(x: parse_number(x), y: parse_number(y))
    _ -> Vector2(x: 0.0, y: 0.0)
  }
}

fn parse_input(input: String) -> List(Machine) {
  string.split(input, "\n\n")
  |> list.filter(fn(line) { line != "" })
  |> list.map(fn(line) {
    let lines = string.split(line, "\n") |> list.filter(fn(line) { line != "" })
    case lines {
      [a, b, prize] -> {
        let a = parse_line(a)
        let b = parse_line(b)
        let prize = parse_line(prize)
        Machine(a: a, b: b, prize: prize)
      }
      _ ->
        Machine(
          a: Vector2(x: 0.0, y: 0.0),
          b: Vector2(x: 0.0, y: 0.0),
          prize: Vector2(x: 0.0, y: 0.0),
        )
    }
  })
  |> list.filter(fn(machine) {
    let Machine(a: a, b: b, prize: prize) = machine
    a.x != 0.0
    && a.y != 0.0
    && b.x != 0.0
    && b.y != 0.0
    && prize.x != 0.0
    && prize.y != 0.0
  })
}

fn is_int(f: Float) -> Bool {
  let i = float.round(f) |> int.to_float()
  float.loosely_equals(f, i, 0.001)
}

pub fn day13_solution() -> Solution {
  Solution(
    part1: fn(input: String) -> SolutionResult {
      let machines = parse_input(input)

      let result =
        machines
        |> list.map(fn(machine) {
          let ax = machine.a.x
          let ay = machine.a.y
          let bx = machine.b.x
          let by = machine.b.y
          let cx = machine.prize.x
          let cy = machine.prize.y

          let k = float.divide(ay, ax) |> result.unwrap(0.0)
          let b =
            float.divide(
              float.subtract(cy, float.multiply(cx, k)),
              float.subtract(by, float.multiply(bx, k)),
            )
            |> result.unwrap(0.0)
          let a =
            float.divide(float.subtract(cx, float.multiply(bx, b)), ax)
            |> result.unwrap(0.0)

          #(a, b)
        })
        |> list.filter(fn(a_b) {
          let #(a, b) = a_b
          is_int(a) && is_int(b)
        })
        |> list.fold(0.0, fn(acc, ab) {
          float.add(acc, float.multiply(ab.0, 3.0)) |> float.add(ab.1)
        })

      FloatResult(result)
    },
    part2: fn(input: String) -> SolutionResult {
      let machines = parse_input(input)

      let result =
        machines
        |> list.map(fn(machine) {
          let ax = machine.a.x
          let ay = machine.a.y
          let bx = machine.b.x
          let by = machine.b.y
          let cx = float.add(machine.prize.x, 10_000_000_000_000.0)
          let cy = float.add(machine.prize.y, 10_000_000_000_000.0)

          let k = float.divide(ay, ax) |> result.unwrap(0.0)
          let b =
            float.divide(
              float.subtract(cy, float.multiply(cx, k)),
              float.subtract(by, float.multiply(bx, k)),
            )
            |> result.unwrap(0.0)
          let a =
            float.divide(float.subtract(cx, float.multiply(bx, b)), ax)
            |> result.unwrap(0.0)

          #(a, b)
        })
        |> list.filter(fn(a_b) {
          let #(a, b) = a_b
          is_int(a) && is_int(b)
        })
        |> list.fold(0.0, fn(acc, ab) {
          float.add(acc, float.multiply(ab.0, 3.0)) |> float.add(ab.1)
        })

      FloatResult(result)
    },
  )
}
