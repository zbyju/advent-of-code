use crate::days::{AdventDay, SolutionOutput};

pub struct Day02;

impl AdventDay for Day02 {
    fn input_base_path(&self) -> String {
        "src/days/day02".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        let games = parse(input);

        let max_red = 12;
        let max_green = 13;
        let max_blue = 14;

        let res: i32 = games
            .iter()
            .filter(|g| {
                g.1.iter()
                    .filter(|s| {
                        s.iter()
                            .filter(|x| match x {
                                Color::Red(i) => *i > max_red,
                                Color::Green(i) => *i > max_green,
                                Color::Blue(i) => *i > max_blue,
                            })
                            .count()
                            > 0
                    })
                    .count()
                    == 0
            })
            .map(|g| g.0)
            .sum();

        SolutionOutput::Int(res as i64)
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        let games = parse(input);

        let res: i32 = games
            .iter()
            .map(|g| {
                g.1.iter().flatten().fold((0, 0, 0), |acc, c| match c {
                    Color::Red(x) => (acc.0.max(*x), acc.1, acc.2),
                    Color::Green(x) => (acc.0, acc.1.max(*x), acc.2),
                    Color::Blue(x) => (acc.0, acc.1, acc.2.max(*x)),
                })
            })
            .map(|x| x.0 * x.1 * x.2)
            .sum();
        SolutionOutput::Int(res as i64)
    }
}

enum Color {
    Blue(i32),
    Red(i32),
    Green(i32),
}

fn parse(input: &str) -> Vec<(i32, Vec<Vec<Color>>)> {
    input
        .lines()
        .enumerate()
        .map(|(index, line)| {
            let game_id = index as i32 + 1;
            let sets = line
                .split(':')
                .nth(1)
                .unwrap_or("")
                .split(';')
                .map(|round| {
                    round
                        .split(',')
                        .filter_map(|color_count| {
                            let mut parts = color_count.trim().split_whitespace();
                            let count = parts.next()?.parse().ok()?;
                            match parts.next()? {
                                "blue" => Some(Color::Blue(count)),
                                "red" => Some(Color::Red(count)),
                                "green" => Some(Color::Green(count)),
                                _ => None,
                            }
                        })
                        .collect()
                })
                .collect();
            (game_id, sets)
        })
        .collect()
}
