#![allow(dead_code)]

use std::fmt;

struct Grid(Vec<Vec<Cell>>);

#[derive(Debug, Clone, Copy)]
struct Cell {
    is_galaxy: bool,
}

fn parse_input(input: &str) -> Grid {
    let mut grid = Vec::with_capacity(input.lines().count());

    for line in input.lines() {
        let mut row = vec![Cell { is_galaxy: false }; line.chars().count()];
        for (c, ch) in line.chars().enumerate() {
            if ch == '#' {
                if let Some(cell) = row.get_mut(c) {
                    cell.is_galaxy = true;
                }
            }
        }
        grid.push(row);
    }

    Grid(grid)
}

impl fmt::Display for Grid {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        for row in &self.0 {
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
    fn double_empties(&mut self) {
        // empty rows
        {
            let mut rows = self.0.first().unwrap().len();
            let mut r = 0;
            let empty_row = vec![Cell { is_galaxy: false }; self.0.len()];
            while r < rows {
                if self.0.get(r).unwrap().iter().any(|row| row.is_galaxy) {
                    r += 1;
                    continue;
                }
                self.0.insert(r + 1, empty_row.clone());
                r += 2;
                rows += 1;
            }
        }

        // empty columns
        {
            let mut cols = self.0.len();
            let rows = self.0.first().unwrap().len();
            let mut c = 0;
            while c < cols {
                let mut empty = true;
                for r in 0..rows {
                    if self.0[r][c].is_galaxy {
                        empty = false;
                        break;
                    }
                }
                if !empty {
                    c += 1;
                    continue;
                }

                for r in 0..rows {
                    self.0.insert(r + 1, vec![Cell { is_galaxy: false }; cols]);
                }
                c += 2;
                cols += 1;
            }
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
