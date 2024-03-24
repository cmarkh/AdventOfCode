#![allow(dead_code)]

type Rows = Vec<Row>;

#[derive(Debug, Clone)]
struct Row {
    springs: String,
    damaged_counts: Vec<u64>,
    unknown_indices: Vec<usize>,
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

        let springs = springs.to_string();
        let mut unknown_indices: Vec<usize> =
            springs.chars().enumerate().filter(|(_, ch)| ch == &'?').map(|(i, _)| i).collect();
        unknown_indices.reverse();
        rows.push(Row { springs, damaged_counts: counts, unknown_indices });
    }

    rows
}

impl Row {
    fn permutations(&self) -> u64 {
        let mut perms = 0;

        let damaged_counts = self.damaged_counts.clone();

        let mut queue: Vec<Row> = Vec::new();
        queue.push(self.clone());

        while let Some(row) = queue.pop() {
            if !row.springs.contains('?') {
                if row_valid(&row.springs, &damaged_counts) {
                    perms += 1;
                }
                continue;
            }

            {
                let mut row2 = row.clone();
                let idx = row2.unknown_indices.pop().unwrap();
                row2.springs.replace_range(idx..(idx + 1), ".");
                if row_partial_valid(&row2.springs, &damaged_counts) {
                    queue.push(row2);
                }
            }
            {
                let mut row2 = row.clone();
                let idx = row2.unknown_indices.pop().unwrap();
                row2.springs.replace_range(idx..(idx + 1), "#");
                if row_partial_valid(&row2.springs, &damaged_counts) {
                    queue.push(row2);
                }
            }
        }

        perms
    }
}

fn row_valid(springs: &str, damaged_counts: &[u64]) -> bool {
    let group_lengths: Vec<u64> =
        springs.split('.').filter(|&group| (!group.is_empty())).map(|group| group.len() as u64).collect();

    group_lengths == damaged_counts
}

fn row_partial_valid(springs: &str, damaged_counts: &[u64]) -> bool {
    let mut count_iter = damaged_counts.iter();
    for group in springs.split('.').filter(|g| !g.is_empty()) {
        if group.contains('?') {
            return true;
        }
        match count_iter.next() {
            Some(&expected_count) => {
                if group.len() as u64 != expected_count {
                    return false;
                }
            }
            None => return false,
        }
    }
    count_iter.next().is_none()
}

fn part_2(rows: Rows) -> u64 {
    let mut sum = 0;
    for row in &rows {
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

    #[case("#.#.###", &[1, 1, 3])]
    fn test_row_valid(springs: &str, damaged_counts: &[u64]) {
        assert!(row_valid(springs, damaged_counts));
    }

    #[case("###.###", &[1, 1, 3])]
    fn test_row_not_valid(springs: &str, damaged_counts: &[u64]) {
        assert!(!row_valid(springs, damaged_counts));
    }

    #[test]
    fn test_permutations_1() {
        let row = Row { springs: "???.###".to_string(), damaged_counts: vec![1, 1, 3], unknown_indices: vec![0, 1, 2] };
        let perms = row.permutations();
        dbg!(&perms);
        assert_eq!(perms, 1);
    }

    #[test]
    fn test_permutations_2() {
        let row = Row {
            springs: ".??..??...?##.".to_string(),
            damaged_counts: vec![1, 1, 3],
            unknown_indices: vec![1, 2, 5, 6, 10],
        };
        let perms = row.permutations();
        dbg!(&perms);
        assert_eq!(perms, 4);
    }

    #[case("ex1.txt" => 525152)]
    #[case("input.txt" => 7732) ]
    fn test_part_2(file: &str) -> u64 {
        let rows = get_input(file);
        part_2(rows)
    }
}
