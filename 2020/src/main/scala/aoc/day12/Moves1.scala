package aoc.day12

trait Step1 {
  def move(state: (Int, Int, String)): (Int, Int, String)
}
case class Move1(distance: Int) extends Step1 {
  override def move(state: (Int, Int, String)): (Int, Int, String) = {
    val vec = state._3 match {
      case "N" => (1, 0)
      case "S" => (-1, 0)
      case "E" => (0, 1)
      case "W" => (0, -1)
    }
    (state._1 + vec._1 * distance, state._2 + vec._2 * distance, state._3)
  }
}
case class MoveDir1(distance: Int, dir: String) extends Step1 {
  override def move(state: (Int, Int, String)): (Int, Int, String) = dir match {
    case "N" => (state._1 + distance, state._2, state._3)
    case "S" => (state._1 - distance, state._2, state._3)
    case "E" => (state._1, state._2 + distance, state._3)
    case "W" => (state._1, state._2 - distance, state._3)
  }
}
case class Turn1(to: String, by: Int) extends Step1 {
  val dirs = Seq("N", "E", "S", "W")
  def move(state: (Int, Int, String)): (Int, Int, String) = {
    val index = dirs.indexOf(state._3)
    val d = if (to == "R") 1 else -1
    val i = (index + by * d) % 4
    val positiveIndex = if (i < 0) i + 4 else i
    (state._1, state._2, dirs(positiveIndex))
  }
}

object Moves1 {
  def parseStep(line: String): Step1 = {
    val d = line.drop(1).toInt
    line.head match {
      case 'N' => MoveDir1(d, "N")
      case 'S' => MoveDir1(d, "S")
      case 'E' => MoveDir1(d, "E")
      case 'W' => MoveDir1(d, "W")
      case 'F' => Move1(d)
      case 'R' => Turn1("R", d / 90)
      case 'L' => Turn1("L", d / 90)
    }
  }
}
