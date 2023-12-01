mod days;
use days::*;
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        eprintln!("Usage: cargo run <day> [test case]");
        return;
    }

    let day = &args[1];
    let test_case = args.get(2).cloned();

    match day.as_str() {
        "day00" => days::day00::solution::Day00.run(test_case),
        "day01" => days::day01::solution::Day01.run(test_case),
        "day02" => days::day02::solution::Day02.run(test_case),
        "day03" => days::day03::solution::Day03.run(test_case),
        "day04" => days::day04::solution::Day04.run(test_case),
        "day05" => days::day05::solution::Day05.run(test_case),
        "day06" => days::day06::solution::Day06.run(test_case),
        "day07" => days::day07::solution::Day07.run(test_case),
        "day08" => days::day08::solution::Day08.run(test_case),
        "day09" => days::day09::solution::Day09.run(test_case),
        "day10" => days::day10::solution::Day10.run(test_case),
        "day11" => days::day11::solution::Day11.run(test_case),
        "day12" => days::day12::solution::Day12.run(test_case),
        "day13" => days::day13::solution::Day13.run(test_case),
        "day14" => days::day14::solution::Day14.run(test_case),
        "day15" => days::day15::solution::Day15.run(test_case),
        "day16" => days::day16::solution::Day16.run(test_case),
        "day17" => days::day17::solution::Day17.run(test_case),
        "day18" => days::day18::solution::Day18.run(test_case),
        "day19" => days::day19::solution::Day19.run(test_case),
        "day20" => days::day20::solution::Day20.run(test_case),
        "day21" => days::day21::solution::Day21.run(test_case),
        "day22" => days::day22::solution::Day22.run(test_case),
        "day23" => days::day23::solution::Day23.run(test_case),
        "day24" => days::day24::solution::Day24.run(test_case),
        "day25" => days::day25::solution::Day25.run(test_case),
        _ => eprintln!("Invalid day specified"),
    }
}
