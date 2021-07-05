package aoc.day08.Instructions

case class Context(programCounter: Int = 0, accumulator: Int = 0) {
  def errorCtx: Context = {
    Context(-1, accumulator)
  }
  def incrementPC: Context = {
    Context(this.programCounter + 1, accumulator)
  }
  def incrementPC(num: Int): Context = {
    Context(this.programCounter + num, accumulator)
  }

  def incrementACC(num: Int): Context = {
    Context(this.programCounter, accumulator + num)
  }
}
