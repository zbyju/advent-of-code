package aoc.day08.Instructions

abstract class Instruction(val name: String,
                           val argument: Any,
                           val calledCount: Int) {
  def apply(ctx: Context): (Instruction, Context) = ???
}