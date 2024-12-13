#![allow(dead_code)]

use std::collections::HashSet;

use color_eyre::{
    eyre::{eyre, OptionExt},
    Result,
};

type Games = Vec<Game>;

#[derive(Debug)]
struct Game {
    a: (i64, i64),
    b: (i64, i64),
    prize: (i64, i64),
}

fn part1(games: &Games) -> i64 {
    let mut cost = 0;

    for game in games {
        cost += dbg!(solve(game));
    }

    cost
}

fn read_file(file: &str) -> Result<Games> {
    let mut games = Games::new();

    let input = std::fs::read_to_string(file)?;

    let mut i = 0;
    let mut lines = input.lines();
    while i < input.lines().count() {
        i += 4;

        let mut line = lines.next().ok_or_eyre("invalid input")?;
        line = match line.strip_prefix("Button A: ") {
            Some(line) => line,
            None => return Err(eyre!("invalid input")),
        };
        let (x, y) = line.split_once(",").ok_or_eyre("invalid input")?;
        let a_x: i64 = x
            .trim()
            .strip_prefix("X+")
            .ok_or_eyre("invalid input")?
            .parse()?;
        let a_y: i64 = y
            .trim()
            .strip_prefix("Y+")
            .ok_or_eyre("invalid input")?
            .parse()?;

        let mut line = lines.next().ok_or_eyre("invalid input")?;
        line = match line.strip_prefix("Button B: ") {
            Some(line) => line,
            None => return Err(eyre!("invalid input")),
        };
        let (x, y) = line.split_once(",").ok_or_eyre("invalid input")?;
        let b_x: i64 = x
            .trim()
            .strip_prefix("X+")
            .ok_or_eyre("invalid input")?
            .parse()?;
        let b_y: i64 = y
            .trim()
            .strip_prefix("Y+")
            .ok_or_eyre("invalid input")?
            .parse()?;

        let mut line = lines.next().ok_or_eyre("invalid input")?;
        line = match line.strip_prefix("Prize: ") {
            Some(line) => line,
            None => return Err(eyre!("invalid input")),
        };
        let (x, y) = line.split_once(",").ok_or_eyre("invalid input")?;
        let prize_x: i64 = x
            .trim()
            .strip_prefix("X=")
            .ok_or_eyre("invalid input")?
            .parse()?;
        let prize_y: i64 = y
            .trim()
            .strip_prefix("Y=")
            .ok_or_eyre("invalid input")?
            .parse()?;

        lines.next();

        let game = Game {
            a: (a_x, a_y),
            b: (b_x, b_y),
            prize: (prize_x, prize_y),
        };
        games.push(game);
    }

    Ok(games)
}

#[derive(Debug, Clone, Copy)]
struct Position {
    x: i64,
    y: i64,
    a_presses: i64,
    b_presses: i64,
}

impl Position {
    fn cost(&self) -> i64 {
        self.a_presses * 3 + self.b_presses
    }

    fn distance(&self, game: &Game) -> i64 {
        (game.prize.0 - self.x) + (game.prize.1 - self.y)
    }

    fn out_of_bounds(&self, game: &Game) -> bool {
        self.x > game.prize.0 || self.y > game.prize.1
    }

    fn reject(&self, game: &Game, min_cost: i64, max_presses: i64) -> bool {
        if self.out_of_bounds(game) {
            return true;
        }
        if self.a_presses > max_presses || self.b_presses > max_presses {
            return true;
        }

        let x = game.prize.0 - self.x;
        let y = game.prize.1 - self.y;

        let min_a = (x / game.a.0).max(y / game.a.1);
        let min_b = (x / game.b.0).max(y / game.b.1);

        if (min_a * 3).min(min_b) >= min_cost {
            return true;
        }

        false
    }
}

fn solve(game: &Game) -> i64 {
    let pos = Position {
        x: 0,
        y: 0,
        a_presses: 0,
        b_presses: 0,
    };

    let mut queue = vec![pos];
    let mut history = HashSet::new();
    let mut min_cost = i64::MAX;

    while let Some(pos) = queue.pop() {
        if !history.insert((pos.x, pos.y)) {
            continue;
        }
        if pos.reject(game, min_cost, 100) {
            continue;
        }
        if pos.x == game.prize.0 && pos.y == game.prize.1 {
            min_cost = min_cost.min(pos.cost());
            println!("{:?} {} {}", pos, pos.cost(), pos.distance(game));
            continue;
        }

        let mut pos2 = pos;
        pos2.a_presses += 1;
        pos2.x += game.a.0;
        pos2.y += game.a.1;
        queue.push(pos2);

        let mut pos2 = pos;
        pos2.b_presses += 1;
        pos2.x += game.b.0;
        pos2.y += game.b.1;
        queue.push(pos2);

        queue.sort_by_key(|a| a.cost());
    }

    match min_cost {
        i64::MAX => 0,
        _ => min_cost,
    }
}

#[cfg(test)]
mod tests {
    use test_case::case;

    use super::*;

    #[case("ex1.txt")]
    fn test_read_file(file: &str) {
        let input = read_file(file).unwrap();
        println!("{:?}", input);
    }

    #[case("ex1.txt", 480)]
    #[case("input.txt", 29517)]
    fn test_part1(file: &str, expected: i64) {
        let grid = read_file(file).unwrap();
        let price = part1(&grid);
        assert_eq!(price, expected);
    }
}
