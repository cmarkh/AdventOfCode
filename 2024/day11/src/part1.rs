#![allow(dead_code)]

use color_eyre::Result;

fn part1(mut stones: Vec<u64>, blinks: u64) -> usize {
    for _ in 0..blinks {
        stones = blink(&stones);
        // println!("{:?}", stones);
    }
    stones.len()
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
    #[case("input.txt", 75, 202019)]
    fn test_part1(file: &str, blinks: u64, expected: usize) {
        let stones = read_file(file).unwrap();
        let count = part1(stones, blinks);
        assert_eq!(count, expected);
    }
}
