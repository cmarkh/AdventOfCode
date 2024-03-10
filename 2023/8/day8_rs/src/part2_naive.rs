#![allow(dead_code)]

use std::{collections::HashMap, rc::Rc};

#[derive(Debug, PartialEq, PartialOrd, Eq, Ord)]
enum Instruction {
    Left,
    Right,
}

#[derive(Debug)]
struct Instructions(Vec<Instruction>);

/// Map<position, (left, right)>
#[derive(Debug, Clone)]
struct Map(HashMap<String, (String, String)>);

/// Map<(Start Position, Instruction Index), (steps_to_end, end)>
#[derive(Debug)]
struct StepsToZ<'a>(HashMap<(&'a str, usize), (u32, &'a str)>);

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

impl<'a> Map {
    fn starts(&'a self) -> Vec<&'a String> {
        self.0.keys().filter(|key| key.ends_with('A')).collect()
    }

    fn ends(&'a self) -> Vec<&'a String> {
        self.0.keys().filter(|key| key.ends_with('Z')).collect()
    }
}

fn part_2(instructions: Instructions, map: Map) -> u32 {
    let instructions = Rc::new(instructions);
    let map = Rc::new(map.clone());

    let mut positions = map.starts();

    for (steps, ins) in instructions.0.iter().cycle().enumerate() {
        dbg!(steps);
        if positions.iter().all(|pos| pos.ends_with('Z')) {
            return steps as u32;
        }
        match ins {
            Instruction::Left => {
                positions = positions
                    .iter()
                    .map(|position| &map.0.get(*position).unwrap().0)
                    .collect()
            }
            Instruction::Right => {
                positions = positions
                    .iter()
                    .map(|position| &map.0.get(*position).unwrap().1)
                    .collect()
            }
        }
    }

    unreachable!();
}

fn least_common_multiple(steps: Vec<u32>) -> u32 {
    let mut min = *steps.iter().min().unwrap();
    while min > 0 {
        for &step in &steps {
            if step % min != 0 {
                min -= 1;
                continue;
            }
        }
    }

    0
}

#[cfg(test)]
mod test {
    use super::*;

    use test_case::case;

    fn get_input(file_name: &str) -> (Instructions, Map) {
        let input = std::fs::read_to_string(format!("src/{}", file_name)).unwrap();
        parse_input(&input)
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("ex2.txt")]
    #[case("input.txt")]
    fn test_input(file_name: &str) {
        let (instructions, map) = get_input(file_name);
        dbg!(instructions, map);
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("ex2.txt")]
    #[case("input.txt")]
    fn test_starts(file_name: &str) {
        let (_, map) = get_input(file_name);
        let starts = map.starts();
        dbg!(starts);
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("ex2.txt")]
    #[case("input.txt")]
    fn test_ends(file_name: &str) {
        let (_, map) = get_input(file_name);
        let ends = map.ends();
        dbg!(ends);
    }

    #[case("ex3.txt" => 6)]
    #[case("input.txt" => 0)]
    fn test_part_2_naive(file_name: &str) -> u32 {
        let (instructions, map) = get_input(file_name);
        let steps = part_2(instructions, map);
        dbg!(steps);
        steps
    }
}
