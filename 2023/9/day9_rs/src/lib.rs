#![allow(dead_code)]

pub mod part1;
pub mod part2;

#[derive(Debug)]
pub struct Rows(Vec<Row>);

#[derive(Debug)]
pub struct Row {
    pub original: Vec<i64>,
    pub calcualted: Vec<Vec<i64>>,
}

pub fn parse_input(input: &str) -> Rows {
    let mut rows = Rows(Vec::new());

    for line in input.lines() {
        let mut row = Row {
            original: Vec::new(),
            calcualted: Vec::new(),
        };
        for digit in line.split(' ') {
            let digit: i64 = digit.parse().unwrap();
            row.original.push(digit);
        }
        rows.0.push(row);
    }

    rows
}

#[cfg(test)]
mod test {
    use super::*;

    use test_case::case;

    pub fn get_input(file_name: &str) -> Rows {
        let input = std::fs::read_to_string(format!("src/{}", file_name)).unwrap();
        parse_input(&input)
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn test_input(file_name: &str) {
        let rows = get_input(file_name);
        dbg!(rows);
    }
}
