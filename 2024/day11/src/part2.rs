#![allow(dead_code)]

use std::collections::HashMap;

use color_eyre::Result;

fn part2(stones: Vec<u64>, blinks: u64) -> u64 {
    let mut map = HashMap::new();
    for stone in stones {
        *map.entry(stone).or_insert(0) += 1;
    }

    for _ in 0..blinks {
        let mut new: HashMap<u64, u64> = HashMap::new();
        for (stone, count) in map.iter() {
            let str = stone.to_string();
            if *stone == 0 {
                *new.entry(1).or_insert(0) += count;
            } else if str.len() % 2 == 0 {
                let lhs: u64 = str[..str.len() / 2].parse().unwrap();
                let rhs: u64 = str[str.len() / 2..].parse().unwrap();
                *new.entry(lhs).or_insert(0) += count;
                *new.entry(rhs).or_insert(0) += count;
            } else {
                *new.entry(stone * 2024).or_insert(0) += count;
            }
        }
        map = new;
    }

    map.values().sum::<u64>()
}

fn read_file(file: &str) -> Result<Vec<u64>> {
    let input = std::fs::read_to_string(file)?;
    let stones: Vec<u64> = input
        .split(' ')
        .map(|s| s.parse())
        .collect::<Result<_, _>>()?;
    Ok(stones)
}

fn blink(stones: &[u64]) -> Vec<u64> {
    let mut new = Vec::new();

    for stone in stones.iter() {
        let str = stone.to_string();
        if *stone == 0 {
            new.push(1);
        } else if str.len() % 2 == 0 {
            let lhs: u64 = str[..str.len() / 2].parse().unwrap();
            let rhs: u64 = str[str.len() / 2..].parse().unwrap();
            new.push(lhs);
            new.push(rhs);
        } else {
            new.push(stone * 2024);
        }
    }

    new
}

#[cfg(test)]
mod tests {
    use test_case::case;

    use super::*;

    #[case("ex1.txt")]
    fn test_read_file(file: &str) {
        let stones = read_file(file).unwrap();
        println!("{:?}", stones);
    }

    #[case("ex2.txt", 6, 22)]
    #[case("ex2.txt", 25, 55312)]
    #[case("input.txt", 25, 202019)]
    #[case("input.txt", 75, 239321955280205)]
    fn test_part2(file: &str, blinks: u64, expected: u64) {
        let stones = read_file(file).unwrap();
        let count = part2(stones, blinks);
        assert_eq!(count, expected);
    }
}
