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
    size: (usize, usize),
}

#[derive(Debug)]
struct Tile {
    coordinate: (usize, usize),
    connections: Vec<(usize, usize)>,
    symbol: char,
    is_loop: bool,
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
                    Some(tile) => {
                        if tile.is_loop {
                            write!(f, "*")?
                        } else {
                            write!(f, "{}", tile.symbol)?
                        }
                    }
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
        size: (0, 0),
    };

    let mut char_grid: Vec<Vec<char>> = Vec::with_capacity(input.lines().count());
    for line in input.lines() {
        let row: Vec<char> = line.chars().collect();
        char_grid.push(row);
    }

    let valid_conn = |r: usize, c: usize| -> Option<char> {
        if let Some(row) = char_grid.get(r)
            && let Some(col) = row.get(c)
            && *col != '.'
        {
            return Some(*col);
        }
        None
    };

    for (r, row) in char_grid.iter().enumerate() {
        for (c, col) in row.iter().enumerate() {
            let coordinate = (r, c);
            let mut connections = Vec::new();
            match col {
                '|' => connections = vec![(r.saturating_sub(1), c), (r + 1, c)],
                '-' => connections = vec![(r, c.saturating_sub(1)), (r, c + 1)],
                'L' => connections = vec![(r.saturating_sub(1), c), (r, c + 1)],
                'J' => connections = vec![(r.saturating_sub(1), c), (r, c.saturating_sub(1))],
                '7' => connections = vec![(r, c.saturating_sub(1)), (r + 1, c)],
                'F' => connections = vec![(r, c + 1), (r + 1, c)],
                '.' => (),
                'S' => {
                    grid.start = (r, c);

                    let mut xy = (r.saturating_sub(1), c);
                    if let Some(ch) = valid_conn(xy.0, xy.1)
                        && (ch == '|' || ch == '7' || ch == 'F')
                    {
                        connections.push(xy);
                    }
                    xy = (r, c.saturating_sub(1));
                    if let Some(ch) = valid_conn(xy.0, xy.1)
                        && (ch == '-' || ch == 'F' || ch == 'L')
                    {
                        connections.push(xy);
                    }
                    xy = (r + 1, c);
                    if let Some(ch) = valid_conn(xy.0, xy.1)
                        && (ch == '|' || ch == 'L' || ch == 'J')
                    {
                        connections.push(xy);
                    }
                    xy = (r, c + 1);
                    if let Some(ch) = valid_conn(xy.0, xy.1)
                        && (ch == '-' || ch == 'J' || ch == '7')
                    {
                        connections.push(xy);
                    }
                }
                _ => unreachable!(),
            }
            grid.map.insert(
                coordinate,
                Tile {
                    coordinate,
                    connections,
                    symbol: *col,
                    is_loop: false,
                },
            );
            grid.size = (r + 1, c + 1);
        }
    }

    grid
}

impl Grid {
    fn mark_loop(&mut self) {
        let mut position = {
            let start = self.map.get_mut(&self.start).unwrap();
            start.is_loop = true;
            *start.connections.first().unwrap()
        };
        let mut former = self.start;

        while position != self.start {
            let pos = self.map.get_mut(&position).unwrap();
            pos.is_loop = true;

            let mut conns = pos.connections.iter();
            let maybe = conns.next().unwrap();
            if maybe != &former {
                former = position;
                position = *maybe;
            } else {
                former = position;
                position = *conns.next().unwrap();
            }
        }
    }

    fn is_tile_inside_loop(&self, tile: &Tile) -> bool {
        if tile.is_loop {
            return false;
        }

        dbg!(tile.coordinate);

        let mut loop_count = 0;
        for r in (0..tile.coordinate.0).rev() {
            let neighbor = self.map.get(&(r, tile.coordinate.1)).unwrap();
            if neighbor.is_loop {
                loop_count += 1;
            }
        }
        if loop_count % 2 == 0 {
            return false;
        }
        dbg!((1, tile.coordinate));

        let mut loop_count = 0;
        for r in tile.coordinate.0..self.size.0 {
            let neighbor = self.map.get(&(r, tile.coordinate.1)).unwrap();
            if neighbor.is_loop {
                loop_count += 1;
            }
        }
        if loop_count % 2 == 0 {
            return false;
        }
        dbg!((2, tile.coordinate));

        let mut loop_count = 0;
        for c in (0..tile.coordinate.1).rev() {
            let neighbor = self.map.get(&(tile.coordinate.0, c)).unwrap();
            if neighbor.is_loop {
                loop_count += 1;
            }
        }
        if loop_count % 2 == 0 {
            return false;
        }
        dbg!((3, tile.coordinate));

        let mut loop_count = 0;
        for c in tile.coordinate.1..self.size.1 {
            let neighbor = self.map.get(&(tile.coordinate.0, c)).unwrap();
            if neighbor.is_loop {
                loop_count += 1;
            }
        }
        if loop_count % 2 == 0 {
            return false;
        }
        dbg!((4, tile.coordinate));

        true
    }
}

pub fn part_2(grid: &mut Grid) -> u64 {
    grid.mark_loop();

    let mut nests = 0;

    for r in 0..grid.size.0 {
        for c in 0..grid.size.1 {
            let tile = grid.map.get(&(r, c)).unwrap();
            if grid.is_tile_inside_loop(tile) {
                dbg!((r, c));
                nests += 1;
            }
        }
    }

    nests
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
    #[case("ex2.txt")]
    #[case("input.txt")]
    fn test_input(file_name: &str) {
        let grid = get_input(file_name);
        dbg!(&grid);
        dbg!(&grid.start);
        dbg!(&grid.size);
    }

    #[ignore]
    #[case("ex2.txt")]
    #[case("ex3.txt")]
    #[case("input.txt")]
    fn test_print_grid(file_name: &str) {
        let mut grid = get_input(file_name);
        grid.mark_loop();
        println!("{}", grid);
    }

    #[case("ex2.txt" => 4)]
    #[case("ex3.txt" => 10)]
    #[case("input.txt" => 6923)]
    fn test_part_2(file_name: &str) -> u64 {
        let mut grid = get_input(file_name);
        let nest_spaces = part_2(&mut grid);
        dbg!(nest_spaces)
    }
}
