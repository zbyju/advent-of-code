package aoc

object PerformanceRunner {

  def main(args: Array[String]): Unit = {
    val sols: Seq[() => Double] =
      Seq(
        day00.Solution.run _,
        day01.Part1.run _,
        day01.Part2.run _,
        day02.Part1.run _,
        day02.Part2.run _,
        day03.Part1.run _,
        day03.Part2.run _,
        day04.Part1.run _,
        day04.Part2.run _,
        day05.Part1.run _,
        day05.Part2.run _,
        day06.Part1.run _,
        day06.Part2.run _,
        day07.Part1.run _,
        day07.Part2.run _,
        day08.Part1.run _,
        day09.Part1.run _,
        day09.Part2.run _,
        day10.Part1.run _,
        day10.Part2.run _,
        day11.Part1.run _,
        day11.Part2.run _,
        day12.Part1.run _,
        day12.Part2.run _,
        day13.Part1.run _,
        day13.Part2.run _,
        day14.Part1.run _,
        day14.Part2.run _,
        day15.Part1.run _,
        day15.Part2.run _,
        day16.Part1.run _,
        day16.Part2.run _,
        day17.Part1.run _,
        day17.Part2.run _
      )
    val overallTime: Double = sols
      .map(sol => {
        val time = sol()
        println("Execution time: " + time)
        time
      })
      .sum
    println(overallTime + "ms")
  }
}
