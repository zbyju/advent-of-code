import aoc.day01.{Part1, Part2}
import org.scalatest.FunSuite

class Day01Test extends FunSuite {

  test("Solution part 1") {
    val sol = Part1("/day01/input01.txt")
    assert(sol.solve() == 514579)
  }

  test("Solution part 2") {
    val sol = Part2("/day01/input01.txt")
    assert(sol.solve() == 241861950)
  }

}
