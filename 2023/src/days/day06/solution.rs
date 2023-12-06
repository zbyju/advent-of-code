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
            let best = time / 2;
            let mut count = 0;

            for charge_time in (1..best).rev() {
                if (time - charge_time) * charge_time > distance_to_beat {
                    count += 1;
                } else {
                    break;
                }
            }
            for charge_time in best..time {
                if (time - charge_time) * charge_time > distance_to_beat {
                    count += 1;
                } else {
                    break;
                }
            }
            count
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
