package aoc

import scala.io.{BufferedSource, Source}

object CommonHelper {
  def getInputFile(path: String): BufferedSource = {
    Source.fromURL(getClass.getResource(path))
  }

  def getInputLines(bufferedSource: BufferedSource): Iterator[String] = {
    bufferedSource.getLines()
  }

  def nanoTime(from: Long, to: Long): Double = {
    (to - from).toDouble / 1_000_000
  }
}
