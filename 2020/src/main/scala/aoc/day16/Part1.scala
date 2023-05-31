package aoc.day16

import aoc.Solution
import aoc.CommonHelper

case class Part1(inputPath: String) extends Solution(inputPath) {
  def isInRange(r: (Int, Int), x: Int): Boolean = x >= r._1 && x <= r._2
  def isInRange2(rs: ((Int, Int), (Int, Int)), x: Int): Boolean =
    isInRange(rs._1, x) || isInRange(rs._2, x)
  override def solve(): Int = {
    val (ranges, _, tickets) = Parser.parse(this.lines.mkString("\n"))
    tickets.flatten
      .map(t => (ranges.forall(r => !isInRange2(r._2, t)), t))
      .filter(_._1 == true)
      .map(_._2)
      .sum
  }
}

object Part1 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part1("/day16/part1.txt")
    val result = sol.solve()
    println(s"Day 16 - Part 1 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
