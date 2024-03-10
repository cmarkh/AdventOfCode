#![allow(dead_code)]

use std::{collections::HashMap, rc::Rc};

#[derive(Debug, PartialEq, PartialOrd, Eq, Ord, Clone, Copy, Hash)]
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
struct StepsToZ<'a>(HashMap<(&'a str, usize), (u64, &'a str)>);

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

    /// (steps, ending position)
    fn steps_to_z(
        &'a self,
        start_position: &'a str,
        instruction_index: usize,
        instructions: &Instructions,
    ) -> Option<(u64, &'a str)> {
        let mut position = start_position;

        dbg!(start_position);

        let mut history = std::collections::HashSet::new();

        for (steps, ins) in instructions
            .0
            .iter()
            .cycle()
            .skip(instruction_index)
            .enumerate()
        {
            // dbg!((position, steps, ins, &self.0.get(position).unwrap()));
            if position.ends_with('Z') && steps > 0 {
                return Some((steps as u64, position));
            }
            // cycle detection
            let key = (position, (instruction_index + steps) % instructions.0.len());
            if history.contains(&key) {
                return None;
            }
            history.insert(key);

            match ins {
                Instruction::Left => position = &self.0.get(position).unwrap().0,
                Instruction::Right => position = &self.0.get(position).unwrap().1,
            }
        }

        unreachable!()
    }

    fn steps_to_z_from_all_starts(&'a self, instructions: &Instructions) -> StepsToZ {
        let mut steps_to_z = StepsToZ(HashMap::new());

        for start in self.starts().into_iter().chain(self.ends()) {
            for i in 0..instructions.0.len() {
                if let Some(res) = self.steps_to_z(start, i, instructions) {
                    steps_to_z.0.insert((start, i), res);
                }
            }
        }

        steps_to_z
    }
}

#[derive(Debug)]
struct Line<'a> {
    current_position: &'a str,
    current_instruction: usize,
    steps: u64,
    map: Rc<Map>,
    instructions: Rc<Instructions>,
    steps_to_z: Rc<StepsToZ<'a>>,
}

impl<'a> Line<'a> {
    fn new(
        start: &'a str,
        instructions: Rc<Instructions>,
        map: Rc<Map>,
        steps_to_z: Rc<StepsToZ<'a>>,
    ) -> Self {
        Self {
            current_position: start,
            current_instruction: 0,
            steps: 0,
            map,
            instructions,
            steps_to_z,
        }
    }

    fn step(&mut self, current_max_steps: u64) {
        while self.steps < current_max_steps {
            dbg!(&(self.current_position, self.current_instruction));
            let (steps, end) = self
                .steps_to_z
                .0
                .get(&(self.current_position, self.current_instruction))
                .unwrap();
            self.steps += steps;
            self.current_position = end;
            self.current_instruction =
                (self.current_instruction + *steps as usize) % self.instructions.0.len();
        }
    }
}

struct Lines<'a>(Vec<Line<'a>>);

impl<'a> Lines<'a> {
    fn max_step(&self) -> u64 {
        self.0.iter().max_by_key(|line| line.steps).unwrap().steps
    }

    fn steps_all_equal(&self) -> (bool, u64) {
        let steps: Vec<u64> = self.0.iter().map(|line| line.steps).collect();
        let is_equal = steps.windows(2).all(|w| w[0] == w[1]);
        let max_steps = *steps.iter().max().unwrap();
        (is_equal, max_steps)
    }
}

fn part_2(instructions: Instructions, map: Map) -> u64 {
    let steps_to_z = map.steps_to_z_from_all_starts(&instructions);
    let instructions = Rc::new(instructions);
    let map = Rc::new(map.clone());
    let steps_to_z = Rc::new(steps_to_z);

    let starts = map.starts();

    let mut lines: Lines = Lines(
        starts
            .iter()
            .map(|start| Line::new(start, instructions.clone(), map.clone(), steps_to_z.clone()))
            .collect(),
    );

    let mut max_step = 1;

    loop {
        dbg!(max_step);
        for line in &mut lines.0 {
            line.step(max_step);
        }
        let (steps_equal, max_steps) = lines.steps_all_equal();
        match steps_equal {
            true => {
                for line in &lines.0 {
                    dbg!(line.current_position);
                }
                return max_steps;
            }
            false => max_step = max_steps,
        }
    }
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

    #[ignore]
    #[case("ex3.txt")]
    #[case("input.txt")]
    fn test_steps_to_z_all_starts(file_name: &str) {
        let (instructions, map) = get_input(file_name);
        let steps = map.steps_to_z_from_all_starts(&instructions);
        dbg!(steps);
    }

    #[case("ex3.txt" => 6)]
    #[case("input.txt" => 0)]
    fn test_part_2(file_name: &str) -> u64 {
        let (instructions, map) = get_input(file_name);
        let steps = part_2(instructions, map);
        dbg!(steps);
        steps
    }
}
