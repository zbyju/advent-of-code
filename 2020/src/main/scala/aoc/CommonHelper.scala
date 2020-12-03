package aoc

import scala.io.{BufferedSource, Source}

object CommonHelper {
  def getInputFile(path: String): BufferedSource = {
    Source.fromURL(getClass.getResource(path))
  }

  def getInputLines(path: String): Iterator[String] = {
    getInputFile(path).getLines()
  }
}
