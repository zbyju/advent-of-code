import aoc.day08.Part1
import aoc.day08.Program.{Part1Execution, Program}
import org.scalatest.FunSuite

class Day08Test extends FunSuite {
    test("Program gets assembled") {
      val sol = Part1("/day08/input01.txt")
      val exec = Part1Execution()
      val program: Program = Program(sol.lines.mkString("\n").split("\n"), exec)
      assert(program.program.length == 9)
    }

    test("Program gets correctly assembled") {
      val sol = Part1("/day08/input01.txt")
      val exec = Part1Execution()
      val program: Program = Program(sol.lines.mkString("\n").split("\n"), exec)
      assert(program.program.length == 9)
      assert(program.program(0).name == "nop")
      assert(program.program(0).argument == 0)
      assert(program.program(1).name == "acc")
      assert(program.program(1).argument == 1)
      assert(program.program(2).name == "jmp")
      assert(program.program(2).argument == 4)
      assert(program.program(3).name == "acc")
      assert(program.program(3).argument == 3)
      assert(program.program(4).name == "jmp")
      assert(program.program(4).argument == -3)
      assert(program.program(5).name == "acc")
      assert(program.program(5).argument == -99)
      assert(program.program(6).name == "acc")
      assert(program.program(6).argument == 1)
      assert(program.program(7).name == "jmp")
      assert(program.program(7).argument == -4)
      assert(program.program(8).name == "acc")
      assert(program.program(8).argument == 6)
    }

    test("Part 1 easy test") {
      val sol = Part1("/day08/input01.txt")
      assert(sol.solve() == 5)
    }
}
