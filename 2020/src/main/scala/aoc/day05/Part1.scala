package aoc.day05

import aoc.Solution
import aoc.CommonHelper

case class Part1(inputPath: String) extends Solution(inputPath) {
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
    for (line <- lines) {
      val tmp = getSeatID(line)
      if (tmp > max) max = tmp
    }
    max
  }
}

object Part1 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part1("/day05/part1.txt")
    val result = sol.solve()
    println(s"Day 06 - Part 1 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
