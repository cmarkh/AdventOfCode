#![allow(dead_code)]

use std::collections::HashSet;

#[derive(Debug, Clone, Copy)]
struct Game {
    count: usize,
    new_hands: usize,
}

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

fn original_games(cards: &[Card]) -> Vec<Game> {
    cards
        .iter()
        .map(|card| Game {
            count: 1,
            new_hands: card.winning_nums.intersection(&card.my_nums).count(),
        })
        .collect()
}

fn part2(cards: &[Card]) -> usize {
    let mut games = original_games(cards);

    for g in 0..games.len() {
        for n in (g + 1)..(g + 1 + games[g].new_hands) {
            games[n].count += games[g].count;
        }
    }

    games.iter().fold(0, |acc, g| acc + g.count)
}

#[cfg(test)]
mod test {
    use super::*;
    use test_case::case;

    #[case("ex1.txt" => 30)]
    #[case("input.txt" => 5833065)]
    fn test_part1(input_name: &str) -> usize {
        let input = std::fs::read_to_string(format!("src/day4/{}", input_name)).unwrap();
        let cards = parse_input(&input);
        part2(&cards)
    }
}
