package aoc.day10

import aoc.Solution

import scala.language.implicitConversions
import scala.collection.immutable.Queue
import scala.util.Try

case class Part2(inputPath: String) extends Solution(inputPath) {
  def countPossibilities(adapters: Seq[Int]): Int = {
    if (adapters.isEmpty) return 0
    val next = adapters.filter(x => x > adapters.head && x <= adapters.head + 3)
    val nextPossibilities = countPossibilities(
      adapters.tail.filter(x => !next.contains(x))
    )
    return next.length + nextPossibilities
  }
  override def solve(): Long = {
    val adapters = lines.map(_.toInt)
    val max = adapters.max
    countPossibilities(adapters.prepended(0).appended(max).sorted)
  }
}

object Part2 {
  def main(args: Array[String]): Unit = {
    val sol = Part2("/day10/part1.txt")
    val result = sol.solve()
    println(s"The result is: $result")
  }
}
