#![allow(dead_code)]

struct Grid(Vec<Vec<char>>);

#[derive(Debug)]
struct Numbers(Vec<Number>);

#[derive(Debug)]
struct Number {
    row: usize,
    start: usize,
    end: usize,
    value: u32,
}

impl Grid {
    fn from_str(string: &str) -> Self {
        Grid(string.lines().map(|line| line.chars().collect()).collect())
    }

    fn get_numbers(&self) -> Numbers {
        let mut nums = Numbers(Vec::new());

        for (r, line) in self.0.iter().enumerate() {
            let mut it = line.iter().enumerate();
            while let Some((i, c)) = it.next() {
                if !c.is_ascii_digit() {
                    continue;
                }
                let start = i;
                let mut end = i;
                for (_, c) in it.by_ref() {
                    if !c.is_ascii_digit() {
                        break;
                    }
                    end += 1;
                }
                let value: &u32 = &line[start..=end]
                    .iter()
                    .fold(0, |acc, c| acc * 10 + c.to_digit(10).unwrap());
                nums.0.push(Number {
                    row: r,
                    start,
                    end,
                    value: *value,
                })
            }
        }

        nums
    }

    fn num_adjencent_to_sym(&self, num: &Number) -> bool {
        for r in (num.row.saturating_sub(1))..=(num.row + 1) {
            if let Some(row) = &self.0.get(r) {
                for i in (num.start.saturating_sub(1))..=(num.end + 1) {
                    if row.get(i).is_some_and(|&c| c != '.' && !c.is_ascii_digit()) {
                        return true;
                    }
                }
            }
        }
        false
    }

    fn part1(&self) -> u32 {
        self.get_numbers()
            .0
            .iter()
            .fold(0, |acc, num| match self.num_adjencent_to_sym(num) {
                true => acc + num.value,
                false => acc,
            })
    }
}

#[cfg(test)]
mod test {
    use super::*;
    use test_case::case;

    #[case("ex1.txt" => 4361)]
    #[case("input.txt" => 557705)]
    fn test_part1(input_name: &str) -> u32 {
        let input = std::fs::read_to_string(format!("src/{}", input_name)).unwrap();
        let grid = Grid::from_str(&input);
        grid.part1()
    }

    #[test]
    #[ignore]
    fn test_grid() {
        let input = std::fs::read_to_string("src/ex1.txt").unwrap();
        let grid = Grid::from_str(&input);
        dbg!(grid.0[0][0]);
    }

    #[test]
    #[ignore]
    fn test_get_nums() {
        let input = std::fs::read_to_string("src/ex1.txt").unwrap();
        let grid = Grid::from_str(&input);
        let nums = grid.get_numbers();
        dbg!(nums);
    }

    #[test]
    #[ignore]
    fn test_adj() {
        let input = std::fs::read_to_string("src/ex1.txt").unwrap();
        let grid = Grid::from_str(&input);
        let nums = grid.get_numbers();
        for num in nums.0 {
            if grid.num_adjencent_to_sym(&num) {
                println!("valid: {:?}", num);
            }
        }
    }
}
