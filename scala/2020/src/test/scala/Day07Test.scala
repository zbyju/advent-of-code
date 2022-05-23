import aoc.day07.{Bag, Part1, Part2, Rule}
import org.scalatest.FunSuite

class Day07Test extends FunSuite {

  test("Bag canContain direct bag") {
    val goldenBag = Bag("golden bag")
    val mainBag = Bag("bright white")
    mainBag.addRule(Rule(goldenBag, 1))
    assert(mainBag.canContain("golden bag"))
  }

  test("Bag canContain bag depth 1") {
    val goldenBag = Bag("golden bag")
    val redBag = Bag("red bag")
    redBag.addRule(Rule(goldenBag, 1))
    val mainBag = Bag("bright white")
    mainBag.addRule(Rule(redBag, 1))
    assert(mainBag.canContain("golden bag"))
  }

  test("Part 1 easy test") {
    val sol = Part1("/day07/input01.txt")
    assert(sol.solve() == 4)
  }

  test("Part 1 final test") {
    val sol = Part1("/day07/final.txt")
    assert(sol.solve() == 302)
  }

  test("Part 2 easy test") {
    val sol = Part2("/day07/input01.txt")
    assert(sol.solve() == 32)
  }

  test("Part 2 easy test - 2") {
    val sol = Part2("/day07/input02.txt")
    assert(sol.solve() == 126)
  }
}
