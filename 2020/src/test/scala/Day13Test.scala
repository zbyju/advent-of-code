import org.scalatest.FunSuite
import aoc.day13.Part1
import aoc.day13.Part2

class Day13Test extends FunSuite {
  test("Part 1 easy test") {
    val sol = Part1("/day13/input01.txt")
    assert(sol.solve() == 295)
  }
  test("Part 2 easy test") {
    val sol = Part2("/day13/input01.txt")
    assert(sol.solve() == 1068781)
  }
}
