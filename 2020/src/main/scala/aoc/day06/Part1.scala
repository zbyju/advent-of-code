package aoc.day06

import aoc.Solution
import aoc.CommonHelper

case class Part1(inputPath: String) extends Solution(inputPath) {

  def countAnswers(group: String): Int = {
    group.replaceAll("\n", "").groupBy(c => c).map(e => 1).sum
  }

  override def solve(): Int = {
    val groups = lines.mkString("\n").split("\n\n")
    groups.map(countAnswers).sum
  }
}

object Part1 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part1("/day06/part1.txt")
    val result = sol.solve()
    println(s"Day 06 - Part 1 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
