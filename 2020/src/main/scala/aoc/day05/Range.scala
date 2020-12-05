package aoc.day05

case class Range(min: Int, max: Int) {
  def length: Int = (max - min + 1)
  def half: Int = length / 2

  def lowerHalf(): Range = Range(min, max - half)
  def upperHalf(): Range = Range(min + half, max)
}
