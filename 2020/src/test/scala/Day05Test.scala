import aoc.day05.{Part1, Part2}
import org.scalatest.{FunSuite, PrivateMethodTester}

class Day05Test extends FunSuite {

  test("Boarding seat test 1") {
    val sol = Part1("/day05/input01.txt")
    assert(sol.getSeatID("BFFFBBFRRR") == 567)
  }
  test("Boarding seat test 2") {
    val sol = Part1("/day05/input01.txt")
    assert(sol.getSeatID("FFFBBBFRRR") == 119)
  }
  test("Boarding seat test 3") {
    val sol = Part1("/day05/input01.txt")
    assert(sol.getSeatID("BBFFBBFRLL") == 820)
  }


  test("Solution part 1") {
    val sol = Part1("/day05/input01.txt")
    assert(sol.solve() == 820)
  }
}
