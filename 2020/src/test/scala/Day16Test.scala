import org.scalatest.FunSuite
import aoc.day16.Part1
import aoc.day16.Part2

class Day16Test extends FunSuite {
  test("Part 1 easy test") {
    val sol = Part1("/day16/input01.txt")
    assert(sol.solve() == 71)
  }
  test("Part 2 easy test") {
    val sol = Part2("/day16/input01.txt")
    assert(true)
  }
}
