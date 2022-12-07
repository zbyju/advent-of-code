package aoc

object PerformanceRunner {
  def time(block: => Any): Double = {
    val v0 = System.nanoTime()
    block
    val v1 = System.nanoTime()
    val t = (v1 - v0).toDouble / 1_000_000
    println("Execution time: " + t)
    return t
  }

  def main(args: Array[String]): Unit = {
    val sols =
      Seq(
        day00.Solution.main _,
        day01.Part1.main _,
        day01.Part2.main _,
        day02.Part1.main _,
        day02.Part2.main _,
        day03.Part1.main _,
        day03.Part2.main _,
        day04.Part1.main _,
        day04.Part2.main _,
        day05.Part1.main _,
        day05.Part2.main _,
        day06.Part1.main _,
        day06.Part2.main _,
        day07.Part1.main _,
        day07.Part2.main _,
        day08.Part1.main _,
        day09.Part1.main _,
        day09.Part2.main _,
        day10.Part1.main _,
        day10.Part2.main _,
        day11.Part1.main _,
        day11.Part2.main _
      )
    val overallTime: Double = sols
      .map(sol => time(sol(Array())))
      .sum
    println(overallTime + "ms")
  }
}
