import aoc.day03.{Part1, Part2}
import org.scalatest.{FunSuite, PrivateMethodTester}

class Day03Test extends FunSuite {

  test("Solution part 1 - 1") {
    val sol = Part1("/day03/input01.txt")
    assert(sol.solve() == 7)
  }

  test("Solution part 1 - 2") {
    val sol = Part1("/day03/input02.txt")
    assert(sol.solve() == 7)
  }

  test("Solution part 2") {
    val sol = Part2("/day03/input02.txt")
    assert(sol.solve() == 336)
  }

}
