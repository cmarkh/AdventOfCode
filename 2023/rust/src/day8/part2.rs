#![allow(dead_code)]

use std::collections::HashMap;

#[derive(Debug, PartialEq, PartialOrd, Eq, Ord)]
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

fn steps_to_all_z(instructions: Instructions, map: Map) {
    let zs: Vec<&String> = map.0.keys().filter(|key| key.ends_with('Z')).collect();
    dbg!(zs);
}

// #[derive(Debug, Clone, PartialEq, Eq)]
// struct StepsToZ {
//     start_pos: String,
//     start_instruction: usize,
//     steps_to_z: usize,
// }

#[derive(Debug, Clone, PartialEq, Eq, Hash)]
struct Start {
    position: String,
    instruction: usize,
}

fn pos_to_z(start: &Start, instructions: &Instructions, map: &Map) -> Option<usize> {
    let mut position = &start.position;

    for (steps, ins) in instructions
        .0
        .iter()
        .cycle()
        .skip(start.instruction)
        .enumerate()
    {
        if position.ends_with('Z') {
            return Some(steps);
        }
        if steps > instructions.0.len() * 10 {
            return None;
        }
        match ins {
            Instruction::Left => position = &map.0.get(position).unwrap().0,
            Instruction::Right => position = &map.0.get(position).unwrap().1,
        }
    }

    unreachable!()
}

fn positions_to_z(
    starting_instruction: usize,
    instructions: &Instructions,
    map: &Map,
) -> HashMap<Start, usize> {
    map.0
        .keys()
        .filter_map(|key| {
            let start = Start {
                position: key.clone(),
                instruction: starting_instruction,
            };
            pos_to_z(&start, instructions, map).map(|steps| (start, steps))
        })
        .collect()
}

fn instructions_to_z(instructions: &Instructions, map: &Map) -> HashMap<Start, usize> {
    let mut out = HashMap::new();
    for (i, _) in instructions.0.iter().enumerate() {
        out.extend(positions_to_z(i, instructions, map))
    }
    out
}

fn part2(instructions: Instructions, map: Map) -> usize {
    let mut positions: Vec<&String> = map.0.keys().filter(|key| key.ends_with('A')).collect();

    for (steps, ins) in instructions.0.iter().cycle().enumerate() {
        if positions.iter().all(|pos| pos.ends_with('Z')) {
            return steps;
        }
        match ins {
            Instruction::Left => {
                for pos in &mut positions {
                    *pos = &map.0.get(*pos).unwrap().0;
                }
            }
            Instruction::Right => {
                for pos in &mut positions {
                    *pos = &map.0.get(*pos).unwrap().1;
                }
            }
        }
    }

    unreachable!()
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

    #[case("ex3.txt" => 6)]
    #[case("input.txt" => 0)]
    fn test_part2(input: &str) -> usize {
        let input = std::fs::read_to_string(format!("src/day8/{}", input)).unwrap();
        let (instructions, map) = parse_input(&input);
        part2(instructions, map)
    }

    #[ignore]
    #[case("ex3.txt")]
    fn test_steps_to_next_z(input: &str) {
        let input = std::fs::read_to_string(format!("src/day8/{}", input)).unwrap();
        let (instructions, map) = parse_input(&input);
        let res = instructions_to_z(&instructions, &map);
        for r in res {
            dbg!(r);
        }
    }
}
