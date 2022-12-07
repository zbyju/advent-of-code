package aoc.day11

import aoc.Solution

import scala.language.implicitConversions
import scala.collection.immutable.Queue
import scala.util.Try

case class Coords(row: Int, col: Int)

case class Part1(inputPath: String) extends Solution(inputPath) {

  override def solve(): Int = {
    val plan = lines.toSeq.map(_.toCharArray().toSeq)
    Part1.iterate(plan).foldLeft(0)((sum, row) => sum + row.count(_ == '#'))
  }
}

object Part1 {
  type Plan = Seq[Seq[Char]]

  def main(args: Array[String]): Unit = {
    val sol = Part1("/day11/part1.txt")
    val result = sol.solve()
    println(s"Day 11 - Part 1 - result: $result")
  }

  def getAdjacent(plan: Plan, coords: Coords): Seq[Char] = {
    val dirs =
      Seq((1, 0), (0, 1), (-1, 0), (0, -1), (1, 1), (-1, 1), (-1, -1), (1, -1))

    dirs
      .map(dir =>
        plan.lift(coords.row + dir._1).map(_.lift(coords.col + dir._2)).flatten
      )
      .flatten
  }

  def mapEmptySeat(adjacent: Seq[Char]): Char = {
    if (adjacent.forall(_ != '#')) '#' else 'L'
  }

  def mapOccupiedSeat(adjacent: Seq[Char]): Char = {
    if (adjacent.count(_ == '#') >= 4) 'L' else '#'
  }

  def nextPlan(plan: Plan): Plan = {
    plan.zipWithIndex.map({ case (row, r) =>
      row.zipWithIndex.map({ case (seat, c) =>
        seat match {
          case '#' => mapOccupiedSeat(getAdjacent(plan, Coords(r, c)))
          case 'L' => mapEmptySeat(getAdjacent(plan, Coords(r, c)))
          case x   => x
        }
      })
    })
  }

  def iterate(plan: Plan): Plan = {
    val newPlan = nextPlan(plan)
    if (plan == newPlan) return newPlan
    iterate(newPlan)
  }
}
