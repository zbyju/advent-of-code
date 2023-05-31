package aoc.day16

import aoc.Solution
import aoc.CommonHelper

case class Part2(inputPath: String) extends Solution(inputPath) {
  def isInRange(r: (Int, Int), x: Int): Boolean = x >= r._1 && x <= r._2
  def isInRange2(rs: ((Int, Int), (Int, Int)), x: Int): Boolean =
    isInRange(rs._1, x) || isInRange(rs._2, x)
  val (ranges, ticket, tickets) = Parser.parse(lines.mkString("\n"))
  val valids = tickets
    .filter(t => !t.exists(i => (ranges.forall(r => !isInRange2(r._2, i)))))

  override def solve(): Long = {
    val ticketRule = "(\\D+): (\\d+)-(\\d+) or (\\d+)-(\\d+)".r
    val (ruleMap, tickets) =
      lines.foldLeft(
        (Map.empty[String, Int => Boolean], Vector.empty[Vector[Int]])
      ) { case (rules, tickets) -> line =>
        line match {
          case ticketRule(name, min1, max1, min2, max2) =>
            val rule = (a: Int) =>
              a >= min1.toInt && a <= max1.toInt || a >= min2.toInt && a <= max2.toInt
            (rules + (name -> rule), tickets)
          case str if str.contains(',') =>
            (rules, tickets :+ str.split(',').map(_.toInt).toVector)
          case _ => (rules, tickets)
        }
      }
    val guesses = valids
      .flatMap(_.zipWithIndex)
      .map(_.swap)
      .groupMap(_._1)(_._2)
      .map { case idx -> values =>
        idx -> ruleMap.foldLeft(Set.empty[String]) { case (out, (name, rule)) =>
          if (values.forall(rule)) out + name else out
        }
      }
      .map { case (idx, rules) => rules -> idx }

    @annotation.tailrec
    def deduce(
        facts: Map[String, Int],
        unknown: Map[Set[String], Int]
    ): Map[String, Int] =
      if (unknown.isEmpty) facts
      else {
        val deduced = unknown.map { case set -> idx =>
          (set diff facts.keySet) -> idx
        }
        deduce(
          facts ++ deduced.filter(_._1.size == 1).map { case (set, idx) =>
            set.head -> idx
          },
          deduced.filter(_._1.size > 1)
        )
      }

    deduce(Map.empty[String, Int], guesses)
      .collect {
        case (name, idx) if name.startsWith("departure") => ticket(idx)
      }
      .map(_.toLong)
      .product
  }
}

object Part2 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part2("/day16/part1.txt")
    val result = sol.solve()
    println(s"Day 16 - Part 2 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
