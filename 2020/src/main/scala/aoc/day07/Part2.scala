package aoc.day07

import aoc.Solution

import scala.language.implicitConversions

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
  def main(args: Array[String]): Unit = {
    val sol = Part2("/day07/part1.txt")
    val result = sol.solve()
    println(s"Day 07 - Part 2 - result: $result")
  }
}
