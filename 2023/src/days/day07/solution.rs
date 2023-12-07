use crate::days::{AdventDay, SolutionOutput};

pub struct Day07;

fn char_value(c: char) -> u8 {
    match c {
        'A' => 14,
        'K' => 13,
        'Q' => 12,
        'J' => 11,
        'T' => 10,
        _ => c as u8 - b'0',
    }
}
fn char_value2(c: char) -> u8 {
    match c {
        'A' => 14,
        'K' => 13,
        'Q' => 12,
        'J' => 1,
        'T' => 10,
        _ => c as u8 - b'0',
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

fn count_cards(hand: &Vec<u8>) -> (u8, u8, u8) {
    let mut counts: [u8; 13] = [0; 13];
    let mut max1 = 0;
    let mut max2 = 0;
    let mut jacks = 0;
    for x in hand {
        if *x == 1 {
            jacks += 1;
            continue;
        }
        counts[(x - 2) as usize] += 1;
    }
    for value in counts {
        if value > max1 {
            max2 = max1;
            max1 = value;
        } else if value > max2 {
            max2 = value;
        }
    }

    (max1, max2, jacks)
}

fn hand_type(hand: &Vec<u8>) -> HandType {
    let (c0, c1, jacks) = count_cards(hand);

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
        let hands: Vec<(Vec<u8>, i64, HandType)> = input
            .lines()
            .map(|l| {
                let split = l.split(' ');
                let hand: Vec<u8> = split
                    .clone()
                    .next()
                    .unwrap()
                    .chars()
                    .map(char_value)
                    .collect();
                let hand_type = hand_type(&hand);
                let bid_str = split.clone().nth(1).unwrap();
                let bid: i64 = bid_str.parse().unwrap();

                (hand, bid, hand_type)
            })
            .collect();

        let mut sorted_hands = hands;
        sorted_hands.sort_by(|a, b| {
            let a_type = a.2 as u8;
            let b_type = b.2 as u8;

            if a_type == b_type {
                for i in 0..5 {
                    let a_val = a.0[i as usize];
                    let b_val = b.0[i as usize];
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
        let hands: Vec<(Vec<u8>, i64, HandType)> = input
            .lines()
            .map(|l| {
                let split = l.split(' ');
                let hand: Vec<u8> = split
                    .clone()
                    .next()
                    .unwrap()
                    .chars()
                    .map(char_value2)
                    .collect();
                let hand_type = hand_type(&hand);
                let bid_str = split.clone().nth(1).unwrap();
                let bid: i64 = bid_str.parse().unwrap();

                (hand, bid, hand_type)
            })
            .collect();

        let mut sorted_hands = hands;
        sorted_hands.sort_by(|a, b| {
            let a_type = a.2 as u8;
            let b_type = b.2 as u8;

            if a_type == b_type {
                for i in 0..5 {
                    let a_val = a.0[i as usize];
                    let b_val = b.0[i as usize];
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
