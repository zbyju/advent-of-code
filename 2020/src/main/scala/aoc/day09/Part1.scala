package aoc.day09

import aoc.Solution

import scala.language.implicitConversions
import scala.collection.immutable.Queue
import scala.util.Try

case class Part1(inputPath: String, preambuleLength: Int)
    extends Solution(inputPath) {

  def isSumOfPrev(q: Seq[Long], num: Long): Boolean = {
    for (i <- (0 until q.length - 1)) {
      for (j <- (i + 1 until q.length)) {
        if (q(i) + q(j) == num) return true
      }
    }
    false
  }

  def recSolve(nums: Seq[Long], q: Seq[Long]): Long = {
    if (nums.length == 0) return -1
    if (!isSumOfPrev(q, nums(0))) return nums(0)
    recSolve(nums.drop(1), q.drop(1) :+ nums(0))
  }

  override def solve(): Long = {
    val nums = lines.flatMap(x => Try(x.toLong).toOption)
    recSolve(nums.drop(preambuleLength), nums.take(preambuleLength))
  }
}

object Part1 {
  def main(args: Array[String]): Unit = {
    val sol = Part1("/day09/part1.txt", 25)
    val result = sol.solve()
    println(s"Day 09 - Part 1 - result: $result")
  }
}
