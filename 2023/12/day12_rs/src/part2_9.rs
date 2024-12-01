#![allow(dead_code)]

use std::collections::{HashMap, HashSet};

type Rows = Vec<Row>;

#[derive(Debug, Clone, Hash, Eq, PartialEq)]
struct Row {
    springs: Vec<char>,
    damaged_counts: Vec<u64>,
}

impl std::fmt::Display for Row {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let springs = self.springs.iter().collect::<String>();
        let damaged_counts: Vec<String> = self.damaged_counts.iter().map(|&count| count.to_string()).collect();
        write!(f, "{} {}", springs, damaged_counts.join(","))
    }
}

fn print_springs(springs: &[char]) {
    println!("{}", springs.iter().collect::<String>())
}

fn parse_input(input: &str) -> Rows {
    let mut rows = Vec::new();

    for line in input.lines() {
        let (base_springs, counts) = line.split_once(' ').unwrap();

        let counts: Vec<u64> = counts.split(',').map(|ch| ch.parse().unwrap()).collect();
        let counts = counts.repeat(5);

        let base_springs: Vec<char> = base_springs.chars().collect();
        let mut springs: Vec<char> = base_springs.clone();
        for _ in 0..4 {
            springs.push('?');
            springs.extend(base_springs.clone());
        }

        rows.push(Row { springs, damaged_counts: counts });
    }

    rows
}

fn springs_match_damaged_counts(springs: &[char], damaged_counts: &[u64]) -> bool {
    if springs.contains(&'?') {
        return false;
    }

    let groups = springs.split(|&ch| ch == '.');
    if groups.clone().count() != damaged_counts.len() {
        return false;
    }

    for (group, expected_damage) in groups.zip(damaged_counts) {
        let damaged_count = group.iter().filter(|&&ch| ch == '#').count() as u64;
        if damaged_count != *expected_damage {
            return false;
        }
    }
    true
}

fn permutations(row: Row) -> usize {
    let mut count = 0;
    let damaged_counts = row.damaged_counts;
    let mut history = HashSet::new();
    let mut queue = vec![row.springs];

    while let Some(mut springs) = queue.pop() {
        print_springs(springs.as_slice());

        if history.contains(&springs) {
            continue;
        } else {
            history.insert(springs.clone());
        }

        if !springs.contains(&'?') {
            if springs_match_damaged_counts(&springs, &damaged_counts) {
                count += 1;
            }
            continue;
        }

        let idx = springs.iter().position(|&ch| ch == '?').unwrap();
        {
            let mut springs2 = springs.clone();
            springs2[idx] = '.';
            queue.push(springs2);
        }
        springs[idx] = '#';
        queue.push(springs);
    }

    count
}

fn part_2(rows: Rows) -> usize {
    let mut sum = 0;
    for row in rows {
        sum += permutations(row);
    }
    sum
}

#[cfg(test)]
mod test {
    use test_case::case;

    use super::*;

    pub fn get_input(file_name: &str) -> Rows {
        let input = std::fs::read_to_string(format!("src/{}", file_name)).unwrap();
        parse_input(&input)
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("ex2.txt")]
    #[case("input.txt")]
    fn print_input(file: &str) {
        let rows = get_input(file);
        for row in &rows {
            println!("{}", row);
        }
    }

    #[case("???.### 1,1,3" => 1)]
    #[case(".??..??...?##. 1,1,3" => 16384) ]
    #[case("?#?#?#?#?#?#?#? 1,3,1,6" => 1) ]
    #[case("????.#...#... 4,1,1" => 16) ]
    #[case("????.######..#####. 1,6,5" => 2500) ]
    #[case("?###???????? 3,2,1" => 506250) ]
    fn test_lookup(line: &str) -> usize {
        let rows = parse_input(line);
        part_2(rows)
    }

    #[case("ex1.txt" => 525152)]
    // #[case("input.txt" => 7732) ]
    fn test_part_2(file: &str) -> usize {
        let rows = get_input(file);
        part_2(rows)
    }
}
