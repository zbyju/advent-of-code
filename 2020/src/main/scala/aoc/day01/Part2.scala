package aoc.day01

import aoc.{CommonHelper, Solution}
import scala.collection.Searching

case class Part2(inputPath: String) extends Solution(inputPath) {
  override def solve(): Int = {
    val expenses = lines.map(_.toInt).sorted

    for (i <- 0 until expenses.length - 2) {
      for (j <- i + 1 until expenses.length - 1) {
        val searchSpace = expenses.slice(j + 1, expenses.length)
        val searchingFor = 2020 - (expenses(i) + expenses(j))
        searchSpace.search(searchingFor) match {
          case Searching.Found(k) =>
            return expenses(i) * expenses(j) * searchSpace(k)
          case _ =>
        }
      }
    }
    return -1
  }
}

object Part2 {
  def main(args: Array[String]): Unit = {
    val sol = Part2("/day01/part2.txt")
    val result = sol.solve()
    println(s"Day 01 - Part 2 - result: $result")
  }
}
