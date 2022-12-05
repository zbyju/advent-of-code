import aoc.day10.Part1
import aoc.day10.Part2
import org.scalatest.FunSuite

class Day10Test extends FunSuite {
  test("Part 1 easy test") {
    val sol = Part1("/day10/input01.txt")
    assert(sol.solve() == 35)
  }

  test("Part 1 easy test 2") {
    val sol = Part1("/day10/input02.txt")
    assert(sol.solve() == 220)
  }

  test("Part 2 easy test") {
    val sol = Part2("/day10/input01.txt")
    assert(sol.solve() == 8)
  }

  test("Part 2 easy test 2") {
    val sol = Part2("/day10/input02.txt")
    assert(sol.solve() == 19208)
  }
}
