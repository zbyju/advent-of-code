package aoc

abstract class Solution(inputPath: String) {
  val linesIt: Iterator[String] = CommonHelper.getInputLines(inputPath)
  val lines : Array[String] = linesIt.toArray

  def solve(): Any = ???
}
