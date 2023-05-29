import org.scalatest.FunSuite
import aoc.day15.Part1
import aoc.day15.Part2

class Day15Test extends FunSuite {
  test("Part 1 easy test") {
    val sol = Part1("/day15/input01.txt")
    assert(sol.solve() == 436)
  }
  test("Part 2 easy test") {
    val sol = Part2("/day15/input01.txt")
    assert(sol.solve() == 175594)
  }
}
