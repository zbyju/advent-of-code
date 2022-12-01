package aoc.day01

import aoc.Solution

case class Part1(inputPath: String) extends Solution(inputPath) {
  override def solve(): Int = {
    val results = for {
      x <- lines;
      y <- lines if x.toInt + y.toInt == 2020
    } yield (x, y)

    val result = results(0)
    result._1.toInt * result._2.toInt
  }
}

object Part1 {
  def main(args: Array[String]): Unit = {
    val sol = Part1("/day01/part1.txt")
    val result = sol.solve()
    println(s"The result is: $result")
  }
}
