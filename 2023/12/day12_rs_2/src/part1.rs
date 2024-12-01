type Rows = Vec<Row>;

#[derive(Debug, Clone)]
struct Row {
    springs: Vec<char>,
    damaged_counts: Vec<i32>,
}

fn parse_input(input: &str) -> Rows {
    let mut rows = Vec::new();

    for line in input.lines() {
        let (base_springs, counts) = line.split_once(' ').unwrap();

        let counts: Vec<i32> = counts.split(',').map(|ch| ch.parse().unwrap()).collect();
        let springs: Vec<char> = base_springs.chars().collect();

        rows.push(Row { springs, damaged_counts: counts });
    }

    rows
}

fn permute_row(row: Row) -> u32 {
    let mut permutations = 0;

    let mut buffer = vec![row];

    while !buffer.is_empty() {
        let row = buffer.pop().unwrap();

        let mut idx = None;
        for (i, char) in row.springs.iter().enumerate() {
            if *char == '?' {
                idx = Some(i);
                break;
            }
        }

        match idx {
            Some(idx) => {
                let mut row1 = row.clone();
                row1.springs[idx] = '#';
                buffer.push(row1);

                let mut row2 = row;
                row2.springs[idx] = '.';
                buffer.push(row2);
            }
            None => match test_row(row) {
                true => permutations += 1,
                false => (),
            },
        }
    }

    permutations
}

fn test_row(row: Row) -> bool {
    let mut counts = row.damaged_counts.clone();
    let mut counts = counts.iter_mut();
    let springs = row.springs.iter();

    let mut count = counts.next().unwrap();
    for spring in springs {
        if *spring == '?' {
            panic!("? found in row");
        }
        if *spring == '.' && *count != 0 {
            return false;
        }
        if *spring == '.' && *count == 0 {
            count = match counts.next() {
                Some(c) => c,
                None => return false,
            };
        }
        if *spring == '#' {
            *count -= 1;
        }
        if *count < 0 {
            return false;
        }
    }

    true
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

    #[case("ex1.txt", 2)]
    #[case("input.txt", 0)]
    fn test_permute_row(file: &str, expected: u32) {
        let rows = get_input(file);
        let count = permute_row(rows[0].clone());
        assert_eq!(count, expected);
    }
}
