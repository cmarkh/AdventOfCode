#![allow(dead_code)]

type Rows = Vec<Row>;

#[derive(Debug, Clone)]
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

        rows.push(Row { springs, damaged_counts: counts });
    }

    rows
}

fn line_permutations(mut row: Row) -> u64 {
    let mut permutations = 0;

    row.damaged_counts = row.damaged_counts.into_iter().rev().collect();
    let mut queue = vec![row];

    'next_row: while let Some(mut row) = queue.pop() {
        println!("{}", row);

        // match (row.springs.is_empty(), row.damaged_counts.is_empty()) {
        //     (true, true) => {
        //         permutations += 1;
        //         continue;
        //     }
        //     (true, false) => {}
        //     (false, true) => continue,
        //     (false, false) => {}
        // }
        if row.damaged_counts.len() > row.springs.len() {
            continue;
        }

        // can skip . in the beginning
        let mut springs = row.springs.chars();
        let mut remaining_springs = String::new();
        #[allow(clippy::while_let_on_iterator)]
        while let Some(ch) = springs.next() {
            match ch {
                '.' => continue,
                ch => {
                    remaining_springs.push(ch);
                    remaining_springs += &springs.collect::<String>();
                    break;
                }
            }
        }
        let mut springs = remaining_springs.chars();

        // replace ? with options # and .
        if springs.next() == Some('?') {
            let springs = springs.collect::<String>().clone();
            queue.push(Row { springs: format!("#{}", springs.clone()), damaged_counts: row.damaged_counts.clone() });
            queue.push(Row { springs, damaged_counts: row.damaged_counts });
            continue 'next_row;
        }

        let count = match row.damaged_counts.pop() {
            Some(count) => count,
            None => {
                if !row.springs.contains('#') {
                    permutations += 1;
                }
                continue 'next_row;
            }
        };

        for _ in 0..count {
            match springs.next() {
                Some('#') => {}
                Some('?') => {}                  // must be a #
                Some('.') => continue 'next_row, // invalid
                None => continue 'next_row,      // invalid
                _ => unreachable!(),
            }
        }
        queue.push(Row { springs: springs.collect::<String>(), damaged_counts: row.damaged_counts });
    }

    permutations
}

fn part_2(rows: Rows) -> u64 {
    let mut sum = 0;
    for row in rows {
        sum += line_permutations(row);
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
