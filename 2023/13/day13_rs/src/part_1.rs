#![allow(dead_code)]

use std::cmp::Ordering;

#[derive(Debug)]
struct Grid(Vec<Vec<bool>>);

impl Grid {
    pub fn new() -> Self {
        Grid(Vec::new())
    }
}

impl std::fmt::Display for Grid {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for row in &self.0 {
            for col in row {
                match col {
                    true => write!(f, "#")?,
                    false => write!(f, ".")?,
                }
            }
            writeln!(f)?;
        }
        Ok(())
    }
}

fn parse_input(input: &str) -> Vec<Grid> {
    let mut grids = Vec::new();

    let mut grid = Grid::new();
    for line in input.lines() {
        if line.is_empty() {
            grids.push(grid);
            grid = Grid::new();
            continue;
        }

        let mut row = Vec::new();
        for char in line.chars() {
            match char {
                '.' => row.push(false),
                '#' => row.push(true),
                _ => unreachable!(),
            }
        }
        grid.0.push(row);
    }
    grids.push(grid);

    grids
}

impl Grid {
    fn find_vertical(&self) -> Option<usize> {
        for center in 1..self.0[0].len() {
            'center: {
                for row in &self.0 {
                    let offset = 'offset: {
                        for offset in 0..center + 1 {
                            if row[center - offset] {
                                break 'offset offset + 1;
                            }
                        }
                        break 'center;
                    };
                    if center + offset >= row.len() || !row[center + offset] {
                        break 'center;
                    }
                    for offset in 1..offset {
                        // all cells between the center and offset should be false
                        if row[center + offset] {
                            break 'center;
                        }
                    }
                }
                // if made it here, the center is good
                return Some(center + 1);
            }
        }

        None
    }

    fn find_horizontal(&self) -> Option<usize> {
        for center in 1..self.0.len() {
            'center: {
                for col in 0..self.0.len() {
                    let offset = 'offset: {
                        for offset in 0..center + 1 {
                            if self.0[center - offset][col] {
                                break 'offset offset + 1;
                            }
                        }
                        break 'center;
                    };
                    if center + offset >= self.0.len() || !self.0[center + offset][col] {
                        break 'center;
                    }
                    for offset in 1..offset {
                        if self.0[center + offset][col] {
                            break 'center;
                        }
                    }
                }
                return Some(center + 1);
            }
        }

        None
    }
}

fn part_1(grids: Vec<Grid>) -> usize {
    let (mut verticals, mut horizontals) = (0, 0);

    for grid in grids {
        // if let Some(center) = grid.find_vertical() {
        //     verticals += center;
        // } else if let Some(center) = grid.find_horizontal() {
        //     horizontals += center;
        // }
        let vertical = grid.find_vertical();
        let horizontal = grid.find_horizontal();
        match (vertical, horizontal) {
            (None, Some(h)) => horizontals += h,
            (Some(v), None) => verticals += v,
            (Some(v), Some(h)) => match v.cmp(&h) {
                Ordering::Equal => unreachable!(),
                Ordering::Less => verticals += v,
                Ordering::Greater => horizontals += h,
            },
            (None, None) => unreachable!(),
        }
        dbg!(verticals, horizontals);
    }

    // verticals += grids.first().unwrap().find_vertical().unwrap();
    // horizontals += grids.last().unwrap().find_horizontal().unwrap();

    verticals + horizontals * 100
}

#[cfg(test)]
mod test {
    use test_case::case;

    use super::*;

    pub fn get_input(file_name: &str) -> Vec<Grid> {
        let input = std::fs::read_to_string(format!("src/{}", file_name)).unwrap();
        parse_input(&input)
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn print_input(file: &str) {
        let grids = get_input(file);
        for grid in &grids {
            print!("{}", grid);
        }
    }

    #[case("ex1.txt" => Some(5))]
    fn test_vertical(file: &str) -> Option<usize> {
        let grids = get_input(file);
        let grid = grids.first().unwrap();
        let center = grid.find_vertical();
        dbg!(center)
    }

    #[case("ex1.txt" => Some(4))]
    fn test_horizontal(file: &str) -> Option<usize> {
        let grids = get_input(file);
        let grid = grids.last().unwrap();
        let center = grid.find_horizontal();
        dbg!(center)
    }

    #[case("ex1.txt" => 405)]
    fn test_part_1(file: &str) -> usize {
        let grids = get_input(file);
        let res = part_1(grids);
        dbg!(res)
    }
}
