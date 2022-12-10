package aoc.day01

import aoc.Solution
import scala.collection.Searching

case class Part1(inputPath: String) extends Solution(inputPath) {
  override def solve(): Int = {
    val expenses = lines.map(_.toInt).sorted

    for (i <- 0 until expenses.length - 1) {
      val searchSpace = expenses.slice(i + 1, expenses.length)
      val searchingFor = 2020 - expenses(i)
      searchSpace.search(searchingFor) match {
        case Searching.Found(j) => return expenses(i) * searchSpace(j)
        case _                  =>
      }

    }
    return -1
  }
}

object Part1 {
  def main(args: Array[String]): Unit = {
    val sol = Part1("/day01/part1.txt")
    val result = sol.solve()
    println(s"Day 01 - Part 1 - result: $result")
  }
}
