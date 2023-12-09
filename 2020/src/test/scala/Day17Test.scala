import org.scalatest.FunSuite
import aoc.day17.Part1
import aoc.day17.Part2

class Day17Test extends FunSuite {
  test("Part 1 easy test") {
    val sol = Part1("/day17/input01.txt")
    assert(sol.solve() == 112)
  }
}
