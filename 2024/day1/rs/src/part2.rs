#![allow(dead_code)]

use std::collections::HashMap;

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

fn part2(left: Vec<isize>, right: Vec<isize>) -> isize {
    let mut similarity = 0;

    let mut right_map = HashMap::new();
    for n in right {
        *right_map.entry(n).or_insert(0) += 1;
    }

    for l in left {
        let r = right_map.entry(l).or_insert(0);
        similarity += l * *r;
    }

    similarity
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

    #[case("ex1.txt", 31)]
    #[case("input.txt", 21306195)]
    fn test_part2(file: &str, expected: isize) {
        let (left, right) = parse_input(file).unwrap();
        let result = part2(left, right);
        assert_eq!(result, expected);
    }
}
