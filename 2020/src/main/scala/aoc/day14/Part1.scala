package aoc.day14

import aoc.Solution
import aoc.CommonHelper

case class Part1(inputPath: String) extends Solution(inputPath) {
  override def solve(): Long = {
    val str: String = "0" * 36
    val (mem, _) = lines
      .foldLeft((Map.empty[Int, String]), str)((acc, line) =>
        line.take(3) match {
          case "mas" => (acc._1, line.split(" = ").last)
          case "mem" => {
            val mem = acc._1
            val mask = acc._2
            val adr = line.split(Array('[', ']'))(1).toInt
            val value =
              line
                .split(" = ")
                .last
                .toInt
                .toBinaryString
                .reverse
                .padTo(36, '0')
                .reverse
            val masked =
              mask
                .zip(value)
                .map { case (m, v) => if (m == 'X') v else m }
                .mkString
            (mem.updated(adr, masked), mask)
          }
        }
      )
    mem.toSeq.map { case (_, v) =>
      v.reverse
        .foldLeft(0L, 0)((acc, x) =>
          (acc._1 + Math.pow(2, acc._2).toLong * x.asDigit, acc._2 + 1)
        )
        ._1
    }.sum
  }
}

object Part1 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part1("/day14/part1.txt")
    val result = sol.solve()
    println(s"Day 14 - Part 1 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
