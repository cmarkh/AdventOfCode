#![allow(dead_code)]

use std::collections::HashMap;

use tap::Pipe;

struct Grid(Vec<Vec<char>>);

#[derive(Debug)]
struct Numbers(Vec<Number>);

#[derive(Debug, Clone, Copy)]
struct Number {
    row: usize,
    start: usize,
    end: usize,
    value: u32,
}

struct Gears(HashMap<Point, Vec<Number>>);

#[derive(PartialEq, PartialOrd, Hash, Eq, Debug, Clone, Copy)]
struct Point {
    row: usize,
    col: usize,
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

    fn get_gears(&self, numbers: &Numbers) -> Gears {
        let mut gears = Gears(HashMap::new());

        for num in &numbers.0 {
            for r in (num.row.saturating_sub(1))..=(num.row + 1) {
                if let Some(row) = &self.0.get(r) {
                    for i in (num.start.saturating_sub(1))..=(num.end + 1) {
                        if row.get(i).is_some_and(|&c| c == '*') {
                            let point = Point { row: r, col: i };
                            let gear = gears.0.entry(point).or_default();
                            gear.push(*num);
                        }
                    }
                }
            }
        }

        let remove: Vec<_> = gears
            .0
            .iter()
            .filter_map(|(gear, nums)| if nums.len() != 2 { Some(*gear) } else { None })
            .collect();
        for gear in remove {
            gears.0.remove(&gear);
        }

        gears
    }

    fn part2(&self) -> u32 {
        self.get_numbers().pipe(|nums| {
            self.get_gears(&nums).0.iter().fold(0, |acc, (_, nums)| {
                acc + nums.iter().fold(1, |acc, num| acc * num.value)
            })
        })
    }
}

#[cfg(test)]
mod test {
    use super::*;
    use test_case::case;

    #[case("ex1.txt" => 467835)]
    #[case("input.txt" => 84266818)]
    fn test_part2(input_name: &str) -> u32 {
        let input = std::fs::read_to_string(format!("src/day3/{}", input_name)).unwrap();
        let grid = Grid::from_str(&input);
        grid.part2()
    }

    #[test]
    #[ignore]
    fn test_grid() {
        let input = std::fs::read_to_string("src/day3/ex1.txt").unwrap();
        let grid = Grid::from_str(&input);
        dbg!(grid.0[0][0]);
    }

    #[test]
    #[ignore]
    fn test_get_nums() {
        let input = std::fs::read_to_string("src/day3/ex1.txt").unwrap();
        let grid = Grid::from_str(&input);
        let nums = grid.get_numbers();
        dbg!(nums);
    }

    #[test]
    #[ignore]
    fn test_gears() {
        let input = std::fs::read_to_string("src/day3/ex1.txt").unwrap();
        let grid = Grid::from_str(&input);
        let nums = grid.get_numbers();
        let gears = grid.get_gears(&nums);
        for gear in &gears.0 {
            dbg!(gear);
        }
    }
}
