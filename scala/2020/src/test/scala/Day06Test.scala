import aoc.day06.{Part1, Part2}
import org.scalatest.FunSuite

class Day06Test extends FunSuite {

  test("Count unique letters test") {
    val sol = Part1("/day06/input01.txt")
    assert(sol.countAnswers(sol.lines.mkString) == 4)
  }

  test("Part 1 easy test") {
    val sol = Part1("/day06/input02.txt")
    assert(sol.solve() == 11)
  }

  test("Part 2 easy test") {
    val sol = Part2("/day06/input02.txt")
    assert(sol.solve() == 6)
  }

  test("Part 1 full test") {
    val sol = Part1("/day06/input03.txt")
    assert(sol.solve() == 6437)
  }

  test("Part 2 full test") {
    val sol = Part2("/day06/input03.txt")
    assert(sol.solve() == 3229)
  }
}
