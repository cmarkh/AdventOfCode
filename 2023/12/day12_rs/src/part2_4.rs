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

fn rows_with_matched_counts(mut rows: Vec<Row>) -> Vec<Row> {
    let mut perms = Vec::new();

    while let Some(row) = rows.pop() {
        if row.springs.len() == row.damaged_counts.len() {
            perms.push(row)
        }
    }

    perms
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

    // #[case("???.### 1,1,3")]
    #[case(".??..??...?##. 1,1,3")]
    #[case("?#?#?#?#?#?#?#? 1,3,1,6")]
    #[case("????.#...#... 4,1,1")]
    #[case("????.######..#####. 1,6,5")]
    #[case("?###???????? 3,2,1")]
    fn test_rows_with_matched_counts(line: &str) {
        let rows = parse_input(line);
        let perms = rows_with_matched_counts(rows);
        for perm in perms {
            println!("{:?} {:?}", perm.springs, perm.damaged_counts)
        }
    }

    // #[case("???.### 1,1,3" => 1)]
    // #[case(".??..??...?##. 1,1,3" => 16384) ]
    // #[case("?#?#?#?#?#?#?#? 1,3,1,6" => 1) ]
    // #[case("????.#...#... 4,1,1" => 16) ]
    // #[case("????.######..#####. 1,6,5" => 2500) ]
    // #[case("?###???????? 3,2,1" => 506250) ]
    // fn test_line(line: &str) -> u64 {
    //     let rows = parse_input(line);
    //     part_2(rows)
    // }

    // #[case("ex1.txt" => 525152)]
    // // #[case("input.txt" => 7732) ]
    // fn test_part_2(file: &str) -> u64 {
    //     let rows = get_input(file);
    //     part_2(rows)
    // }
}
