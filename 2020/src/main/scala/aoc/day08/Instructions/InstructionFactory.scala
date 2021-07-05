package aoc.day08.Instructions

class InstructionFactory() {
  def apply(line: String): Instruction = {
    val InstructionRegex = "^(\\w\\w\\w) ([+-])(\\d+)$".r
    line match {
      case InstructionRegex(name, sign, argument) => {
        val arg = if(sign == "+") argument.toInt else -argument.toInt
        name.toLowerCase match {
          case "nop" => NOP(arg, 0)
          case "acc" => ACC(arg, 0)
          case "jmp" => JMP(arg, 0)
          case _ => ERR(0, 0)
        }
      }
      case _ => ERR(0, 0)
    }
  }
}
