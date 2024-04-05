#![allow(dead_code)]

type Rows = Vec<Row>;

#[derive(Debug, Clone)]
struct Row {
    springs: Vec<String>,
    damaged_counts: Vec<u64>,
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
        let springs = springs.split('.').filter(|str| !str.is_empty()).map(|str| str.to_string()).collect();

        rows.push(Row { springs, damaged_counts: counts });
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
        if let Some(spring) = self.springs.first() {
            if !spring.contains('?') {
                let group = self.springs.remove(0);
                let count = self.damaged_counts.remove(0);
                if group.len() != count as usize {
                    return Status::Invalid;
                }
            }
        }

        if let Some(spring) = self.springs.last() {
            if !spring.contains('?') {
                let group = self.springs.pop().unwrap();
                let count = self.damaged_counts.pop().unwrap();
                if group.len() != count as usize {
                    return Status::Invalid;
                }
            }
        }

        match (self.springs.is_empty(), self.damaged_counts.is_empty()) {
            (true, true) => Status::Done,
            (true, false) => Status::Invalid,
            (false, true) => {
                for group in &self.springs {
                    if group.contains('#') {
                        return Status::Invalid;
                    }
                }
                Status::Done // remaining ? can be all .
            }
            (false, false) => Status::Partial,
        }
    }

    fn permutations(&mut self) -> u64 {
        let mut perms = 0;

        let mut queue: Vec<Row> = Vec::new();
        queue.push(self.clone());

        while let Some(mut row) = queue.pop() {
            // println!("{:?}", row.springs);
            match row.check() {
                Status::Done => perms += 1,
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
                        if let Some(group) = row.springs.last_mut() {
                            *group = group.replacen('?', "#", 1);
                        }
                        queue.push(row);
                    }
                }
            }
        }

        perms
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

    #[case("???.### 1,1,3" => 1)]
    #[case(".??..??...?##. 1,1,3" => 16384) ]
    #[case("?#?#?#?#?#?#?#? 1,3,1,6" => 1) ]
    #[case("????.#...#... 4,1,1" => 16) ]
    #[case("????.######..#####. 1,6,5" => 2500) ]
    #[case("?###???????? 3,2,1" => 506250) ]
    fn test_line(line: &str) -> u64 {
        let rows = parse_input(line);
        part_2(rows)
    }

    #[case("ex1.txt" => 525152)]
    #[case("input.txt" => 7732) ]
    fn test_part_2(file: &str) -> u64 {
        let rows = get_input(file);
        part_2(rows)
    }
}
