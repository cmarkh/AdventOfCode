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
pub struct Grid {
    map: HashMap<(usize, usize), Tile>,
    start: (usize, usize),
}

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
        for ((y, x), _) in self.map.iter() {
            if *x > max_x {
                max_x = *x;
            }
            if *y > max_y {
                max_y = *y;
            }
        }

        for y in 0..=max_y {
            for x in 0..=max_x {
                match self.map.get(&(y, x)) {
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
    let mut grid = Grid {
        map: HashMap::new(),
        start: (0, 0),
    };

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
            dbg!(r, c);
            match col {
                '|' => connections = [(r.saturating_sub(1), c), (r + 1, c)],
                '-' => connections = [(r, c.saturating_sub(1)), (r, c + 1)],
                'L' => connections = [(r.saturating_sub(1), c), (r, c + 1)],
                'J' => connections = [(r.saturating_sub(1), c), (r, c.saturating_sub(1))],
                '7' => connections = [(r, c.saturating_sub(1)), (r + 1, c)],
                'F' => connections = [(r, c + 1), (r + 1, c)],
                '.' => (),
                'S' => {
                    grid.start = (r, c);

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
                    if valid_conn(xy.0, xy.1) && conn_idx < 2 {
                        connections[conn_idx] = xy;
                        conn_idx += 1;
                    }
                    xy = (r, c + 1);
                    if valid_conn(xy.0, xy.1) && conn_idx < 2 {
                        connections[conn_idx] = xy;
                    }
                }
                _ => unreachable!(),
            }
            grid.map.insert(
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

impl Grid {
    fn walk(&self) -> HashMap<(usize, usize), u64> {
        let mut history = HashMap::<(usize, usize), u64>::new();

        let mut queue: Vec<((usize, usize), u64)> = Vec::new();
        queue.push((self.start, 0)); // (row, col), steps

        while let Some(position) = queue.pop() {
            if let Some(steps) = history.get(&position.0) {
                if position.1 >= *steps {
                    continue;
                }
            }
            history.insert(position.0, position.1);

            let tile = self.map.get(&position.0).unwrap();
            queue.push((tile.connections[0], position.1 + 1));
            queue.push((tile.connections[1], position.1 + 1));
        }

        history
    }
}

pub fn part_1(grid: Grid) -> u64 {
    let history = grid.walk();

    let mut max = 0;
    for (position, steps) in history {
        if steps > max {
            max = steps;
            dbg!(max);
            dbg!(position);
            dbg!(grid.map.get(&position).unwrap());
        }
    }

    max
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
        dbg!(&grid);
        dbg!(&grid.start);
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn test_print_grid(file_name: &str) {
        let grid = get_input(file_name);
        println!("{}", grid);
    }

    #[case("ex1.txt" => 4)]
    #[case("input.txt" => 0)]
    fn test_part_1(file_name: &str) -> u64 {
        let grid = get_input(file_name);
        let max_steps = part_1(grid);
        dbg!(&max_steps);
        max_steps
    }
}
