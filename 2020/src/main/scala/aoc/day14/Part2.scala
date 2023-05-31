package aoc.day14

import aoc.Solution
import aoc.CommonHelper

case class Part2(inputPath: String) extends Solution(inputPath) {
  override def solve(): BigInt = {
    val str: String = "0" * 36
    val (mem, _) = lines
      .foldLeft((Map.empty[String, Int]), str)((acc, line) =>
        line.take(3) match {
          case "mas" => (acc._1, line.split(" = ").last)
          case "mem" => {
            val mem = acc._1
            val mask = acc._2
            val adr = line
              .split(Array('[', ']'))(1)
              .toInt
              .toBinaryString
              .reverse
              .padTo(36, '0')
              .reverse
              .zip(mask)
              .map { case (a, m) => if (m == '0') a else m }
            val value =
              line
                .split(" = ")
                .last
                .toInt
            val adrMasked =
              adr
                .foldLeft(List(""))((acc, x) =>
                  x match {
                    case 'X' =>
                      acc.flatMap(out =>
                        List(out.appended('0'), out.appended('1'))
                      )

                    case c => acc.map(out => out.appended(c))
                  }
                )

            val newMem =
              adrMasked.foldLeft(mem)((acc, a) =>
                acc.updated(a.mkString, value)
              )
            (newMem, mask)
          }
        }
      )
    mem.toSeq.map { case (_, v) => v }.sum
  }
}

object Part2 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part2("/day14/part1.txt")
    val result = sol.solve()
    println(s"Day 14 - Part 2 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
