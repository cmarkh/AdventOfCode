#![allow(dead_code)]

use std::collections::{HashMap, HashSet};

type Rows = Vec<Row>;

#[derive(Debug, Clone, Hash, Eq, PartialEq)]
struct Row {
    springs: String,
    damaged_counts: Vec<u64>,
}

impl std::fmt::Display for Row {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let damaged_counts: Vec<String> = self.damaged_counts.iter().map(|&count| count.to_string()).collect();
        write!(f, "{} {}", self.springs, damaged_counts.join(","))
    }
}

fn parse_input(input: &str) -> Rows {
    let mut rows = Vec::new();

    for line in input.lines() {
        let (base_springs, counts) = line.split_once(' ').unwrap();

        let counts: Vec<u64> = counts.split(',').map(|ch| ch.parse().unwrap()).collect();
        let counts = counts.repeat(5);

        let mut springs = base_springs.to_string();
        for _ in 0..4 {
            springs += "?";
            springs += base_springs;
        }
        // let springs = springs.split('.').filter(|str| !str.is_empty()).map(|str| str.to_string()).collect();

        rows.push(Row { springs, damaged_counts: counts });
    }

    rows
}

fn lookup(springs: String) -> HashMap<String, Vec<String>> {
    let mut lookup: HashMap<String, Vec<String>> = HashMap::new();
    let mut queue = vec![springs];

    while let Some(sprigs) = queue.pop() {
        if lookup.contains_key(&sprigs) {
            continue;
        }
        let mut new_lookups = Vec::new();
        for idx in 0..sprigs.len() {
            if sprigs.chars().nth(idx).unwrap() != '?' {
                {
                    let mut new_springs = sprigs.clone();
                    new_springs.replace_range(idx..(idx + 1), ".");
                    queue.push(new_springs.clone());
                    new_lookups.push(new_springs);
                }
                {
                    let mut new_springs = sprigs.clone();
                    new_springs.replace_range(idx..(idx + 1), "#");
                    queue.push(new_springs.clone());
                    new_lookups.push(new_springs);
                }
            }
        }
        lookup.insert(sprigs, new_lookups);
    }

    lookup
}

fn part_2(rows: Rows) -> u64 {
    let mut sum = 0;
    for row in rows {
        // sum += line_permutations(row);
        let lookup = lookup(row.springs);
        for (key, val) in lookup {
            // println!("{}: {:?}", key, val);
        }
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
            println!("{:?}", row);
        }
    }

    #[case("???.### 1,1,3" => 1)]
    #[case(".??..??...?##. 1,1,3" => 16384) ]
    #[case("?#?#?#?#?#?#?#? 1,3,1,6" => 1) ]
    #[case("????.#...#... 4,1,1" => 16) ]
    #[case("????.######..#####. 1,6,5" => 2500) ]
    #[case("?###???????? 3,2,1" => 506250) ]
    fn test_lookup(line: &str) -> u64 {
        let rows = parse_input(line);
        part_2(rows)
    }

    #[case("ex1.txt" => 525152)]
    // #[case("input.txt" => 7732) ]
    fn test_part_2(file: &str) -> u64 {
        let rows = get_input(file);
        part_2(rows)
    }
}
