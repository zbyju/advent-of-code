package aoc.day08.Instructions

case class ERR(override val argument: Int = 0,
               override val calledCount: Int) extends Instruction("err", argument, calledCount) {
  override def apply(ctx: Context): (ERR, Context) = {
    (ERR(argument, calledCount + 1), ctx.errorCtx)
  }
}
