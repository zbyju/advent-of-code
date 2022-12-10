package aoc.day02

import aoc.Solution
import aoc.CommonHelper

case class Part1(inputPath: String) extends Solution(inputPath) {

  private def parsePassword(str: String): (Int, Int, Char, String) = {
    var tmp = str.split('-')
    val min: Int = tmp(0).toInt
    var str2 = tmp(1)

    tmp = str2.split(" ")
    val max: Int = tmp(0).toInt
    str2 = tmp(1) + tmp(2)

    tmp = str2.split(":")
    val param: String = tmp(0)
    val password: String = tmp(1).trim

    (min, max, param(0), password)
  }

  private def ruleIsOk(
      min: Int,
      max: Int,
      param: Char,
      password: String
  ): Boolean = {
    val count = password.count(_ == param)
    count >= min && count <= max
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

object Part1 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part1("/day02/part1.txt")
    val result = sol.solve()
    println(s"Day 02 - Part 1 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
