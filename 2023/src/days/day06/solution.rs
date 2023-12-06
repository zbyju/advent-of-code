use crate::days::{AdventDay, SolutionOutput};

pub struct Day06;

fn solve(input: String) -> SolutionOutput {
    let numbers: Vec<Vec<i64>> = input
        .lines()
        .map(|l| {
            l.split(':')
                .nth(1)
                .unwrap()
                .split(' ')
                .filter(|x| !x.is_empty())
                .map(|x| x.parse().unwrap())
                .collect()
        })
        .collect();

    let zipped: Vec<(i64, i64)> = numbers
        .get(0)
        .unwrap()
        .iter()
        .zip(numbers.get(1).unwrap().iter())
        .map(|(&a, &b)| (a, b))
        .collect();

    let number_of_ways_to_beat: Vec<i64> = zipped
        .iter()
        .map(|&(time, distance_to_beat)| {
            // Quadratic formula (a=1, b=time, c=distance_to_beat)
            let x = (time - f64::sqrt((time * time - 4 * distance_to_beat) as f64) as i64) / 2;
            let b = time - x;
            b - (b * (time - b) <= distance_to_beat) as i64
                - x
                - (x * (time - x) <= distance_to_beat) as i64
                + 1
        })
        .collect();

    SolutionOutput::Int(number_of_ways_to_beat.iter().product())
}

impl AdventDay for Day06 {
    fn input_base_path(&self) -> String {
        "src/days/day06".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        solve(input.to_string())
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        solve(input.replace(' ', ""))
    }
}
