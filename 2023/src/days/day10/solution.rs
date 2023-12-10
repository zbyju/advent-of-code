use std::collections::{HashMap, HashSet, VecDeque};

use crate::days::{AdventDay, SolutionOutput};

pub struct Day10;

fn parse(input: &str) -> (Vec<Vec<char>>, (usize, usize)) {
    let map: Vec<Vec<char>> = input.lines().map(|l| l.chars().collect()).collect();

    let len = map.get(0).unwrap().len() + 1;
    let i = input.find('S').unwrap();
    let y = i / len;
    let x = i % len;

    (map, (x, y))
}

fn determine_s(map: &Vec<Vec<char>>, pos_s: (usize, usize)) -> char {
    let up = *map
        .get(pos_s.1.wrapping_sub(1))
        .and_then(|row| row.get(pos_s.0))
        .unwrap_or(&'.');
    let down = *map
        .get(pos_s.1.wrapping_add(1))
        .and_then(|row| row.get(pos_s.0))
        .unwrap_or(&'.');
    let right = *map
        .get(pos_s.1)
        .and_then(|row| row.get(pos_s.0.wrapping_add(1)))
        .unwrap_or(&'.');
    let left = *map
        .get(pos_s.1)
        .and_then(|row| row.get(pos_s.0.wrapping_sub(1)))
        .unwrap_or(&'.');

    if "-LF".contains(left) && "-LF".contains(right) {
        return '-';
    } else if "|7F".contains(up) && "|7F".contains(down) {
        return '|';
    } else if "|LJ".contains(down) && "-J7".contains(right) {
        return 'F';
    } else if "|LJ".contains(down) && "-LF".contains(left) {
        return '7';
    } else if "|7F".contains(up) && "-J7".contains(right) {
        return 'L';
    } else if "|7F".contains(up) && "-LF".contains(left) {
        return 'J';
    }

    return '.';
}

fn next(map: &Vec<Vec<char>>, pos: (usize, usize)) -> Vec<(usize, usize)> {
    let mut char = map
        .get(pos.1)
        .and_then(|row| row.get(pos.0))
        .unwrap_or(&'.');

    if char == &'S' {
        let up = *map
            .get(pos.1.wrapping_sub(1))
            .and_then(|row| row.get(pos.0))
            .unwrap_or(&'.');
        let down = *map
            .get(pos.1.wrapping_add(1))
            .and_then(|row| row.get(pos.0))
            .unwrap_or(&'.');
        let right = *map
            .get(pos.1)
            .and_then(|row| row.get(pos.0.wrapping_add(1)))
            .unwrap_or(&'.');
        let left = *map
            .get(pos.1)
            .and_then(|row| row.get(pos.0.wrapping_sub(1)))
            .unwrap_or(&'.');

        if "-LF".contains(left) && "-LF".contains(right) {
            char = &'-';
        } else if "|7F".contains(up) && "|7F".contains(down) {
            char = &'|';
        } else if "|LJ".contains(down) && "-J7".contains(right) {
            char = &'F';
        } else if "|LJ".contains(down) && "-LF".contains(left) {
            char = &'7';
        } else if "|7F".contains(up) && "-J7".contains(right) {
            char = &'L';
        } else if "|7F".contains(up) && "-LF".contains(left) {
            char = &'J';
        }
    }

    match *char {
        '-' => vec![
            (pos.0.wrapping_sub(1), pos.1),
            (pos.0.wrapping_add(1), pos.1),
        ],
        '|' => vec![
            (pos.0, pos.1.wrapping_sub(1)),
            (pos.0, pos.1.wrapping_add(1)),
        ],
        'F' => vec![
            (pos.0.wrapping_add(1), pos.1),
            (pos.0, pos.1.wrapping_add(1)),
        ],
        '7' => vec![
            (pos.0.wrapping_sub(1), pos.1),
            (pos.0, pos.1.wrapping_add(1)),
        ],
        'L' => vec![
            (pos.0.wrapping_add(1), pos.1),
            (pos.0, pos.1.wrapping_sub(1)),
        ],
        'J' => vec![
            (pos.0.wrapping_sub(1), pos.1),
            (pos.0, pos.1.wrapping_sub(1)),
        ],
        _ => vec![],
    }
}

impl AdventDay for Day10 {
    fn input_base_path(&self) -> String {
        "src/days/day10".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        let (mut map, start) = parse(input);

        let s_char = determine_s(&map, start);
        map[start.1][start.0] = s_char;

        let mut q: VecDeque<(usize, usize)> = VecDeque::new();
        let mut distances: HashMap<(usize, usize), i64> = HashMap::new();
        let mut visited: HashSet<(usize, usize)> = HashSet::new();
        let mut max = 0;
        q.push_back(start);
        distances.insert(start, 0);
        visited.insert(start);

        while !q.is_empty() {
            let cur = q.pop_front().unwrap();
            let d = *distances.get(&cur).unwrap();
            let ns = next(&map, cur);

            for n in ns {
                if visited.contains(&n)
                    || map.get(n.1).and_then(|row| row.get(n.0)).unwrap_or(&'.') == &'.'
                {
                    continue;
                }

                let nd = d + 1;
                if nd > max {
                    max = nd;
                }

                q.push_back(n);
                distances.insert(n, nd);
                visited.insert(n);
            }
        }

        SolutionOutput::Int(max)
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        let (mut map, start) = parse(input);

        let s_char = determine_s(&map, start);
        map[start.1][start.0] = s_char;

        let mut q: VecDeque<(usize, usize)> = VecDeque::new();
        let mut distances: HashMap<(usize, usize), i64> = HashMap::new();
        let mut visited: HashSet<(usize, usize)> = HashSet::new();
        let mut max = 0;
        q.push_back(start);
        distances.insert(start, 0);
        visited.insert(start);

        while !q.is_empty() {
            let cur = q.pop_front().unwrap();
            let d = *distances.get(&cur).unwrap();
            let ns = next(&map, cur);

            for n in ns {
                if visited.contains(&n)
                    || map.get(n.1).and_then(|row| row.get(n.0)).unwrap_or(&'.') == &'.'
                {
                    continue;
                }

                let nd = d + 1;
                if nd > max {
                    max = nd;
                }

                q.push_back(n);
                distances.insert(n, nd);
                visited.insert(n);
            }
        }

        let lp: Vec<Vec<char>> = map
            .iter()
            .enumerate()
            .map(|(y, row)| {
                row.iter()
                    .enumerate()
                    .map(|(x, c)| if visited.contains(&(x, y)) { *c } else { '.' })
                    .collect()
            })
            .collect();

        let mut res = 0;

        for (y, row) in lp.iter().enumerate() {
            let mut num_intersections = 0;

            for (x, ch) in row.iter().enumerate() {
                match ch {
                    '|' | 'J' | 'L' if visited.contains(&(x, y)) => {
                        num_intersections += 1;
                    }
                    '.' if num_intersections % 2 == 1 => res += 1,
                    _ => {}
                }
            }
        }

        SolutionOutput::Int(res)
    }
}
