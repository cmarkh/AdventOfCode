#![allow(dead_code)]

use crate::{Row, Rows};

fn diffs(digits: &[i64]) -> Vec<i64> {
    let mut row = Vec::with_capacity(digits.len() - 1);

    let mut digits = digits.iter().rev();
    let mut prev = digits.next().unwrap();
    for next in digits {
        row.push(prev - next);
        prev = next;
    }

    row.reverse();
    row
}

impl Row {
    fn calculate_diffs_2(&mut self) {
        self.calcualted.push(diffs(&self.original));
        loop {
            let row = &self.calcualted.last().unwrap();
            if row.iter().all(|digit| *digit == 0) {
                break;
            }
            self.calcualted.push(diffs(row));
        }
    }

    fn diff_new_digits(&self) -> i64 {
        let mut diff = 0;
        for row in self.calcualted.iter().rev() {
            diff = row.first().unwrap() - diff;
        }
        self.original.first().unwrap() - diff
    }
}

pub fn part2(rows: &mut Rows) -> i64 {
    let mut sum = 0;
    for row in &mut rows.0 {
        row.calculate_diffs_2();
        sum += row.diff_new_digits();
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
        let diffs = diffs(&rows.0[2].original);
        dbg!(diffs);
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn test_row_diffs(file_name: &str) {
        let mut rows = get_input(file_name);
        rows.0[2].calculate_diffs_2();
        dbg!(&rows.0[2].calcualted);
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn test_sum_new_nums(file_name: &str) {
        let mut rows = get_input(file_name);
        let row = rows.0.last_mut().unwrap();
        row.calculate_diffs_2();
        dbg!(&row.original);
        dbg!(&row.calcualted);
        dbg!(&row.diff_new_digits());
    }

    #[case("ex1.txt" => 2)]
    #[case("input.txt" => 1108)]
    fn test_part_2(file_name: &str) -> i64 {
        let mut rows = get_input(file_name);
        let res = part2(&mut rows);
        dbg!(res);
        res
    }
}
