#![allow(dead_code)]

use std::collections::HashSet;

type Rows = Vec<Row>;

#[derive(Debug, Clone)]
struct Row {
    springs: Vec<String>,
    damaged_counts: Vec<u64>,
    string: String,
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
        let springs = springs.split('.').map(|str| str.to_string()).collect();

        rows.push(Row { springs, damaged_counts: counts, string: String::new() });
    }

    rows
}

enum Status {
    Done,
    Invalid,
    Partial,
}

impl Row {
    fn check(&mut self) -> Status {
        if self.springs.is_empty() && self.damaged_counts.is_empty() {
            return Status::Done;
        }
        if self.springs.is_empty() || self.damaged_counts.is_empty() {
            return Status::Invalid;
        }

        // if !self.springs.first().unwrap().contains('?') {
        //     let group = self.springs.remove(0);
        //     let count = self.damaged_counts.remove(0);
        //     if group.len() != count as usize {
        //         return Status::Invalid;
        //     }
        // }

        if !self.springs.last().unwrap().contains('?') {
            let group = self.springs.pop().unwrap();
            let count = self.damaged_counts.pop().unwrap();
            if group.len() != count as usize {
                return Status::Invalid;
            }
            self.string = format!("{}.{}", self.string, group)
        }

        if self.springs.is_empty() {
            if self.damaged_counts.is_empty() {
                Status::Done
            } else {
                Status::Invalid
            }
        } else {
            Status::Partial
        }
    }

    fn permutations(&mut self) -> u64 {
        let mut perms: HashSet<String> = HashSet::new();

        let mut queue: Vec<Row> = Vec::new();
        queue.push(self.clone());

        while let Some(mut row) = queue.pop() {
            // println!("{:?}", row.springs);
            match row.check() {
                Status::Done => _ = perms.insert(dbg!(row.string)),
                Status::Invalid => continue,
                Status::Partial => {
                    if !row.springs.last().unwrap().contains('?') {
                        queue.push(row);
                        continue;
                    }
                    {
                        let mut row = row.clone();
                        let group = row.springs.pop().unwrap();
                        if let Some(idx) = group.find('?') {
                            // replace ? with . and split on .
                            if !group[..idx].is_empty() {
                                row.springs.push(group[..idx].to_string());
                            }
                            if !group[idx + 1..].is_empty() {
                                row.springs.push(group[idx + 1..].to_string());
                            }
                        } else {
                            row.springs.push(group);
                        }
                        queue.push(row);
                    }
                    {
                        let mut row = row.clone();
                        let group = row.springs.pop().unwrap();
                        let group = group.replacen('?', "#", 1);
                        row.springs.push(group);
                        queue.push(row);
                    }
                }
            }
        }

        dbg!(&perms);
        perms.len() as u64
    }
}

fn part_2(mut rows: Rows) -> u64 {
    let mut sum = 0;
    for row in &mut rows {
        sum += row.permutations();
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

    #[case("ex1.txt" => 525152)]
    #[case("input.txt" => 7732) ]
    fn test_part_2(file: &str) -> u64 {
        let rows = get_input(file);
        part_2(rows)
    }
}
