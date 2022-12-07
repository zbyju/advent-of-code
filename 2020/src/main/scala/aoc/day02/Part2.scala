package aoc.day02

import aoc.Solution

case class Part2(inputPath: String) extends Solution(inputPath) {

  private def parsePassword(str: String): (Int, Int, Char, String) = {
    var tmp = str.split('-')
    val index1: Int = tmp(0).toInt
    var str2 = tmp(1)

    tmp = str2.split(" ")
    val index2: Int = tmp(0).toInt
    str2 = tmp(1) + tmp(2)

    tmp = str2.split(":")
    val param: String = tmp(0)
    val password: String = tmp(1).trim

    (index1 - 1, index2 - 1, param(0), password)
  }

  private def ruleIsOk(
      index1: Int,
      index2: Int,
      param: Char,
      password: String
  ): Boolean = {
    password(index1) == param ^ password(index2) == param
  }

  override def solve(): Int = {
    var result = 0
    for (passwordEntry <- lines) {
      val (min, max, param, password) = parsePassword(passwordEntry)
      if (ruleIsOk(min, max, param, password)) result += 1
    }
    result
  }
}

object Part2 {
  def main(args: Array[String]): Unit = {
    val sol = Part2("/day02/part2.txt")
    val result = sol.solve()
    println(s"Day 02 - Part 2 - result: $result")
  }
}
