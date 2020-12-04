package aoc

abstract class Solution(inputPath: String) {
  val file = CommonHelper.getInputFile(inputPath)
  val linesIt: Iterator[String] = CommonHelper.getInputLines(file)
  val lines : Array[String] = linesIt.toArray

  def solve(): Any = ???
}
