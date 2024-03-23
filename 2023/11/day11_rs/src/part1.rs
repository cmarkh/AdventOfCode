#![allow(dead_code)]

use std::fmt;

type Grid = crate::grid::Grid<Cell>;

#[derive(Debug, Clone, Copy)]
struct Cell {
    is_galaxy: bool,
}

fn parse_input(input: &str) -> Grid {
    let mut grid =
        Grid::new(input.lines().count(), input.lines().next().unwrap().chars().count(), Cell { is_galaxy: false });

    for (r, line) in input.lines().enumerate() {
        for (c, ch) in line.chars().enumerate() {
            if ch == '#' {
                if let Some(cell) = grid.get_mut(r, c) {
                    cell.is_galaxy = true;
                }
            }
        }
    }

    grid
}

impl fmt::Display for Grid {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        for row in &self.data {
            for cell in row {
                if cell.is_galaxy {
                    write!(f, "#")?;
                } else {
                    write!(f, ".")?;
                }
            }
            writeln!(f)?;
        }

        Ok(())
    }
}

impl Grid {
    fn row_empty(&self, r: usize) -> bool {
        if let Some(row) = self.data.get(r) {
            for c in row {
                if c.is_galaxy {
                    return false;
                }
            }
        }
        true
    }

    fn column_empty(&self, c: usize) -> bool {
        for r in 0..self.rows {
            if let Some(cell) = self.get(r, c) {
                if cell.is_galaxy {
                    return false;
                }
            }
        }
        true
    }

    fn double_empties(&mut self) {
        let mut r = 0;
        while r < self.rows {
            if !self.row_empty(r) {
                r += 1;
                continue;
            }
            self.insert_row(r + 1, vec![Cell { is_galaxy: false }; self.columns]);
            r += 2;
        }

        let mut c = 0;
        while c < self.columns {
            if !self.column_empty(c) {
                c += 1;
                continue;
            }
            self.insert_column(c + 1, vec![Cell { is_galaxy: false }; self.rows]);
            c += 2;
        }
    }

    fn collect_galaxy_coords(&self) -> Vec<(usize, usize)> {
        let mut galaxies = Vec::new();
        for (r, row) in self.data.iter().enumerate() {
            for (c, cell) in row.iter().enumerate() {
                if cell.is_galaxy {
                    galaxies.push((r, c));
                }
            }
        }
        galaxies
    }
}

fn shortest_path(a: (usize, usize), b: (usize, usize)) -> usize {
    let (a_r, a_c) = a;
    let (b_r, b_c) = b;

    let dist_r = if a_r > b_r { a_r - b_r } else { b_r - a_r };
    let dist_c = if a_c > b_c { a_c - b_c } else { b_c - a_c };

    dist_r + dist_c
}

fn pairs(coords: Vec<(usize, usize)>) -> Vec<((usize, usize), (usize, usize))> {
    let mut pairs = Vec::new();

    for i in 0..coords.len() {
        for l in (i + 1)..coords.len() {
            pairs.push((coords[i], coords[l]))
        }
    }

    pairs
}

fn part_1(grid: Grid) -> usize {
    let galaxies = grid.collect_galaxy_coords();
    let pairs = pairs(galaxies);

    let mut sum = 0;
    for pair in &pairs {
        sum += shortest_path(pair.0, pair.1);
    }

    sum
}

#[cfg(test)]
mod test {

    use test_case::case;

    use super::*;

    pub fn get_input(file_name: &str) -> Grid {
        let input = std::fs::read_to_string(format!("src/{}", file_name)).unwrap();
        parse_input(&input)
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn test_input(file: &str) {
        let galaxies = get_input(file);
        println!("{}", galaxies);
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn test_double_empties(file: &str) {
        let mut galaxies = get_input(file);
        galaxies.double_empties();
        println!("{}", galaxies);
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn print_galaxies(file: &str) {
        let mut grid = get_input(file);
        grid.double_empties();
        let galaxies = grid.collect_galaxy_coords();
        for coord in &galaxies {
            println!("{:?}", coord);
        }
        dbg!(&galaxies.len());
    }

    #[case((0,4), (10,9) => 15)]
    #[case((2,0), (7,12) => 17)]
    #[case((7,12), (2,0) => 17)]
    fn test_shortest_path(a: (usize, usize), b: (usize, usize)) -> usize {
        let dist = shortest_path(a, b);
        dbg!(dist)
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn print_pairs(file: &str) {
        let mut grid = get_input(file);
        grid.double_empties();
        let galaxies = grid.collect_galaxy_coords();
        let pairs = pairs(galaxies);
        for pair in &pairs {
            println!("{:?}", pair);
        }
        dbg!(pairs.len());
    }

    #[case("ex1.txt" => 374)]
    #[case("input.txt" => 10154062)]
    fn test_part_1(file: &str) -> usize {
        let mut grid = get_input(file);
        grid.double_empties();
        let res = part_1(grid);
        dbg!(res)
    }
}
