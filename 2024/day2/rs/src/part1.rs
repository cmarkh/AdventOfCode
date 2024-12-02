#![allow(dead_code)]

use color_eyre::Result;

type Level = Vec<isize>;
type Levels = Vec<Level>;

fn parse_input(file: &str) -> Result<Levels> {
    let mut levels = Levels::new();

    let input = std::fs::read_to_string(file)?;
    for line in input.lines() {
        let digits = line
            .split(" ")
            .map(|s| s.parse::<isize>())
            .collect::<std::result::Result<Vec<isize>, _>>()?;
        levels.push(digits);
    }

    Ok(levels)
}

fn increasing(d1: isize, d2: isize) -> bool {
    d2 - d1 > 0
}

fn is_safe(level: Level) -> bool {
    let inc = level[1] - level[0] > 0;

    for i in 1..level.len() {
        let delta = level[i] - level[i - 1];
        if (delta > 0) != inc {
            return false;
        }
        if delta == 0 {
            return false;
        }
        if delta.abs() > 3 {
            return false;
        }
    }

    true
}

fn part1(levels: Levels) -> isize {
    let mut safe = 0;

    for level in levels {
        if is_safe(level) {
            safe += 1;
        }
    }

    safe
}

#[cfg(test)]
mod tests {
    use test_case::case;

    use super::*;

    #[case("ex1.txt")]
    fn test_parse_input(file: &str) {
        let levels = parse_input(file).unwrap();
        for row in levels {
            println!("{:?}", row);
        }
    }

    #[case("ex1.txt", 2)]
    #[case("input.txt", 483)]
    fn test_part1(file: &str, expected: isize) {
        let levels = parse_input(file).unwrap();
        let result = part1(levels);
        assert_eq!(result, expected);
    }
}
