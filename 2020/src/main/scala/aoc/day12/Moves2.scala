package aoc.day12

case class State2(wp: (Int, Int), sh: (Int, Int, String)) {}

trait Step2 {
  def move(state: State2): State2
}
case class Move2(d: Int) extends Step2 {
  override def move(s: State2): State2 = {
    State2(
      s.wp,
      (
        s.sh._1 + s.wp._1 * d,
        s.sh._2 + s.wp._2 * d,
        s.sh._3
      )
    )
  }
}
case class MoveDir2(d: Int, dir: String) extends Step2 {
  override def move(s: State2): State2 = dir match {
    case "N" =>
      State2(
        (s.wp._1 + d, s.wp._2),
        s.sh
      )
    case "S" =>
      State2(
        (s.wp._1 - d, s.wp._2),
        s.sh
      )
    case "E" =>
      State2(
        (s.wp._1, s.wp._2 + d),
        s.sh
      )
    case "W" =>
      State2(
        (s.wp._1, s.wp._2 - d),
        s.sh
      )

  }
}
case class Turn2(by: Int) extends Step2 {
  override def move(s: State2): State2 = by match {
    case 0 | 360 => s
    case 90 =>
      State2(
        (-s.wp._2, s.wp._1),
        s.sh
      )
    case 180 =>
      State2(
        (-s.wp._1, -s.wp._2),
        s.sh
      )

    case 270 =>
      State2(
        (s.wp._2, -s.wp._1),
        s.sh
      )
  }
}

object Moves2 {
  def parseStep(line: String): Step2 = {
    val d = line.drop(1).toInt
    (line.head, d) match {
      case ('N', _)   => MoveDir2(d, "N")
      case ('S', _)   => MoveDir2(d, "S")
      case ('E', _)   => MoveDir2(d, "E")
      case ('W', _)   => MoveDir2(d, "W")
      case ('F', _)   => Move2(d)
      case ('R', _)   => Turn2(d)
      case ('L', 90)  => Turn2(270)
      case ('L', 270) => Turn2(90)
      case ('L', _)   => Turn2(d)
    }
  }
}
