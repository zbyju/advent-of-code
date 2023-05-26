package aoc.day12

import aoc.Solution
import aoc.CommonHelper

case class Part1(inputPath: String) extends Solution(inputPath) {
  override def solve(): Int = {
    val res = this.lines
      .map(str => Moves.parseStep(str))
      .foldLeft((0, 0, "E"))((acc, mv: Step) => {
        println(acc, mv, mv.move(acc))
        mv.move(acc)
      })
    Math.abs(res._1) + Math.abs(res._2)
  }
}

object Part1 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part1("/day12/part1.txt")
    val result = sol.solve()
    println(s"Day 12 - Part 1 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
