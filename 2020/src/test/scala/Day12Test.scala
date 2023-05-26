import org.scalatest.FunSuite
import aoc.day12.Part1

class Day12Test extends FunSuite {
  test("Part 1 easy test") {
    val sol = Part1("/day12/input01.txt")
    assert(sol.solve() == 25)
  }
}
