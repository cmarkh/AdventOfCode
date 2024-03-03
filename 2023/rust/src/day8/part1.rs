#![allow(dead_code)]

use std::collections::HashMap;

#[derive(Debug)]
enum Instruction {
    Left,
    Right,
}

#[derive(Debug)]
struct Instructions(Vec<Instruction>);

#[derive(Debug)]
struct Map(HashMap<String, (String, String)>);

fn parse_input(input: &str) -> (Instructions, Map) {
    let mut input = input.lines();

    let instructions = Instructions(
        input
            .next()
            .unwrap()
            .chars()
            .map(|c| match c {
                'L' => Instruction::Left,
                'R' => Instruction::Right,
                _ => panic!(),
            })
            .collect(),
    );

    input.next().ok_or("?").unwrap();
    let map = Map(input
        .map(|line| {
            let mut split = line.split(" = ");
            let key = split.next().ok_or("hold up").unwrap().to_string();
            let mut split = split.next().unwrap().split(", ");
            let left = split.next().unwrap().replace('(', "");
            let right = split.next().unwrap().replace(')', "");
            (key, (left, right))
        })
        .collect());

    (instructions, map)
}

fn part1(instructions: Instructions, map: Map) -> u32 {
    let mut position = &"AAA".to_string();

    for (steps, ins) in instructions.0.iter().cycle().enumerate() {
        if position == "ZZZ" {
            return steps as u32;
        }
        match ins {
            Instruction::Left => position = &map.0.get(position).unwrap().0,
            Instruction::Right => position = &map.0.get(position).unwrap().1,
        }
    }

    u32::MAX
}

#[cfg(test)]
mod test {
    use super::*;

    use test_case::case;

    #[ignore]
    #[case("ex1.txt")]
    #[case("ex2.txt")]
    #[case("input.txt")]
    fn test_input(input: &str) {
        let input = std::fs::read_to_string(format!("src/day8/{}", input)).unwrap();
        let (instructions, map) = parse_input(&input);
        dbg!(instructions, map);
    }

    #[case("ex1.txt" => 2)]
    #[case("ex2.txt" => 6)]
    #[case("input.txt" => 20777)]
    fn test_part1(input: &str) -> u32 {
        let input = std::fs::read_to_string(format!("src/day8/{}", input)).unwrap();
        let (instructions, map) = parse_input(&input);
        part1(instructions, map)
    }
}
