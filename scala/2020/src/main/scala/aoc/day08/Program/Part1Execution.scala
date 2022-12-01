package aoc.day08.Program

case class Part1Execution() extends Execution() {
  override def apply(program: Program): Int = {
    var shouldContinue = true
    while(shouldContinue) {
      val instructionToExecute = program.program(program.ctx.programCounter)
      if(instructionToExecute.calledCount == 0) {
        val (newOp, newCtx) = instructionToExecute.apply(program.ctx)
        program.program(program.ctx.programCounter) = newOp
        program.ctx = newCtx
      } else {
        shouldContinue = false
      }
    }
    program.ctx.accumulator
  }
}
