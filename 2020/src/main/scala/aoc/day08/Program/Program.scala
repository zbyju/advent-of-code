package aoc.day08.Program

import aoc.day08.Instructions.{Context, Instruction, InstructionFactory}

case class Program(lines: Array[String], execution: Execution) {
  var ctx: Context = Context()
  val instructFct = new InstructionFactory()
  val program: Array[Instruction] = {
    lines.map(line => instructFct(line))
  }

  def apply(): Int = {
    execution.apply(this)
  }
}
