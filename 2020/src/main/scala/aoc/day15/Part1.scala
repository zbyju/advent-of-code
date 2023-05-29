package aoc.day15

import aoc.Solution
import aoc.CommonHelper

case class Part1(inputPath: String, rounds: Int = 2020)
    extends Solution(inputPath) {
  def initState(input: Array[Int]): (Map[Int, Int], Int) = {
    input
      .foldLeft((Map.empty[Int, Int], 1))((acc, x) => {
        (acc._1 + (x -> acc._2), acc._2 + 1)
      })
  }
  override def solve(): Int = {
    val input = this.lines.head.split(",").map(_.toInt)
    val (map, round) = initState(input.dropRight(1))
    ((round + 1) to rounds)
      .foldLeft((map, input.last))((acc, round) => {
        val roundSpoken = acc._1.getOrElse(acc._2, 0)
        val spoken = if (roundSpoken == 0) 0 else round - roundSpoken - 1
        (acc._1.updated(acc._2, (round - 1)), spoken)
      })
      ._2
  }
}

object Part1 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part1("/day15/part1.txt")
    val result = sol.solve()
    println(s"Day 15 - Part 1 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
