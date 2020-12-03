import aoc.day02.{Part1, Part2}
import org.scalatest.{FunSuite, PrivateMethodTester}

class Day02Test extends FunSuite {

  test("Solution part 1") {
    val sol = Part1("/day02/input01.txt")
    assert(sol.solve() == 2)
  }

  test("Solution part 2") {
    val sol = Part2("/day02/input01.txt")
    assert(sol.solve() == 1)
  }

}
