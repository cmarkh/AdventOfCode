#![allow(dead_code)]

#[derive(Debug)]
struct Game {
    id: i32,
    hands: Vec<Hand>,
}

#[derive(Default, Clone, Copy, Debug)]
struct Hand {
    red: i32,
    green: i32,
    blue: i32,
}

/// Returns (game ID, cubes)
fn parse_input(input: &str) -> Vec<Game> {
    let mut games = Vec::new();
    let re_id = regex::Regex::new(r"Game (\d+)").unwrap();
    let re_hand = regex::Regex::new(r"(\d+) (red|blue|green)").unwrap();

    for line in input.lines() {
        let (game_id, game_str) = line.split_once(':').unwrap();

        let game_id = re_id
            .captures(game_id)
            .unwrap()
            .get(1)
            .unwrap()
            .as_str()
            .parse::<i32>()
            .unwrap();

        let hand_strs = game_str.split(';');
        let mut hands = Vec::new();
        for hand_str in hand_strs {
            let mut hand = Hand::default();
            for cap in re_hand.captures_iter(hand_str) {
                let count: i32 = cap[1].parse().unwrap();
                match &cap[2] {
                    "red" => hand.red = count,
                    "blue" => hand.blue = count,
                    "green" => hand.green = count,
                    _ => panic!(),
                }
            }
            hands.push(hand);
        }
        games.push(Game { id: game_id, hands });
    }

    games
}

const MAX_HAND: Hand = Hand {
    red: 12,
    green: 13,
    blue: 14,
};

fn part1(games: &Vec<Game>) -> i32 {
    let mut res = 0;

    let valid_game = |game: &Game| {
        for hand in &game.hands {
            if hand.red > MAX_HAND.red || hand.blue > MAX_HAND.blue || hand.green > MAX_HAND.green {
                return false;
            }
        }
        true
    };

    for game in games {
        if valid_game(game) {
            res += game.id;
        }
    }

    res
}

fn part2(games: &Vec<Game>) -> i32 {
    let mut res = 0;

    for game in games {
        let mut min_hand = Hand::default();
        for hand in &game.hands {
            if hand.red > min_hand.red {
                min_hand.red = hand.red;
            }
            if hand.green > min_hand.green {
                min_hand.green = hand.green;
            }
            if hand.blue > min_hand.blue {
                min_hand.blue = hand.blue;
            }
        }
        res += min_hand.red * min_hand.green * min_hand.blue;
    }

    res
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let input = std::fs::read_to_string("src/input.txt").unwrap();
        let games = parse_input(&input);
        let res = part1(&games);
        assert_eq!(res, 2505);
    }

    #[test]
    fn test_part2() {
        let input = std::fs::read_to_string("src/input.txt").unwrap();
        let games = parse_input(&input);
        let res = part2(&games);
        assert_eq!(res, 70265);
    }

    #[test]
    fn test_ex1() {
        let input = std::fs::read_to_string("src/ex1.txt").unwrap();
        let games = parse_input(&input);
        let res = part1(&games);
        assert_eq!(res, 8);
    }

    #[test]
    fn test_ex2() {
        let input = std::fs::read_to_string("src/ex1.txt").unwrap();
        let games = parse_input(&input);
        let res = part2(&games);
        assert_eq!(res, 2286);
    }

    #[test]
    #[ignore]
    fn test_parse_input() {
        let input = std::fs::read_to_string("src/ex1.txt").unwrap();
        let parsed = parse_input(&input);
        dbg!(parsed);
    }
}
