use crate::days::{AdventDay, SolutionOutput};

pub struct Day09;

fn parse(input: &str) -> Vec<Vec<i64>> {
    input
        .lines()
        .map(|l| l.split(' ').map(|x| x.parse().unwrap()).collect())
        .collect()
}

fn extrapolate(history: &Vec<i64>, acc: i64) -> i64 {
    if history
        .iter()
        .filter(|x| *x != history.get(0).unwrap())
        .count()
        == 0
    {
        return acc + history.last().unwrap();
    }
    let next: Vec<i64> = history
        .windows(2)
        .map(|w| w.get(1).unwrap() - w.get(0).unwrap())
        .collect();
    extrapolate(&next, acc + history.last().unwrap())
}
fn extrapolate2(history: &Vec<i64>) -> i64 {
    if history
        .iter()
        .filter(|x| *x != history.get(0).unwrap())
        .count()
        == 0
    {
        return *history.get(0).unwrap();
    }
    let next: Vec<i64> = history
        .windows(2)
        .map(|w| w.get(1).unwrap() - w.get(0).unwrap())
        .collect();
    history.get(0).unwrap() - extrapolate2(&next)
}

impl AdventDay for Day09 {
    fn input_base_path(&self) -> String {
        "src/days/day09".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        let histories = parse(input);

        let res = histories.iter().map(|h| extrapolate(h, 0)).sum();

        SolutionOutput::Int(res)
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        let histories = parse(input);

        let res = histories.iter().map(|h| extrapolate2(h)).sum();
        SolutionOutput::Int(res)
    }
}
