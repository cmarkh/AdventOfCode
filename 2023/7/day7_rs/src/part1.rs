use std::{cmp::Ordering, collections::HashMap};

#[derive(Debug, Clone)]
pub struct Hands(Vec<Hand>);

#[derive(Debug, Clone, Default, PartialEq)]
pub struct Hand {
    pub cards: Vec<char>,
    pub set: HashMap<char, u32>,
    pub bid: u32,
}

pub fn parse_input(input: &str) -> Result<Hands, Box<dyn std::error::Error>> {
    Ok(Hands(
        input
            .lines()
            .map(|line| {
                let mut split = line.split_ascii_whitespace();
                let cards: Vec<char> = split.next().ok_or("no split")?.chars().collect();
                let set = cards.iter().fold(HashMap::new(), |mut acc, c| {
                    *acc.entry(*c).or_insert(0) += 1;
                    acc
                });
                let bid = split.next().ok_or("no bid")?.parse::<u32>()?;
                Ok(Hand { cards, set, bid })
            })
            .collect::<Result<Vec<Hand>, Box<dyn std::error::Error>>>()?,
    ))
}

impl Hand {
    pub fn five_of_a_kind(&self) -> bool {
        self.set.len() == 1
    }

    pub fn four_of_a_kind(&self) -> bool {
        self.set.len() == 2
            && self.set.values().any(|count| *count == 4)
            && self.set.values().any(|count| *count == 1)
    }

    pub fn full_house(&self) -> bool {
        self.set.len() == 2
            && self.set.values().any(|count| *count == 3)
            && self.set.values().any(|count| *count == 2)
    }

    pub fn three_of_a_kind(&self) -> bool {
        self.set.values().any(|count| *count == 3)
    }

    pub fn two_pair(&self) -> bool {
        self.set.values().fold(0, |acc, count| match *count {
            2 => acc + 1,
            _ => acc,
        }) == 2
    }

    pub fn one_pair(&self) -> bool {
        self.set.values().any(|count| *count == 2)
    }

    pub fn high_card(&self) -> bool {
        true
    }

    pub fn rank(&self) -> u32 {
        match () {
            _ if self.five_of_a_kind() => 6,
            _ if self.four_of_a_kind() => 5,
            _ if self.full_house() => 4,
            _ if self.three_of_a_kind() => 3,
            _ if self.two_pair() => 2,
            _ if self.one_pair() => 1,
            _ if self.high_card() => 0,
            _ => panic!(),
        }
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
        match self.rank().cmp(&other.rank()) {
            Ordering::Equal => self
                .cards
                .iter()
                .zip(other.cards.iter())
                .map(|(m, o)| card_value(*m).cmp(&card_value(*o)))
                .find(|&order| order != Ordering::Equal)
                .unwrap_or(Ordering::Equal),
            other => other,
        }
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
        let input = std::fs::read_to_string(format!("src/{}", input)).unwrap();
        let hands = parse_input(&input).unwrap();
        dbg!(hands);
    }

    #[test]
    fn test_five_of_a_kind() {
        let hand = Hand {
            set: vec![('A', 5)].into_iter().collect(),
            ..Default::default()
        };
        assert!(hand.five_of_a_kind())
    }

    #[test]
    fn test_not_five_of_a_kind() {
        let hand = Hand {
            set: vec![('A', 4), ('B', 1)].into_iter().collect(),
            ..Default::default()
        };
        assert!(!hand.five_of_a_kind())
    }

    #[test]
    fn test_four_of_a_kind() {
        let hand = Hand {
            set: vec![('A', 4), ('B', 1)].into_iter().collect(),
            ..Default::default()
        };
        assert!(hand.four_of_a_kind())
    }

    #[test]
    fn test_full_house() {
        let hand = Hand {
            set: vec![('A', 3), ('B', 2)].into_iter().collect(),
            ..Default::default()
        };
        assert!(hand.full_house())
    }

    #[ignore]
    #[test]
    fn test_ranks() {
        let input = std::fs::read_to_string("src/ex.txt").unwrap();
        let hands = parse_input(&input).unwrap();
        for hand in hands.0 {
            dbg!(&hand.cards, hand.rank());
        }
    }

    #[ignore]
    #[test]
    fn test_sort() {
        let input = std::fs::read_to_string("src/ex.txt").unwrap();
        let mut hands = parse_input(&input).unwrap();
        hands.0.sort();
        hands.0.reverse();
        for hand in hands.0 {
            dbg!(&hand.cards, hand.rank());
        }
    }

    #[case("ex.txt"=> 6440)]
    #[case("input.txt" => 252295678)]
    fn test_part1(input: &str) -> u32 {
        let input = std::fs::read_to_string(format!("src/{}", input)).unwrap();
        let hands = parse_input(&input).unwrap();
        part1(hands)
    }
}
