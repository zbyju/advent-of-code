package aoc.day04

import aoc.Solution

import scala.util.matching.Regex

case class Part2(inputPath : String) extends Solution(inputPath) {

  val requiredKeys = Array("byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid")
  val optionalKeys = Array("cid")

  private def getPassportsFromuInput(): Seq[String] = {
    lines.mkString("\n").split("\n\n")
  }

  private def createPresentKeyMap(): scala.collection.mutable.Map[String, Boolean] = {
    var presentKeyMap = scala.collection.mutable.Map[String, Boolean]()
    for(key <- requiredKeys) {
      presentKeyMap += (key -> false)
    }
    presentKeyMap
  }

  private def isValueValid(key: String, value: String): Boolean = {
    key match {
      case "byr" => value.toInt >= 1920 && value.toInt <= 2002
      case "iyr" => value.toInt >= 2010 && value.toInt <= 2020
      case "eyr" => value.toInt >= 2020 && value.toInt <= 2030
      case "hgt" => {
        val reg: Regex = "[0-9]{2}[0-9]*(cm|in)".r
        if(!(reg matches value)) return false

        val unit = value.takeRight(2)
        val number = value.dropRight(2).toInt
        unit match {
          case "cm" => number >= 150 && number <= 193
          case "in" => number >= 59 && number <= 76
          case _ => false
        }
      }
      case "hcl" => {
        val reg: Regex = "#([0-9]|[a-f]){6}".r
        reg matches value
      }
      case "ecl" => {
        val reg: Regex = "amb|blu|brn|gry|grn|hzl|oth".r
        reg matches value
      }
      case "pid" => {
        val reg: Regex = "[0-9]{9}".r
        reg matches value
      }
      case "cid" => true
      case _ => false
    }
  }

  private def isPassportValid(passport: String): Boolean = {
    val validKeys: scala.collection.mutable.Map[String, Boolean] = createPresentKeyMap()
    val fields = passport.split(" |\n")

    for(field <- fields) {
      val splittedField = field.split(":")
      val (key, value) = (splittedField(0), splittedField(1))

      validKeys(key) = isValueValid(key, value)
    }

    validKeys.values.forall(_ == true)
  }

  override def solve(): Int = {
    val passports = getPassportsFromuInput()
    var numberOfValidPassports = 0
    for(passport <- passports) {
      if(isPassportValid(passport)) {
        numberOfValidPassports += 1
      }
    }
    numberOfValidPassports
  }
}

object Part2 {
  def main(args: Array[String]): Unit = {
    val sol = Part2("/day04/part2.txt")
    val result = sol.solve()
    println(s"The result is: $result")
  }
}
