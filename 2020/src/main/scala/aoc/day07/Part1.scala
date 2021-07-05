package aoc.day07

import aoc.Solution

import scala.language.implicitConversions

case class Rule(bag: Bag, count: Int) {
  override def equals(that: Any): Boolean = {
    that match {
      case that: Rule => this.bag == that.bag
      case that: String => this.bag == that
      case _ => false
    }
  }
}

case class Bag(name: String) {
  var rules: Seq[Rule] = Seq()

  override def equals(that: Any): Boolean = {
    that match {
      case that: Bag => this.name == that.name
      case that: String => this.name == that
      case _ => false
    }
  }

  def canContain(bagName: String): Boolean = {
    if (rules.isEmpty) false
    else if (rules.contains(bagName)) true
    else rules.exists(rule => rule.bag.canContain(bagName))
  }

  def hasToContain: Int = {
    if (rules.isEmpty) 0
    else this.rules.map(rule => {
      rule.count + rule.count * rule.bag.hasToContain
    }).sum
  }

  def addRule(rule: Rule): Unit = {
    this.rules = this.rules :+ rule
  }

  override def toString: String = this.name + " " + this.rules.length + " rules"
}

object Bag {
  // Take Bag and Rules and add them to the bag
  def addRulesToBags(bagAndRules: (Bag, String), bags: Seq[(Bag, String)]): Bag = {
    val pattern = "^(\\d+) ([\\w ?]+) (bag|bags)$".r
    val rules: Seq[Rule] = bagAndRules._2.split(", ").map(rule => {
      if(rule.trim == "no other bags") return bagAndRules._1
      val pattern(count, name, _) = rule.trim
      val ruleBag: Bag = bags.find(bag => bag._1.name == name).get._1
      Rule(ruleBag, count.toInt)
    })
    rules.foreach(rule => bagAndRules._1.addRule(rule))
    bagAndRules._1
  }

  // Create bag without rules and return rules as string
  private def createBagAndRules(rule: String): (Bag, String) = {
    val pattern = "^([\\w ]+) bags contain (.*).$".r
    val pattern(name, rules) = rule
    (Bag(name), rules)
  }

  // String of ALL rules -> Bags
  def apply(rules: Array[String]): Seq[Bag] = {
    val bagsAndRules = rules.map(rule => createBagAndRules(rule))
    bagsAndRules.map(bag => addRulesToBags(bag, bagsAndRules))
  }
}

case class Part1(inputPath : String) extends Solution(inputPath) {
  override def solve(): Int = {
    val bags = Bag(lines)
    bags.count(bag => bag.canContain("shiny gold"))
  }
}

object Part1 {
  def main(args: Array[String]): Unit = {
    val sol = Part1("/day07/part1.txt")
    val result = sol.solve()
    println(s"The result is: $result")
  }
}
