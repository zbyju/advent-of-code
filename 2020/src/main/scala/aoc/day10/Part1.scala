package aoc.day10

import aoc.Solution

import scala.language.implicitConversions
import scala.collection.immutable.Queue
import scala.util.Try
import aoc.CommonHelper

case class Part1(inputPath: String) extends Solution(inputPath) {
  type Diffs = (Int, Int)

  override def solve(): Int = {
    val adapters = lines.map(_.toInt)
    val max = adapters.max
    val diffs =
      adapters
        .prepended(0)
        .appended(max + 3)
        .sorted
        .sliding(2)
        .map(x => x.last - x.head)
        .foldLeft((0, 0))((sum, diff) =>
          diff match {
            case 1 => (sum._1 + 1, sum._2)
            case 3 => (sum._1, sum._2 + 1)
          }
        )
    diffs._1 * diffs._2
  }
}

object Part1 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part1("/day10/part1.txt")
    val result = sol.solve()
    println(s"Day 10 - Part 1 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
