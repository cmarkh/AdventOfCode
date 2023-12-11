use std::{cmp::Ordering, collections::HashMap};

#[derive(Debug, Clone)]
pub struct Hands(Vec<Hand>);

#[derive(Debug, Clone, Default, PartialEq)]
pub struct Hand {
    pub cards: Vec<char>,
    pub bid: u32,
}

pub fn parse_input(input: &str) -> Result<Hands, Box<dyn std::error::Error>> {
    Ok(Hands(
        input
            .lines()
            .map(|line| {
                let mut split = line.split_ascii_whitespace();
                let cards: Vec<char> = split.next().ok_or("no split")?.chars().collect();
                let bid = split.next().ok_or("no bid")?.parse::<u32>()?;
                Ok(Hand { cards, bid })
            })
            .collect::<Result<Vec<Hand>, Box<dyn std::error::Error>>>()?,
    ))
}

#[derive(Debug, Clone, Copy, PartialEq, Eq, PartialOrd, Ord)]
pub enum Rank {
    HighCard,
    OnePair,
    TwoPair,
    ThreeOfAKind,
    FullHouse,
    FourOfAKind,
    FiveOfAKind,
}

impl Hand {
    pub fn set(&self) -> HashMap<char, u32> {
        self.cards.iter().fold(HashMap::new(), |mut acc, c| {
            *acc.entry(*c).or_insert(0) += 1;
            acc
        })
    }
}

impl Eq for Hand {}

impl PartialOrd for Hand {
    fn partial_cmp(&self, other: &Self) -> Option<Ordering> {
        Some(self.cmp(other))
    }
}

impl Ord for Hand {
    fn cmp(&self, other: &Self) -> Ordering {
        let hand_cmp = || -> Ordering {
            self.cards
                .iter()
                .zip(other.cards.iter())
                .map(|(m, o)| card_value(*m).cmp(&card_value(*o)))
                .find(|&order| order != Ordering::Equal)
                .unwrap_or(Ordering::Equal)
        };

        match rank(&self.set()).cmp(&rank(&other.set())) {
            Ordering::Equal => hand_cmp(),
            other => other,
        }
    }
}

pub fn rank(set: &HashMap<char, u32>) -> Rank {
    match set {
        set if set.len() == 1 => Rank::FiveOfAKind,
        set if set.len() == 2
            && set.values().any(|count| *count == 4)
            && set.values().any(|count| *count == 1) =>
        {
            Rank::FourOfAKind
        }
        set if set.len() == 2
            && set.values().any(|count| *count == 3)
            && set.values().any(|count| *count == 2) =>
        {
            Rank::FullHouse
        }
        set if set.values().any(|count| *count == 3) => Rank::ThreeOfAKind,
        set if set.values().fold(0, |acc, count| match *count {
            2 => acc + 1,
            _ => acc,
        }) == 2 =>
        {
            Rank::TwoPair
        }
        set if set.values().any(|count| *count == 2) => Rank::OnePair,
        _ => Rank::HighCard,
    }
}

pub fn card_value(c: char) -> u32 {
    match c {
        '2'..='9' => c.to_digit(10).unwrap(),
        'T' => 9 + 1,
        'J' => 9 + 2,
        'Q' => 9 + 3,
        'K' => 9 + 4,
        'A' => 9 + 5,
        _ => panic!(),
    }
}

pub fn part1(mut hands: Hands) -> u32 {
    let mut winnings = 0;

    hands.0.sort();
    for (i, hand) in hands.0.iter().enumerate() {
        winnings += (i as u32 + 1) * hand.bid;
    }

    winnings
}

#[cfg(test)]
mod test {
    use super::*;

    use test_case::case;

    #[ignore]
    #[case("ex.txt")]
    #[case("input.txt")]
    fn test_parse_input(input: &str) {
        let input = std::fs::read_to_string(format!("src/day7/{}", input)).unwrap();
        let hands = parse_input(&input).unwrap();
        dbg!(hands);
    }

    #[ignore]
    #[case("ex.txt")]
    #[case("input.txt")]
    fn test_rank_alt_j(input: &str) {
        let input = std::fs::read_to_string(format!("src/day7/{}", input)).unwrap();
        let mut hands = parse_input(&input).unwrap();
        hands.0.sort();
        for hand in hands.0 {
            dbg!(&hand, rank(&hand.set()));
        }
    }

    #[case("ex.txt"=> 6440)]
    #[case("input.txt" => 252295678)]
    fn test_part1(input: &str) -> u32 {
        let input = std::fs::read_to_string(format!("src/day7/{}", input)).unwrap();
        let hands = parse_input(&input).unwrap();
        part1(hands)
    }
}
