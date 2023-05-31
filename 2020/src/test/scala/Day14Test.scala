import org.scalatest.FunSuite
import aoc.day14.Part1
import aoc.day14.Part2

class Day14Test extends FunSuite {
  test("Part 1 easy test") {
    val sol = Part1("/day14/input01.txt")
    assert(sol.solve() == 165)
  }
  test("Part 2 easy test") {
    val sol = Part2("/day14/input02.txt")
    assert(sol.solve() == 208)
  }
}
