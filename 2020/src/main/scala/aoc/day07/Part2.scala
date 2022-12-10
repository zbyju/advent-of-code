package aoc.day07

import aoc.Solution

import scala.language.implicitConversions
import aoc.CommonHelper

case class Part2(inputPath: String) extends Solution(inputPath) {
  def printBags(bags: Seq[Bag]): Unit = {
    bags.foreach(bag => println(bag))
  }
  override def solve(): Int = {
    val bags = Bag(lines)
    val golden = bags.find(bag => bag.name == "shiny gold").get
    golden.hasToContain
  }
}

object Part2 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part2("/day07/part1.txt")
    val result = sol.solve()
    println(s"Day 07 - Part 2 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
