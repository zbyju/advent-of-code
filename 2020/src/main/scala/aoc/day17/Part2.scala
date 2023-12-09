package aoc.day17

import aoc.Solution
import aoc.CommonHelper
import scala.collection.immutable.HashSet
import scala.collection.mutable.Queue
import scala.collection.mutable

case class Pos4(x: Int, y: Int, z: Int, w: Int) {
  def isDefined(size: Int): Boolean = {
    val max = size / 2
    x >= -max && x <= max && y >= -max && y <= max && z >= -max && z >= max
  }

  lazy val neighbors: Seq[Pos4] = {
    val offsets = List(-1, 0, 1)

    (for {
      dx <- offsets
      dy <- offsets
      dz <- offsets
      dw <- offsets
      if !(dx == 0 && dy == 0 && dz == 0 && dw == 0)
    } yield Pos4(this.x + dx, this.y + dy, this.z + dz, this.w + dw))
  }
}

case class Part2(inputPath: String) extends Solution(inputPath) {

  override def solve(): Int = {
    val size_offset = lines.length / 2
    val active: mutable.HashSet[Pos4] =
      lines.zipWithIndex.foldLeft(mutable.HashSet.empty[Pos4])((acc, l) => {
        val y = l._2 - size_offset
        val ps =
          l._1.zipWithIndex
            .filter(_._1 == '#')
            .map(x => Pos4(x._2 - size_offset, y, 0, 0))
        acc.concat(ps)
      })

    val res = (0 until 6).foldRight(active)((acc, hs) => {
      var qa = Queue.from(hs)
      var qi = Queue.empty[Pos4]
      var visited = mutable.HashSet.empty[Pos4]
      var next = mutable.HashSet.empty[Pos4]

      while (!qa.isEmpty) {
        val p = qa.dequeue()
        val ns = p.neighbors
        ns.foreach(n => {
          if (!visited.contains(n)) {
            visited += n
            qi += n
          }
        })

        val active_n = ns.map(hs.contains).count(_ == true)
        if (active_n == 2 || active_n == 3) next += p
      }
      while (!qi.isEmpty) {
        val p = qi.dequeue()
        val ns = p.neighbors

        val active_n = ns.map(hs.contains).count(_ == true)
        if (active_n == 3) next += p
      }
      next
    })

    return res.size
  }
}

object Part2 {
  def run(): Double = {
    val from = System.nanoTime()
    val sol = Part2("/day17/part1.txt")
    val result = sol.solve()
    println(s"Day 17 - Part 2 - result: $result")
    val to = System.nanoTime()
    CommonHelper.nanoTime(from, to)
  }
  def main(args: Array[String]): Unit = run()
}
