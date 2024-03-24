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

fn permutations(springs: &str, damaged_counts: &[u64]) -> u64 {
    let mut perms = 0;

    let mut queue: Vec<String> = Vec::new();
    queue.push(springs.to_string());

    while let Some(springs) = queue.pop() {
        if !springs.contains('?') {
            if row_valid(&springs, damaged_counts) {
                perms += 1;
            }
            continue;
        }

        {
            let springs2 = springs.replacen('?', ".", 1);
            if row_partial_valid(&springs2, damaged_counts) {
                queue.push(springs2);
            }
        }
        {
            let springs2 = springs.replacen('?', "#", 1);
            if row_partial_valid(&springs2, damaged_counts) {
                queue.push(springs2);
            }
        }
    }

    perms
}

fn row_valid(springs: &str, damaged_counts: &[u64]) -> bool {
    let groups = springs.split('.').filter(|group| !group.is_empty());
    if groups.clone().count() != damaged_counts.len() {
        return false;
    }

    for (group, count) in groups.zip(damaged_counts.iter()) {
        if group.len() as u64 != *count {
            return false;
        }
    }

    true
}

fn row_partial_valid(springs: &str, damaged_counts: &[u64]) -> bool {
    let groups = springs.split('.').filter(|group| !group.is_empty());

    for (group, count) in groups.zip(damaged_counts.iter()) {
        if group.contains('?') {
            return true;
        }
        if group.len() as u64 != *count {
            return false;
        }
    }

    true
}

fn part_2(rows: Rows) -> u64 {
    let mut sum = 0;
    for row in &rows {
        let mut row_sum = 1;

        let spring = row.springs.clone();
        let springq = format!("{}?", row.springs.clone());
        let qspring = format!("?{}", row.springs.clone());
        let qspringq = format!("?{}?", row.springs.clone());

        {
            // part 1
            let mut part_sum = 0;
            part_sum += permutations(&spring, &row.damaged_counts);
            part_sum += permutations(&springq, &row.damaged_counts);
            row_sum *= part_sum;
        }
        {
            // part 2,3,4
            let mut part_sum = 0;
            part_sum += permutations(&spring, &row.damaged_counts);
            part_sum += permutations(&springq, &row.damaged_counts);
            part_sum += permutations(&qspring, &row.damaged_counts);
            part_sum += permutations(&qspringq, &row.damaged_counts);
            row_sum *= part_sum * 3;
        }
        {
            // part 5
            let mut part_sum = 0;
            part_sum += permutations(&spring, &row.damaged_counts);
            part_sum += permutations(&qspring, &row.damaged_counts);
            row_sum *= part_sum;
        }

        sum += row_sum
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
        let row = Row { springs: "???.###".to_string(), damaged_counts: vec![1, 1, 3] };
        let perms = permutations(&row.springs, &row.damaged_counts);
        dbg!(&perms);
        assert_eq!(perms, 1);
    }

    #[test]
    fn test_permutations_2() {
        let row = Row { springs: ".??..??...?##.".to_string(), damaged_counts: vec![1, 1, 3] };
        let perms = permutations(&row.springs, &row.damaged_counts);
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
