mod days;
use days::*;
use std::env;

fn main() {
    let args: Vec<String> = env::args().collect();

    if args.len() < 2 {
        run_all();
        return;
    }

    let day = &args[1];
    let test_case = args.get(2).cloned();
    run_single(day, test_case)
}

fn run_all() {
    use std::time::Instant;
    let now = Instant::now();
    days::day00::solution::Day00.run(None, false);
    days::day01::solution::Day01.run(None, false);
    days::day02::solution::Day02.run(None, false);
    days::day03::solution::Day03.run(None, false);
    days::day04::solution::Day04.run(None, false);
    days::day05::solution::Day05.run(None, false);
    days::day06::solution::Day06.run(None, false);
    days::day07::solution::Day07.run(None, false);
    days::day08::solution::Day08.run(None, false);
    days::day09::solution::Day09.run(None, false);
    days::day10::solution::Day10.run(None, false);
    days::day11::solution::Day11.run(None, false);
    days::day12::solution::Day12.run(None, false);
    days::day13::solution::Day13.run(None, false);
    days::day14::solution::Day14.run(None, false);
    days::day15::solution::Day15.run(None, false);
    days::day16::solution::Day16.run(None, false);
    days::day17::solution::Day17.run(None, false);
    days::day18::solution::Day18.run(None, false);
    days::day19::solution::Day19.run(None, false);
    days::day20::solution::Day20.run(None, false);
    days::day21::solution::Day21.run(None, false);
    days::day22::solution::Day22.run(None, false);
    days::day23::solution::Day23.run(None, false);
    days::day24::solution::Day24.run(None, false);
    days::day25::solution::Day25.run(None, false);
    let elapsed = now.elapsed();
    println!("Time to solve all:   {:.2?}", elapsed);
}

fn run_single(day: &str, test_case: Option<String>) {
    match day {
        "day00" => days::day00::solution::Day00.run(test_case, true),
        "day01" => days::day01::solution::Day01.run(test_case, true),
        "day02" => days::day02::solution::Day02.run(test_case, true),
        "day03" => days::day03::solution::Day03.run(test_case, true),
        "day04" => days::day04::solution::Day04.run(test_case, true),
        "day05" => days::day05::solution::Day05.run(test_case, true),
        "day06" => days::day06::solution::Day06.run(test_case, true),
        "day07" => days::day07::solution::Day07.run(test_case, true),
        "day08" => days::day08::solution::Day08.run(test_case, true),
        "day09" => days::day09::solution::Day09.run(test_case, true),
        "day10" => days::day10::solution::Day10.run(test_case, true),
        "day11" => days::day11::solution::Day11.run(test_case, true),
        "day12" => days::day12::solution::Day12.run(test_case, true),
        "day13" => days::day13::solution::Day13.run(test_case, true),
        "day14" => days::day14::solution::Day14.run(test_case, true),
        "day15" => days::day15::solution::Day15.run(test_case, true),
        "day16" => days::day16::solution::Day16.run(test_case, true),
        "day17" => days::day17::solution::Day17.run(test_case, true),
        "day18" => days::day18::solution::Day18.run(test_case, true),
        "day19" => days::day19::solution::Day19.run(test_case, true),
        "day20" => days::day20::solution::Day20.run(test_case, true),
        "day21" => days::day21::solution::Day21.run(test_case, true),
        "day22" => days::day22::solution::Day22.run(test_case, true),
        "day23" => days::day23::solution::Day23.run(test_case, true),
        "day24" => days::day24::solution::Day24.run(test_case, true),
        "day25" => days::day25::solution::Day25.run(test_case, true),
        _ => eprintln!("Invalid day specified"),
    }
}
