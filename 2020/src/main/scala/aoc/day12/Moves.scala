package aoc.day12

trait Step {
  def move(state: (Int, Int, String)): (Int, Int, String)
}
case class Move(distance: Int) extends Step {
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
case class MoveDir(distance: Int, dir: String) extends Step {
  override def move(state: (Int, Int, String)): (Int, Int, String) = dir match {
    case "N" => (state._1 + distance, state._2, state._3)
    case "S" => (state._1 - distance, state._2, state._3)
    case "E" => (state._1, state._2 + distance, state._3)
    case "W" => (state._1, state._2 - distance, state._3)
  }
}
case class Turn(to: String, by: Int) extends Step {
  val dirs = Seq("N", "E", "S", "W")
  def move(state: (Int, Int, String)): (Int, Int, String) = {
    val index = dirs.indexOf(state._3)
    val d = if (to == "R") 1 else -1
    val i = (index + by * d) % 4
    val positiveIndex = if (i < 0) i + 4 else i
    (state._1, state._2, dirs(positiveIndex))
  }
}

object Moves {
  def parseStep(line: String): Step = {
    val d = line.drop(1).toInt
    line.head match {
      case 'N' => MoveDir(d, "N")
      case 'S' => MoveDir(d, "S")
      case 'E' => MoveDir(d, "E")
      case 'W' => MoveDir(d, "W")
      case 'F' => Move(d)
      case 'R' => Turn("R", d / 90)
      case 'L' => Turn("L", d / 90)
    }
  }
}
