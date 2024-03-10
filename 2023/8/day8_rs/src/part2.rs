#![allow(dead_code)]

use std::{collections::HashMap, hash::Hash};

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

#[derive(Debug)]
struct Line<'a> {
    current_position: &'a str,
    current_instruction: usize,
    steps: u64,
    map: &'a Map,
    instructions: &'a Instructions,
}

impl<'a> Line<'a> {
    fn new(start: &'a str, instructions: &'a Instructions, map: &'a Map) -> Self {
        Self {
            current_position: start,
            current_instruction: 0,
            steps: 0,
            map,
            instructions,
        }
    }

    fn step_one(&mut self) {
        self.current_position = match self.instructions.0.get(self.current_instruction).unwrap() {
            Instruction::Left => &self.map.0.get(self.current_position).unwrap().0,
            Instruction::Right => &self.map.0.get(self.current_position).unwrap().1,
        };
        self.current_instruction = (self.current_instruction + 1) % self.instructions.0.len();
        self.steps += 1;
    }

    fn steps_to_z_iter(&mut self) -> u64 {
        let mut steps = 0;
        while !self.current_position.ends_with('Z') || steps == 0 {
            self.step_one();
            steps += 1;
        }
        steps
    }

    fn step_to_z(&mut self, steps_to_z: &mut StepsToZ<'a>) {
        let key = (self.current_position, self.current_instruction);
        if let Some(entry) = steps_to_z.get_mut(&key) {
            self.steps += entry.0;
            self.current_position = entry.1;
            self.current_instruction = entry.2;
        } else {
            let steps = self.steps_to_z_iter();
            let value = (steps, self.current_position, self.current_instruction);
            steps_to_z.insert(key, value);
        }
    }
}

/// Map<(Start Position, Start Instruction Index), (Steps, End Position, End Instruction Index)>
type StepsToZ<'a> = HashMap<(&'a str, usize), (u64, &'a str, usize)>;

#[derive(Debug)]
struct Lines<'a> {
    lines: Vec<Line<'a>>,
    steps_to_z: StepsToZ<'a>,
}

impl<'a> Lines<'a> {
    fn new(instructions: &'a Instructions, map: &'a Map) -> Self {
        let starts = map.starts();
        let lines = starts
            .into_iter()
            .map(|start| Line::new(start, instructions, map))
            .collect();
        Lines {
            lines,
            steps_to_z: HashMap::new(),
        }
    }

    fn step_all(&mut self) -> bool {
        let max_step = self
            .lines
            .iter()
            .max_by_key(|line| line.steps)
            .unwrap()
            .steps;

        let equal = self.lines.iter().all(|line| line.steps == max_step);

        for line in &mut self.lines {
            if line.steps < max_step || equal {
                line.step_to_z(&mut self.steps_to_z);
                // println!(
                //     "steps: {}\tpossition: {}\tinstruction: {}",
                //     line.steps, line.current_position, line.current_instruction
                // );
            }
        }

        let steps_test = self.lines[0].steps;
        for line in &self.lines {
            if line.steps != steps_test || !line.current_position.ends_with('Z') {
                return false;
            }
        }
        true
    }
}

fn part_2(instructions: Instructions, map: Map) -> u64 {
    let mut lines = Lines::new(&instructions, &map);

    let mut done = false;
    while !done {
        done = lines.step_all();
    }

    lines.lines[0].steps
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
    #[case("ex3.txt")]
    #[case("input.txt")]
    fn test_input(file_name: &str) {
        let (instructions, map) = get_input(file_name);
        dbg!(instructions, map);
    }

    #[ignore]
    #[case("ex3.txt")]
    #[case("input.txt")]
    fn test_run(file_name: &str) {
        let (instructions, map) = get_input(file_name);
        let starts = map.starts();
        let start = starts.first().unwrap();
        let mut line = Line::new(start, &instructions, &map);

        while !line.current_position.ends_with('Z') {
            line.step_one();
            dbg!((line.current_position, line.steps));
        }
    }

    #[ignore]
    #[case("ex3.txt")]
    #[case("input.txt")]
    fn test_run_2(file_name: &str) {
        let (instructions, map) = get_input(file_name);
        let ends = map.ends();
        let end = ends.first().unwrap();
        let mut line = Line::new(end, &instructions, &map);

        while !line.current_position.ends_with('Z') || line.steps == 0 {
            line.step_one();
            dbg!((line.current_position, line.steps));
        }
    }

    #[ignore]
    #[case("ex3.txt")]
    #[case("input.txt")]
    fn test_run_3(file_name: &str) {
        let (instructions, map) = get_input(file_name);
        let ends = map.ends();
        let end = ends.first().unwrap();
        let mut line = Line::new(end, &instructions, &map);

        let mut z_steps = Vec::new();
        for i in 0..u64::MAX {
            line.step_one();
            if line.current_position.ends_with('Z') {
                z_steps.push(i);
                dbg!((line.current_position, i));
            }
        }
    }

    #[ignore]
    #[case("ex3.txt")]
    #[case("input.txt")]
    fn test_steps_to_z(file_name: &str) {
        let (instructions, map) = get_input(file_name);
        let mut lines = Lines::new(&instructions, &map);
        lines.step_all();
        for line in &lines.lines {
            dbg!((line.steps, line.current_position));
        }
    }

    #[case("ex3.txt" => 6)]
    #[case("input.txt" => 13289612809129)]
    fn test_part_2(file_name: &str) -> u64 {
        let (instructions, map) = get_input(file_name);
        let steps = part_2(instructions, map);
        dbg!(steps);
        steps
    }
}
