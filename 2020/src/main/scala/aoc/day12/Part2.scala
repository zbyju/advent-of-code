package aoc.day12

import aoc.Solution
import aoc.CommonHelper

case class Part2(inputPath: String) extends Solution(inputPath) {
  override def solve(): Int = {
    val res = this.lines
      .map(str => Moves2.parseStep(str))
      .foldLeft(State2((1, 10), (0, 0, "E")))((acc, mv: Step2) => mv.move(acc))
    Math.abs(res.sh._1) + Math.abs(res.sh._2)
  }
}

object Part2 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part2("/day12/part1.txt")
    val result = sol.solve()
    println(s"Day 12 - Part 2 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
