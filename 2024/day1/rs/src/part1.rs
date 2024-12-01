#![allow(dead_code)]

use color_eyre::Result;

fn parse_input(file: &str) -> Result<(Vec<isize>, Vec<isize>)> {
    let mut left: Vec<isize> = Vec::new();
    let mut right: Vec<isize> = Vec::new();

    let input = std::fs::read_to_string(file)?;
    for line in input.lines() {
        let split: Vec<&str> = line.split(" ").filter(|ch| !ch.is_empty()).collect();
        left.push(split[0].parse()?);
        right.push(split[1].parse()?);
    }

    Ok((left, right))
}

fn part1(mut left: Vec<isize>, mut right: Vec<isize>) -> isize {
    let mut distance = 0;

    left.sort();
    right.sort();

    for (l, r) in left.iter().zip(right.iter()) {
        distance += (l - r).abs();
    }

    distance
}

#[cfg(test)]
mod test {
    use test_case::case;

    use super::*;

    #[case("ex1.txt")]
    fn test_parse_input(file: &str) {
        let (left, right) = parse_input(file).unwrap();
        println!("{:?}\n{:?}", left, right)
    }

    #[case("ex1.txt", 11)]
    #[case("input.txt", 1651298)]
    fn test_part1(file: &str, expected: isize) {
        let (left, right) = parse_input(file).unwrap();
        let result = part1(left, right);
        assert_eq!(result, expected);
    }
}
