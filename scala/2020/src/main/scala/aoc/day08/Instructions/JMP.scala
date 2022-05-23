package aoc.day08.Instructions

case class JMP(override val argument: Int,
               override val calledCount: Int) extends Instruction("jmp", argument, calledCount) {
  override def apply(ctx: Context): (JMP, Context) = {
    val newCtx = ctx.incrementPC(argument)
    (JMP(argument, calledCount + 1), newCtx)
  }
}