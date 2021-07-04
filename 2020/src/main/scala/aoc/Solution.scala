package aoc

import scala.io.BufferedSource

abstract class Solution(inputPath: String) {
  val file: BufferedSource = CommonHelper.getInputFile(inputPath)
  val linesIt: Iterator[String] = CommonHelper.getInputLines(file)
  val lines : Array[String] = linesIt.toArray

  def solve(): Any = ???
}
