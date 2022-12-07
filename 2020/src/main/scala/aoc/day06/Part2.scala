package aoc.day06

import aoc.Solution

case class Part2(inputPath: String) extends Solution(inputPath) {

  def countAnswers(group: String): Int = {
    ('a' to 'z')
      .map(letter =>
        group
          .split("\n")
          .forall(person => person.contains(letter))
      )
      .map(b => if (b) 1 else 0)
      .sum
  }

  override def solve(): Int = {
    val groups = lines.mkString("\n").split("\n\n")
    groups.map(countAnswers).sum
  }
}

object Part2 {
  def main(args: Array[String]): Unit = {
    val sol = Part2("/day06/part1.txt")
    val result = sol.solve()
    println(s"Day 06 - Part 2 - result: $result")
  }
}
