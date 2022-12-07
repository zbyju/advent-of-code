package aoc.day03

import aoc.Solution

case class Part2(inputPath: String) extends Solution(inputPath) {
  override def solve(): Int = {
    val slopes = Array(
      (1, 1),
      (3, 1),
      (5, 1),
      (7, 1),
      (1, 2)
    )

    val width = lines(0).length
    val height = lines.length
    var result = 1

    for (slope <- slopes) {
      val positions =
        for (i <- slope._2 until height by slope._2)
          yield (i, (i / slope._2) * slope._1 % width)

      var treeCount = 0
      for (pos <- positions) {
        if (lines(pos._1)(pos._2) == '#') {
          treeCount += 1
        }
      }
      result *= treeCount
    }
    result
  }
}

object Part2 {
  def main(args: Array[String]): Unit = {
    val sol = Part2("/day03/part2.txt")
    val result = sol.solve()
    println(s"Day 03 - Part 2 - result: $result")
  }
}
