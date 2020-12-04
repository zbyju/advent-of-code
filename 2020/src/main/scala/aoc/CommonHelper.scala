package aoc

import scala.io.{BufferedSource, Source}

object CommonHelper {
  def getInputFile(path: String): BufferedSource = {
    Source.fromURL(getClass.getResource(path))
  }

  def getInputLines(bufferedSource: BufferedSource): Iterator[String] = {
    bufferedSource.getLines()
  }
}
