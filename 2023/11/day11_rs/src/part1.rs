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
}
