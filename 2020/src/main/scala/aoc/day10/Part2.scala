package aoc.day10

import aoc.Solution

import scala.collection.mutable.HashMap

case class Part2(inputPath: String) extends Solution(inputPath) {
  val memo = HashMap[Seq[Int], Long]()

  def countPossibilities(adapters: Seq[Int]): Long = {
    if (adapters.isEmpty) return 0
    if (adapters.length == 1) return 1
    val memorized = memo.get(adapters)
    if (memorized.isDefined) return memorized.get
    val next = adapters.filter(x => x > adapters.head && x <= adapters.head + 3)
    val res =
      next.map(n => countPossibilities(adapters.filter(a => a >= n))).sum
    memo(adapters) = res
    res
  }
  override def solve(): Long = {
    val adapters = lines.map(_.toInt)
    val max = adapters.max
    countPossibilities(adapters.prepended(0).appended(max + 3).sorted)
  }
}

object Part2 {
  def main(args: Array[String]): Unit = {
    val sol = Part2("/day10/part1.txt")
    val result = sol.solve()
    println(s"The result is: $result")
  }
}
