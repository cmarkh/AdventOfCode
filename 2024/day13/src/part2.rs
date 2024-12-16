#![allow(dead_code)]

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
        let x = (game.a.0, game.b.0, game.prize.0);
        let y = (game.a.1, game.b.1, game.prize.1);

        let (a, b) = match solve2(x, y) {
            Some((a, b)) => (a, b),
            None => continue,
        };
        if a < 0 || b < 0 {
            continue;
        }

        println!("\n{}a + {}b = {}", game.a.0, game.b.0, game.prize.0);
        println!("{}a + {}b = {}", game.a.1, game.b.1, game.prize.1);
        println!("a: {}, b: {}", a, b);

        cost += a * 3 + b;
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
            prize: (prize_x + 10000000000000, prize_y + 10000000000000),
        };
        games.push(game);
    }

    Ok(games)
}

fn solve2((a1, b1, s1): (i64, i64, i64), (a2, b2, s2): (i64, i64, i64)) -> Option<(i64, i64)> {
    // a1*x + b1*y = s1
    // a2*x + b2*y = s2

    // let one = (a1 * b2, b1 * b2, s1 * b2);
    // let two = (a2 * b1, b2 * b1, s2 * b1);
    // let diff = (a1 * b2 - a2 * b1, s1 * b2 - s2 * b1);

    let x = (s1 * b2 - s2 * b1) / (a1 * b2 - a2 * b1);

    // let one = (a1 * a2, b1 * a2, s1 * a2);
    // let two = (a2 * a1, b2 * a1, s2 * a1);
    // let diff = (b1 * a2 - b2 * a1, s1 * a2 - s2 * a1);

    let y = (s1 * a2 - s2 * a1) / (b1 * a2 - b2 * a1);

    // need to check if rounding occured
    if x * a1 + y * b1 == s1 && x * a2 + y * b2 == s2 {
        Some((x, y))
    } else {
        None
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

    #[case((94, 22, 8400), (34,67,5400), Some((80, 40)))]
    #[case((68, 14, 5034), (17, 76, 2051), Some((71, 10)))]
    fn test_solve2(a: (i64, i64, i64), b: (i64, i64, i64), expected: Option<(i64, i64)>) {
        let result = solve2(a, b);
        dbg!(result);
        assert_eq!(result, expected);
    }

    #[case("input.txt", 103570327981381)]
    fn test_part1(file: &str, expected: i64) {
        let grid = read_file(file).unwrap();
        let price = part1(&grid);
        assert_eq!(price, expected);
    }
}
