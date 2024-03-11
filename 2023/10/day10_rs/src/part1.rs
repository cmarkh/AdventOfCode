#![allow(dead_code)]

/*
| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
 */

use std::{collections::HashMap, fmt};

#[derive(Debug)]
struct Grid(HashMap<(usize, usize), Tile>);

#[derive(Debug)]
struct Tile {
    coordinate: (usize, usize),
    connections: [(usize, usize); 2],
    symbol: String,
}

impl fmt::Display for Grid {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let mut max_x = 0;
        let mut max_y = 0;
        for ((y, x), _) in self.0.iter() {
            if *x > max_x {
                max_x = *x;
            }
            if *y > max_y {
                max_y = *y;
            }
        }

        for y in 0..=max_y {
            for x in 0..=max_x {
                match self.0.get(&(y, x)) {
                    Some(tile) => write!(f, "{}", tile.symbol)?,
                    None => write!(f, " ")?,
                }
            }
            writeln!(f)?;
        }

        Ok(())
    }
}

fn parse_input(input: &str) -> Grid {
    let mut grid = Grid(HashMap::new());

    let mut char_grid: Vec<Vec<char>> = Vec::with_capacity(input.lines().count());
    for line in input.lines() {
        let row: Vec<char> = line.chars().collect();
        char_grid.push(row);
    }

    let valid_conn = |r: usize, c: usize| -> bool {
        if let Some(row) = char_grid.get(r)
            && let Some(col) = row.get(c)
            && *col != '.'
        {
            return true;
        }
        false
    };

    for (r, row) in char_grid.iter().enumerate() {
        for (c, col) in row.iter().enumerate() {
            let coordinate = (r, c);
            let mut connections: [(usize, usize); 2] = [(0, 0); 2];
            match col {
                '|' => connections = [(r - 1, c), (r + 1, c)],
                '-' => connections = [(r, c - 1), (r, c + 1)],
                'L' => connections = [(r - 1, c), (r, c + 1)],
                'J' => connections = [(r - 1, c), (r, c - 1)],
                '7' => connections = [(r, c - 1), (r + 1, c)],
                'F' => connections = [(r, c + 1), (r - 1, c)],
                '.' => (),
                'S' => {
                    let mut conn_idx = 0;
                    let mut xy = (r - 1, c);
                    if valid_conn(xy.0, xy.1) {
                        connections[conn_idx] = xy;
                        conn_idx += 1;
                    }
                    xy = (r, c - 1);
                    if valid_conn(xy.0, xy.1) {
                        connections[conn_idx] = xy;
                        conn_idx += 1;
                    }
                    xy = (r + 1, c);
                    if valid_conn(xy.0, xy.1) {
                        connections[conn_idx] = xy;
                        conn_idx += 1;
                    }
                    xy = (r, c + 1);
                    if valid_conn(xy.0, xy.1) {
                        connections[conn_idx] = xy;
                    }
                }
                _ => unreachable!(),
            }
            grid.0.insert(
                coordinate,
                Tile {
                    coordinate,
                    connections,
                    symbol: col.to_string(),
                },
            );
        }
    }

    grid
}

#[cfg(test)]
mod test {
    use super::*;

    use test_case::case;

    pub fn get_input(file_name: &str) -> Grid {
        let input = std::fs::read_to_string(format!("src/{}", file_name)).unwrap();
        parse_input(&input)
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn test_input(file_name: &str) {
        let grid = get_input(file_name);
        dbg!(grid);
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn test_print_grid(file_name: &str) {
        let grid = get_input(file_name);
        println!("{}", grid);
    }
}
