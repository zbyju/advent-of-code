package aoc.day08.Instructions

case class ACC(override val argument: Int,
               override val calledCount: Int) extends Instruction("acc",argument, calledCount) {
  override def apply(ctx: Context): (ACC, Context) = {
    val newCtx = ctx.incrementPC.incrementACC(argument)
    (ACC(argument, calledCount + 1), newCtx)
  }
}
