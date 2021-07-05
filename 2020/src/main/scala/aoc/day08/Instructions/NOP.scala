package aoc.day08.Instructions

case class NOP(override val argument: Int,
               override val calledCount: Int) extends Instruction("nop", argument, calledCount) {
  override def apply(ctx: Context): (NOP, Context) = {
    val newCtx = ctx.incrementPC
    (NOP(argument, calledCount + 1), newCtx)
  }
}
