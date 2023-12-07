use crate::days::{AdventDay, SolutionOutput};

pub struct Day07;

fn char_value(c: char) -> u8 {
    match c {
        'A' => 14,
        'K' => 13,
        'Q' => 12,
        'J' => 11,
        'T' => 10,
        _ => c as u8 - '0' as u8,
    }
}

#[repr(u8)]
#[derive(Debug, Eq, PartialEq, Copy, Clone)]
enum HandType {
    FiveOfKind,
    FourOfKind,
    FullHouse,
    ThreeOfKind,
    TwoPair,
    OnePair,
    HighCard,
}

fn hand_type(card: &Vec<char>) -> HandType {
    let mut sorted = card.clone();
    sorted.sort_by(|a, b| char_value(*a).cmp(&char_value(*b)));
    let mut counts = sorted.into_iter().fold(Vec::new(), |mut acc, ch| {
        if let Some((last_char, count)) = acc.last_mut() {
            if *last_char == ch {
                *count += 1;
            } else {
                acc.push((ch, 1));
            }
        } else {
            acc.push((ch, 1));
        }
        acc
    });

    counts.sort_by_key(|x| x.1);

    let jacks = counts
        .iter()
        .filter(|&(ch, _)| *ch == '1')
        .nth(0)
        .map(|x| x.1)
        .unwrap_or(0);

    let mut cs = counts.iter().rev().filter(|x| x.0 != '1');
    let c0 = cs.next().unwrap_or(&('1', 0)).1;
    let c1 = cs.next().unwrap_or(&('1', 0)).1;

    if c0 + jacks == 5 {
        return HandType::FiveOfKind;
    }
    if c0 + jacks == 4 {
        return HandType::FourOfKind;
    }
    if c0 + jacks == 3 && c1 + (jacks - (3 - c0)) == 2 {
        return HandType::FullHouse;
    }
    if c0 + jacks == 3 {
        return HandType::ThreeOfKind;
    }
    if c0 + jacks == 2 && c1 + (jacks - (2 - c0)) == 2 {
        return HandType::TwoPair;
    }
    if c0 + jacks == 2 {
        return HandType::OnePair;
    }
    HandType::HighCard
}

impl AdventDay for Day07 {
    fn input_base_path(&self) -> String {
        "src/days/day07".to_string()
    }

    fn part1(&self, input: &str) -> SolutionOutput {
        let cards: Vec<(Vec<char>, i64, HandType)> = input
            .lines()
            .map(|l| {
                let split = l.split(' ');
                let card: Vec<char> = split.clone().nth(0).unwrap().chars().collect();
                let hand_type = hand_type(&card);
                let bid_str = split.clone().nth(1).unwrap();
                let bid: i64 = bid_str.parse().unwrap();

                (card, bid, hand_type)
            })
            .collect();

        let mut sorted_hands = cards;
        sorted_hands.sort_by(|a, b| {
            let a_type = a.2 as u8;
            let b_type = b.2 as u8;

            if a_type == b_type {
                for i in 0..5 {
                    let a_val = char_value(a.0[i as usize]);
                    let b_val = char_value(b.0[i as usize]);
                    if a_val != b_val {
                        return b_val.cmp(&a_val);
                    }
                }
            }

            a_type.cmp(&b_type)
        });

        let res = sorted_hands
            .iter()
            .rev()
            .enumerate()
            .fold(0, |acc, h| acc + (h.1 .1) * ((h.0 as i64) + 1));

        SolutionOutput::Int(res)
    }

    fn part2(&self, input: &str) -> SolutionOutput {
        let cards: Vec<(Vec<char>, i64, HandType)> = input
            .lines()
            .map(|l| {
                let split = l.split(' ');
                let card: Vec<char> = split
                    .clone()
                    .nth(0)
                    .unwrap()
                    .replace('J', "1")
                    .chars()
                    .collect();
                let hand_type = hand_type(&card);
                let bid_str = split.clone().nth(1).unwrap();
                let bid: i64 = bid_str.parse().unwrap();

                (card, bid, hand_type)
            })
            .collect();

        let mut sorted_hands = cards;
        sorted_hands.sort_by(|a, b| {
            let a_type = a.2 as u8;
            let b_type = b.2 as u8;

            if a_type == b_type {
                for i in 0..5 {
                    let a_val = char_value(a.0[i as usize]);
                    let b_val = char_value(b.0[i as usize]);
                    if a_val != b_val {
                        return b_val.cmp(&a_val);
                    }
                }
            }

            a_type.cmp(&b_type)
        });

        let res = sorted_hands
            .iter()
            .rev()
            .enumerate()
            .fold(0, |acc, h| acc + (h.1 .1) * ((h.0 as i64) + 1));

        SolutionOutput::Int(res)
    }
}
