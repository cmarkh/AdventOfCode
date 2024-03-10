#![allow(dead_code)]

use crate::{Row, Rows};

fn diffs(digits: &[i64]) -> Vec<i64> {
    let mut row = Vec::with_capacity(digits.len() - 1);

    let mut digits = digits.iter();
    let mut prev = digits.next().unwrap();
    for next in digits {
        row.push(next - prev);
        prev = next;
    }

    row
}

impl Row {
    fn calculate_diffs(&mut self) {
        self.calcualted.push(diffs(&self.original));
        loop {
            let row = &self.calcualted.last().unwrap();
            if row.iter().all(|digit| *digit == 0) {
                break;
            }
            self.calcualted.push(diffs(row));
        }
    }

    fn sum_new_digits(&self) -> i64 {
        let mut sum = 0;
        for row in self.calcualted.iter().rev() {
            sum += row.last().unwrap();
        }
        sum + self.original.last().unwrap()
    }
}

pub fn part1(rows: &mut Rows) -> i64 {
    let mut sum = 0;
    for row in &mut rows.0 {
        row.calculate_diffs();
        sum += row.sum_new_digits();
    }
    sum
}

#[cfg(test)]
mod test {
    use super::*;
    use crate::test::get_input;

    use test_case::case;

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn test_diffs(file_name: &str) {
        let rows = get_input(file_name);
        let diffs = diffs(&rows.0[0].original);
        dbg!(diffs);
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn test_row_diffs(file_name: &str) {
        let mut rows = get_input(file_name);
        rows.0[0].calculate_diffs();
        dbg!(&rows.0[0].calcualted);
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn test_sum_new_nums(file_name: &str) {
        let mut rows = get_input(file_name);
        rows.0[0].calculate_diffs();
        dbg!(&rows.0[0].sum_new_digits());
    }

    #[ignore]
    #[case("ex1.txt" => 114)]
    #[case("input.txt" => 1877825184)]
    fn test_part_1(file_name: &str) -> i64 {
        let mut rows = get_input(file_name);
        let res = part1(&mut rows);
        dbg!(res);
        res
    }
}
