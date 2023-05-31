package aoc.day16

object Parser {
  type Range = (Int, Int)
  type Ranges = Map[String, (Range, Range)]
  type Ticket = Seq[Int]
  type Tickets = Seq[Ticket]

  def parse(input: String): (Ranges, Ticket, Tickets) = {
    val split = input.split("\n\n")
    val ranges: Ranges = parseRanges(split.head.split('\n'))
    val ticket: Ticket = parseTicket(split(1).split('\n'))
    val tickets: Tickets = parseTickets(split.last.split('\n'))

    (ranges, ticket, tickets)
  }

  def parseRanges(lines: Array[String]): Ranges = {
    lines
      .map(x => {
        val s1 = x.split(": ")
        val s2 = s1(1).split(" or ")
        (s1.head, s2.head, s2.last)
      })
      .foldLeft(Map.empty[String, (Range, Range)])((acc: Ranges, x) => {
        val r1 = x._2.split('-')
        val r2 = x._3.split('-')
        acc.updated(
          x._1,
          ((r1(0).toInt, r1(1).toInt), (r2(0).toInt, r2(1).toInt))
        )
      })
  }
  def parseTicket(lines: Array[String]): Ticket =
    lines.last.split(',').map(_.toInt)
  def parseTickets(lines: Array[String]): Tickets =
    lines.drop(1).map(_.split(',').map(_.toInt).toSeq)

}
