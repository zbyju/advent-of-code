package aoc.day11

import aoc.Solution

import scala.collection.mutable.HashMap
import aoc.CommonHelper

case class Part2(inputPath: String) extends Solution(inputPath) {

  override def solve(): Int = {
    val plan = lines.toSeq.map(_.toCharArray().toSeq)
    Part2.iterate(plan).foldLeft(0)((sum, row) => sum + row.count(_ == '#'))
  }
}

object Part2 {
  type Plan = Seq[Seq[Char]]

  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part2("/day11/part1.txt")
    val result = sol.solve()
    println(s"Day 11 - Part 2 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()

  def walkDir(plan: Plan, coords: Coords, dir: (Int, Int)): Option[Char] = {
    if (
      coords.row >= plan.length || coords.row < 0 || coords.col >= plan(
        coords.row
      ).length || coords.col < 0
    ) return None
    plan(coords.row)(coords.col) match {
      case '.' =>
        walkDir(plan, Coords(coords.row + dir._1, coords.col + dir._2), dir)
      case x => Some(x)
    }
  }

  def getAdjacent(plan: Plan, coords: Coords): Seq[Char] = {
    val dirs =
      Seq((1, 0), (0, 1), (-1, 0), (0, -1), (1, 1), (-1, 1), (-1, -1), (1, -1))

    dirs
      .map(dir =>
        walkDir(plan, Coords(coords.row + dir._1, coords.col + dir._2), dir)
      )
      .flatten
  }

  def mapEmptySeat(adjacent: Seq[Char]): Char = {
    if (adjacent.forall(_ != '#')) '#' else 'L'
  }

  def mapOccupiedSeat(adjacent: Seq[Char]): Char = {
    if (adjacent.count(_ == '#') >= 5) 'L' else '#'
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
