package aoc.day08

import aoc.Solution
import aoc.day08.Program.{Part1Execution, Program}
import aoc.CommonHelper

case class Part1(inputPath: String) extends Solution(inputPath) {
  override def solve(): Int = {
    val part1Exec = Part1Execution()
    val program: Program = Program(lines.mkString("\n").split("\n"), part1Exec)
    program()
  }
}

object Part1 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part1("/day08/part1.txt")
    val result = sol.solve()
    println(s"Day 08 - Part 1 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
