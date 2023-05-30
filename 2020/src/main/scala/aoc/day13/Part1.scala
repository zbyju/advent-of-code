package aoc.day13

import aoc.Solution
import aoc.CommonHelper

case class Part1(inputPath: String) extends Solution(inputPath) {
  override def solve(): Int = {
    val earliest = lines.head.toInt
    val ids = lines.last.split(",").filter(x => x != "x").map(_.toInt)
    val (bestId, time) = ids.foldLeft(ids.head, Int.MaxValue)((acc, x) => {
      val timeBus = x * Math.ceil(earliest.toDouble / x.toDouble).toInt
      val waiting = timeBus - earliest
      if (waiting < acc._2) (x, waiting) else acc
    })
    bestId * time
  }
}

object Part1 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part1("/day13/part1.txt")
    val result = sol.solve()
    println(s"Day 13 - Part 1 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
