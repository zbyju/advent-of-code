package aoc.day03

import aoc.Solution

case class Part1(inputPath : String) extends Solution(inputPath) {
  override def solve(): Int = {
    val width = lines(0).length
    val height = lines.length
    val positions = for (
      i <- 1 until height
    ) yield (i, (i * 3 % width))

    var treeCount = 0
    for(pos <- positions) {
      if(lines(pos._1)(pos._2) == '#') {
        treeCount += 1
      }
    }
    treeCount
  }
}

object Part1 {
  def main(args: Array[String]): Unit = {
    val sol = Part1("/day03/part1.txt")
    val result = sol.solve()
    println(s"The result is: $result")
  }
}
