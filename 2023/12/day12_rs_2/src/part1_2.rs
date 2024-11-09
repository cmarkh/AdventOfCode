#![allow(unused)]

use std::collections::HashSet;

type Rows = Vec<Row>;
type Springs = String;

#[derive(Debug, Clone)]
struct Row {
    springs: Springs,
    damaged_counts: Vec<i32>,
}

fn parse_input(input: &str) -> Rows {
    let mut rows = Vec::new();

    for line in input.lines() {
        let (base_springs, counts) = line.split_once(' ').unwrap();

        let counts: Vec<i32> = counts.split(',').map(|ch| ch.parse().unwrap()).collect();
        // let springs: Vec<char> = base_springs.chars().collect();
        let springs = base_springs.to_string();

        rows.push(Row { springs, damaged_counts: counts });
    }

    rows
}

fn permutations(row: &Row) -> HashSet<Springs> {
    let counts = row.damaged_counts.clone();
    let len = row.springs.len();

    let mut perms = HashSet::new();

    let mut shorty = String::new();
    for (i, c) in counts.iter().enumerate() {
        if i != 0 {
            shorty.push('.');
        }
        shorty.push_str(&"#".repeat(*c as usize));
    }

    let mut tested = HashSet::new();

    let mut buffer = vec![shorty.clone()];
    buffer.push(format!(".{}", shorty));
    buffer.push(format!("{}.", shorty));
    buffer.push(format!(".{}.", shorty));

    while let Some(string) = buffer.pop() {
        if string.len() == len {
            perms.insert(string);
            continue;
        }
        if string.len() > len {
            continue;
        }

        let indicies = string.match_indices('.').collect::<Vec<_>>();
        for (idx, _) in indicies {
            if !test_permutation(&row.springs[..idx], &string[..idx]) {
                continue;
            }
            let new_string = format!("{}..{}", &string[..idx], &string[idx + 1..]);
            if !tested.contains(&new_string) {
                tested.insert(new_string.clone());
            } else {
                continue;
            }
            buffer.push(new_string);
        }
    }

    perms
}

fn test_permutation(springs: &str, perm: &str) -> bool {
    springs.chars().zip(perm.chars()).all(|(s, p)| s == p || s == '?')
}

fn row_permutations(row: &Row) -> usize {
    let mut count = 0;

    let perms = permutations(row);
    for perm in perms {
        if test_permutation(&row.springs, &perm) {
            println!("{}", perm);
            count += 1;
        }
    }

    count
}

fn part1(rows: Rows) -> usize {
    let mut count = 0;
    for row in rows {
        dbg!(&row);
        count += dbg!(row_permutations(&row));
    }
    count
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
    #[case("input.txt")]
    fn print_input(file: &str) {
        let rows = get_input(file);
        for row in &rows {
            println!("{:?}", row);
        }
    }

    #[case("ex1.txt")]
    fn test_permutations(file: &str) {
        let rows = get_input(file);
        let perms = permutations(&rows[1]);
        for perm in perms {
            println!("{}", perm);
        }
    }

    #[case("ex1.txt", 4)]
    fn test_row(file: &str, expected: usize) {
        let rows = get_input(file);
        let count = row_permutations(&rows[1]);
        assert_eq!(count, expected);
    }

    #[case("ex1.txt", 21)]
    #[case("input.txt", 7732)]
    fn test_part_1(file: &str, expected: usize) {
        let rows = get_input(file);
        let count = part1(rows);
        assert_eq!(count, expected);
    }
}
