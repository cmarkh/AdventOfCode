#![allow(dead_code)]

type Rows = Vec<Row>;

#[derive(Debug, Clone)]
struct Row {
    springs: String,
    damaged_counts: Vec<u64>,
}

fn parse_input(input: &str) -> Rows {
    let mut rows = Vec::new();

    for line in input.lines() {
        let (springs, counts) = line.split_once(' ').unwrap();

        let counts: Vec<u64> = counts.split(',').map(|ch| ch.parse().unwrap()).collect();
        rows.push(Row { springs: springs.to_string(), damaged_counts: counts });
    }

    rows
}

impl Row {
    fn permutations(&self) -> u64 {
        let mut perms = 0;

        let mut queue: Vec<Row> = Vec::new();
        queue.push(self.clone());

        while let Some(row) = queue.pop() {
            if !row.springs.contains('?') {
                if row.valid() {
                    perms += 1;
                }
                continue;
            }

            {
                let mut row2 = row.clone();
                row2.springs = row2.springs.replacen('?', ".", 1);
                queue.push(row2);
            }
            {
                let mut row2 = row.clone();
                row2.springs = row2.springs.replacen('?', "#", 1);
                queue.push(row2);
            }
        }

        perms
    }

    fn valid(&self) -> bool {
        let groups: Vec<&str> = self.springs.split('.').filter(|group| !group.is_empty()).collect();
        if groups.len() != self.damaged_counts.len() {
            return false;
        }

        for (group, count) in groups.iter().zip(self.damaged_counts.iter()) {
            if group.len() as u64 != *count {
                return false;
            }
        }

        true
    }
}

fn part_1(rows: Rows) -> u64 {
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
    #[case("input.txt")]
    fn print_input(file: &str) {
        let rows = get_input(file);
        for row in &rows {
            println!("{:?}", row);
        }
    }

    #[test]
    fn test_valid_row_1() {
        let row = Row { springs: "#.#.###".to_string(), damaged_counts: vec![1, 1, 3] };
        assert!(row.valid());
    }

    #[test]
    fn test_valid_row_2() {
        let row = Row { springs: "###.###".to_string(), damaged_counts: vec![1, 1, 3] };
        assert!(!row.valid());
    }

    #[test]
    fn test_valid_row_3() {
        let row = Row { springs: "#.#.##".to_string(), damaged_counts: vec![1, 1, 3] };
        assert!(!row.valid());
    }

    #[test]
    fn test_permutations_1() {
        let row = Row { springs: "???.###".to_string(), damaged_counts: vec![1, 1, 3] };
        let perms = row.permutations();
        dbg!(&perms);
        assert_eq!(perms, 1);
    }

    #[test]
    fn test_permutations_2() {
        let row = Row { springs: ".??..??...?##.".to_string(), damaged_counts: vec![1, 1, 3] };
        let perms = row.permutations();
        dbg!(&perms);
        assert_eq!(perms, 4);
    }

    #[case("ex1.txt" => 21)]
    #[case("input.txt" => 7732) ]
    fn test_part_1(file: &str) -> u64 {
        let rows = get_input(file);
        part_1(rows)
    }
}
