import aoc.day04.{Part2, Part1}
import org.scalatest.FunSuite

class Day04Test extends FunSuite{
  test("Solution part 1") {
    val sol = Part1("/day04/input01.txt")
    assert(sol.solve() == 2)
  }

  test("Solution part 2 - All ok") {
    val sol = Part2("/day04/input03.txt")
    assert(sol.solve() == 4)
  }

  test("Solution part 2 - All wrong") {
    val sol = Part2("/day04/input04.txt")
    assert(sol.solve() == 0)
  }

  test("Solution part 2 - Mixed") {
    val sol = Part2("/day04/input02.txt")
    assert(sol.solve() == 4)
  }
}
