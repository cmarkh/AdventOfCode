#![allow(dead_code)]

use std::collections::HashSet;

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
        write!(f, "{}  \t\t{}", springs, damaged_counts.join(","))
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

#[derive(Debug, Clone)]
enum Status {
    Valid,
    Partial(Row),
    Invalid,
}

impl Row {
    fn test(&self) -> Status {
        let groups = self.springs.split(|&ch| ch == '.').filter(|group| !group.is_empty());

        for (i, (group, expected_damage)) in groups.clone().zip(&self.damaged_counts).enumerate() {
            let damaged_count = group.iter().filter(|&&ch| ch == '#').count() as u64;
            if damaged_count != *expected_damage || group.contains(&'?') {
                let damaged_counts = self.damaged_counts[i..].to_vec();
                let groups = groups.collect::<Vec<_>>()[i..].to_vec();
                if groups.is_empty() || damaged_counts.is_empty() {
                    return Status::Invalid;
                }

                let springs =
                    groups.iter().map(|group| group.iter().collect::<String>()).collect::<Vec<String>>().join(".");

                let row = Row { springs: springs.chars().collect(), damaged_counts };
                // println!("{}  Partial", row);
                return Status::Partial(row);
            }
        }

        match groups.count() == self.damaged_counts.len() {
            true => Status::Valid,
            false => Status::Invalid,
        }
    }
}

fn permutations(row: Row) -> usize {
    let mut count = 0;
    let mut history = HashSet::new();
    let mut queue = vec![row];

    let mut i = 0;
    while let Some(row) = queue.pop() {
        i += 1;
        // println!("{row}");

        // if i > 40 {
        //     break;
        // }

        if history.contains(&row) {
            continue;
        } else {
            history.insert(row.clone());
        }

        // TODO: shift the ? split into here and cache that
        let mut row = match row.test() {
            Status::Valid => {
                count += 1;
                println!("{} Valid", row);
                continue;
            }
            Status::Partial(row) => row,
            Status::Invalid => continue,
        };

        let idx = match row.springs.iter().position(|&ch| ch == '?') {
            Some(idx) => idx,
            None => {
                queue.push(row);
                continue;
            }
        };
        {
            let mut springs2 = row.springs.clone();
            springs2[idx] = '.';
            queue.push(Row { springs: springs2, damaged_counts: row.damaged_counts.clone() });
        }
        row.springs[idx] = '#';
        queue.push(row);
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
    fn test_row(line: &str) -> usize {
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
