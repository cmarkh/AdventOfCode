#![allow(dead_code)]

use std::collections::HashSet;

#[derive(Debug)]
struct Card {
    winning_nums: HashSet<u32>,
    my_nums: HashSet<u32>,
}

fn parse_input(input: &str) -> Vec<Card> {
    let parse = |line: &str| {
        line.split(' ')
            .filter_map(|n| match n.trim() {
                "" => None,
                _ => Some(n.parse::<u32>().unwrap()),
            })
            .collect()
    };

    input
        .lines()
        .map(|line| {
            let line = line.split(':').nth(1).unwrap();
            let (winning_nums, my_nums) = line.split_once('|').unwrap();
            let winning_nums = parse(winning_nums);
            let my_nums = parse(my_nums);
            Card {
                winning_nums,
                my_nums,
            }
        })
        .collect()
}

fn part1(cards: &[Card]) -> u32 {
    cards
        .iter()
        .map(
            |card| match card.winning_nums.intersection(&card.my_nums).count() as u32 {
                0 => 0,
                c => 2_u32.pow(c - 1),
            },
        )
        .sum()
}

#[cfg(test)]
mod test {
    use super::*;
    use test_case::case;

    #[case("ex1.txt" => 13)]
    #[case("input.txt" => 20667)]
    fn test_part1(input_name: &str) -> u32 {
        let input = std::fs::read_to_string(format!("src/{}", input_name)).unwrap();
        let cards = parse_input(&input);
        part1(&cards)
    }
}
