package aoc.day04

import aoc.Solution

case class Part1(inputPath: String) extends Solution(inputPath) {

  val requiredKeys = Array("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid")
  val optionalKeys = Array("cid")

  private def getPassportsFromuInput(): Seq[String] = {
    lines.mkString("\n").split("\n\n")
  }

  private def createPresentKeyMap()
      : scala.collection.mutable.Map[String, Boolean] = {
    var presentKeyMap = scala.collection.mutable.Map[String, Boolean]()
    for (key <- requiredKeys) {
      presentKeyMap += (key -> false)
    }
    presentKeyMap
  }

  private def isPassportValid(passport: String): Boolean = {
    val presentKeys: scala.collection.mutable.Map[String, Boolean] =
      createPresentKeyMap()

    val fields = passport.split(" |\n")
    for (field <- fields) {
      val splittedField = field.split(":")
      val (key, value) = (splittedField(0), splittedField(1))

      presentKeys(key) = true
    }

    presentKeys.values.forall(_ == true)
  }

  override def solve(): Int = {
    val passports = getPassportsFromuInput()
    var numberOfValidPassports = 0
    for (passport <- passports) {
      if (isPassportValid(passport)) {
        numberOfValidPassports += 1
      }
    }
    numberOfValidPassports
  }
}

object Part1 {
  def main(args: Array[String]): Unit = {
    val sol = Part1("/day04/part2.txt")
    val result = sol.solve()
    println(s"Day 04 - Part 1 - result: $result")
  }
}
