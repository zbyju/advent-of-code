package aoc.day13

import aoc.Solution
import aoc.CommonHelper

case class Part2(inputPath: String) extends Solution(inputPath) {
  override def solve(): BigInt = {
    var schedule = this.lines.last.split(',')
    val busIds = schedule.filter(_ != "x").map(BigInt(_)).toList
    val offsets = schedule.indices
      .filter(idx => schedule(idx) != "x")
      .map(BigInt(_))
      .toArray
    val t0 = (busIds zip offsets).map { case (n, a) => if (a > 0) n - a else a }
    val lcm = busIds.product

    @annotation.tailrec
    def chineseRemainder(
        ns: List[BigInt],
        as: List[BigInt],
        t: BigInt
    ): BigInt = {
      def mulInv(a: BigInt, b: BigInt): BigInt = {
        @annotation.tailrec
        def loop(a: BigInt, b: BigInt, x0: BigInt, x1: BigInt): BigInt = {
          if (a > 1) loop(b, a % b, x1 - (a / b) * x0, x0)
          else x1
        }
        if (b == 1) 1
        else {
          val x1 = loop(a, b, 0, 1)
          if (x1 < 0) x1 + b
          else x1
        }
      }

      if (ns.isEmpty) t
      else {
        val p = lcm / ns.head
        chineseRemainder(ns.tail, as.tail, t + as.head * mulInv(p, ns.head) * p)
      }
    }
    chineseRemainder(busIds, t0, 0) % lcm
  }
}

object Part2 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part2("/day13/part1.txt")
    val result = sol.solve()
    println(s"Day 13 - Part 2 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
