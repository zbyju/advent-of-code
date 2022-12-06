import aoc.day11.Part1
import aoc.day11.Part2
import org.scalatest.FunSuite
import aoc.CommonHelper

class Day11Test extends FunSuite {
  test("Part 1 - test generating next plans") {
    for (i <- 1 to 5) {
      val prevFile = CommonHelper.getInputFile(f"/day11/input0${i}.txt")
      val prevPlan =
        CommonHelper.getInputLines(prevFile).toSeq.map(_.toCharArray().toSeq)
      val nextFile = CommonHelper.getInputFile(f"/day11/input0${i + 1}.txt")
      val nextPlan =
        CommonHelper.getInputLines(nextFile).toSeq.map(_.toCharArray().toSeq)
      val sol = Part1.nextPlan(prevPlan)
      assert(sol == nextPlan)
    }
  }

  test("Part 1 - example test") {
    val sol = Part1("/day11/input01.txt")
    assert(sol.solve() == 37)
  }

  test("Part 2 - test generating next plans") {
    for (i <- 1 to 6) {
      val prevFile = CommonHelper.getInputFile(f"/day11/input1${i}.txt")
      val prevPlan =
        CommonHelper.getInputLines(prevFile).toSeq.map(_.toCharArray().toSeq)
      val nextFile = CommonHelper.getInputFile(f"/day11/input1${i + 1}.txt")
      val nextPlan =
        CommonHelper.getInputLines(nextFile).toSeq.map(_.toCharArray().toSeq)
      val sol = Part2.nextPlan(prevPlan)
      assert(sol == nextPlan)
    }
  }

  test("Part 2 - example test") {
    val sol = Part2("/day11/input01.txt")
    assert(sol.solve() == 26)
  }
}
