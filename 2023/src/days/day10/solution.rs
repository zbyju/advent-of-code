use std::collections::HashSet;

use crate::days::{AdventDay, SolutionOutput};

pub struct Day10;

#[derive(PartialEq, Eq, Hash, Copy, Clone, Debug)]
struct Pos(usize, usize);

impl Pos {
    fn new(x: usize, y: usize) -> Self {
        Pos(x, y)
    }

    fn val(&self) -> (i64, i64) {
        (self.0 as i64, self.1 as i64)
    }

    fn x(&self) -> usize {
        self.0
    }

    fn y(&self) -> usize {
        self.1
    }

    fn next(&self, dir: &Dir) -> Option<Pos> {
        let x = (self.x() as isize).checked_add(dir.x());
        let y = (self.y() as isize).checked_add(dir.y());
        match (x, y) {
            (Some(x), Some(y)) => Some(Pos(x as usize, y as usize)),
            _ => None,
        }
    }

    fn nexts(&self, grid: &[Vec<char>]) -> Vec<Pos> {
        let ch = self.char_at(grid);
        match ch {
            '|' => vec![self.next(&Dir::Up).unwrap(), self.next(&Dir::Down).unwrap()],
            '-' => vec![
                self.next(&Dir::Left).unwrap(),
                self.next(&Dir::Right).unwrap(),
            ],
            'L' => vec![
                self.next(&Dir::Up).unwrap(),
                self.next(&Dir::Right).unwrap(),
            ],
            'J' => vec![self.next(&Dir::Up).unwrap(), self.next(&Dir::Left).unwrap()],
            'F' => vec![
                self.next(&Dir::Down).unwrap(),
                self.next(&Dir::Right).unwrap(),
            ],
            '7' => vec![
                self.next(&Dir::Down).unwrap(),
                self.next(&Dir::Left).unwrap(),
            ],
            _ => vec![],
        }
    }

    fn at<'a>(&self, grid: &'a [Vec<char>]) -> Option<&'a char> {
        grid.get(self.y()).and_then(|row| row.get(self.x()))
    }

    fn char_at<'a>(&self, grid: &'a [Vec<char>]) -> &'a char {
        self.at(grid).unwrap_or(&'.')
    }
}

enum Dir {
    Up,
    Down,
    Right,
    Left,
}

impl Dir {
    fn x(&self) -> isize {
        match self {
            Dir::Up => 0,
            Dir::Down => 0,
            Dir::Left => -1,
            Dir::Right => 1,
        }
    }

    fn y(&self) -> isize {
        match self {
            Dir::Up => -1,
            Dir::Down => 1,
            Dir::Left => 0,
            Dir::Right => 0,
        }
    }
}

fn parse(input: &str) -> (Vec<Vec<char>>, Pos) {
    let map: Vec<Vec<char>> = input.lines().map(|l| l.chars().collect()).collect();

    let len = map.get(0).unwrap().len() + 1;
    let i = input.find('S').unwrap();
    let y = i / len;
    let x = i % len;

    (map, Pos::new(x, y))
}

fn determine_s(map: &[Vec<char>], pos_s: &Pos) -> char {
    let mut possibilities = HashSet::from(['|', '-', 'L', 'J', 'F', '7']);

    let up = pos_s.next(&Dir::Up).and_then(|p| p.at(map)).unwrap_or(&'.');
    if "-LJ.".contains(*up) {
        possibilities.remove(&'|');
        possibilities.remove(&'L');
        possibilities.remove(&'J');
    }
    let down = pos_s
        .next(&Dir::Down)
        .and_then(|p| p.at(map))
        .unwrap_or(&'.');
    if "-F7.".contains(*down) {
        possibilities.remove(&'|');
        possibilities.remove(&'7');
        possibilities.remove(&'F');
    }
    let right = pos_s
        .next(&Dir::Right)
        .and_then(|p| p.at(map))
        .unwrap_or(&'.');
    if "|LF.".contains(*right) {
        possibilities.remove(&'-');
        possibilities.remove(&'L');
        possibilities.remove(&'F');
    }
    let left = pos_s
        .next(&Dir::Left)
        .and_then(|p| p.at(map))
        .unwrap_or(&'.');
    if "|7J.".contains(*left) {
        possibilities.remove(&'-');
        possibilities.remove(&'7');
        possibilities.remove(&'J');
    }
    *possibilities.iter().next().unwrap()
}

fn shoelace(points: &[(i64, i64)]) -> i64 {
    let n = points.len();
    let mut area = 0;

    for i in 0..n {
        let (x1, y1) = points[i];
        let (x2, y2) = if i + 1 < n { points[i + 1] } else { points[0] };

        area += (x1 * y2) - (y1 * x2);
    }

    area.abs() / 2
}

impl AdventDay for Day10 {
    fn input_base_path(&self) -> String {
        "src/days/day10".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        let (mut map, start) = parse(input);

        let s_char = determine_s(&map, &start);
        map[start.y()][start.x()] = s_char;

        let mut cur = start.nexts(&map)[0];
        let mut previous: Pos = start;
        let mut visited: HashSet<Pos> = HashSet::from([start]);

        while cur != start {
            visited.insert(cur);
            let ns = cur.nexts(&map);
            let next = ns.iter().find(|&&pos| pos != previous).unwrap();

            previous = cur;
            cur = *next;
        }

        SolutionOutput::Int(visited.len() as i64 / 2)
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        let (mut map, start) = parse(input);

        let s_char = determine_s(&map, &start);
        map[start.y()][start.x()] = s_char;

        let mut cur = start.nexts(&map)[0];
        let mut previous: Pos = start;
        let mut seen: Vec<(i64, i64)> = vec![start.val()];
        let mut visited: HashSet<Pos> = HashSet::from([start]);

        while cur != start {
            visited.insert(cur);
            seen.push(cur.val());
            let ns = cur.nexts(&map);
            let next = ns.iter().find(|&&pos| pos != previous).unwrap();

            previous = cur;
            cur = *next;
        }

        let a = shoelace(&seen);
        let res = a + 1 - seen.len() as i64 / 2;
        SolutionOutput::Int(res)
    }
}
