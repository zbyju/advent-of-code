package aoc.day01

import aoc.{CommonHelper, Solution}

case class Part2(inputPath : String) extends Solution(inputPath) {
  override def solve(): Int = {
    val results = for {
      x <- lines
      y <- lines
      z <- lines if x.toInt + y.toInt + z.toInt == 2020
    } yield (x, y, z)

    val result = results(0)
    result._1.toInt * result._2.toInt * result._3.toInt
  }
}

object Part2 {
  def main(args: Array[String]): Unit = {
    val sol = Part2("/day01/part2.txt")
    val result = sol.solve()
    println(s"The result is: $result")
  }
}


