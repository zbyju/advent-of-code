import aoc.day09.Part1
import aoc.day09.Part2
import org.scalatest.FunSuite

class Day09Test extends FunSuite {
  test("Part 1 easy test") {
    val sol = Part1("/day09/input01.txt", 5)
    assert(sol.solve() == 127)
  }

  test("Part 2 easy test") {
    val sol = Part2("/day09/input01.txt", 5)
    assert(sol.solve() == 62)
  }
}
