package aoc.day00

import aoc.CommonHelper

object Solution {
  def run(): Double = {
    val from = System.nanoTime()
    println("Hello World!")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }

  def main(args: Array[String]): Unit = run()
}
