package aoc.day05

import aoc.Solution

case class Part2(inputPath: String) extends Solution(inputPath) {
  private def calcSeatID(row: Int, col: Int): Int = row * 8 + col

  def getSeatID(line: String): Int = {
    var row = Range(0, 127)
    var col = Range(0, 7)

    for (c <- line) {
      c match {
        case 'F' => row = row.lowerHalf()
        case 'B' => row = row.upperHalf()
        case 'R' => col = col.upperHalf()
        case 'L' => col = col.lowerHalf()
      }
    }
    calcSeatID(row.max, col.max)
  }

  override def solve(): Int = {
    var max = 0
    val allSeats: scala.collection.mutable.Set[Int] =
      scala.collection.mutable.Set()
    for (line <- lines) {
      val tmp = getSeatID(line)
      if (tmp > max) max = tmp
      allSeats.addOne(tmp)
    }
    var result = -1
    for (i <- 0 until max) {
      if (
        !allSeats.contains(i) && allSeats
          .contains(i - 1) && allSeats.contains(i + 1)
      ) result = i
    }
    result
  }
}

object Part2 {
  def main(args: Array[String]): Unit = {
    val sol = Part2("/day05/part2.txt")
    val result = sol.solve()
    println(s"Day 05 - Part 2 - result: $result")
  }
}
