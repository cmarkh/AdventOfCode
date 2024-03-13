#![allow(dead_code)]

use std::{collections::HashMap, fmt};

#[derive(Debug)]
pub struct Grid {
    map: HashMap<(usize, usize), Tile>,
    start: (usize, usize),
    size: (usize, usize),
    pipe: Vec<(usize, usize)>,
}

#[derive(Debug)]
struct Tile {
    coordinate: (usize, usize),
    connections: Vec<(usize, usize)>,
    symbol: char,
    symbol_enum: Symbol,
    is_loop: bool,
    is_wall: bool,
}

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

#[derive(Debug)]
enum Symbol {
    Vertical,
    Horizontal,
    NorthEast,
    NorthWest,
    SouthEast,
    SouthWest,
    Ground,
    Start,
}

impl Symbol {
    fn new(symbol: char) -> Self {
        match symbol {
            '|' => Self::Vertical,
            '-' => Self::Horizontal,
            'L' => Self::NorthEast,
            'J' => Self::NorthWest,
            '7' => Self::SouthWest,
            'F' => Self::SouthEast,
            '.' => Self::Ground,
            'S' => Self::Start,
            _ => unreachable!(),
        }
    }
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
                        } else if tile.is_wall {
                            write!(f, "w")?
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
    let mut grid = Grid { map: HashMap::new(), start: (0, 0), size: (0, 0), pipe: Vec::new() };

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
                    symbol_enum: Symbol::new(*col),
                    is_loop: false,
                    is_wall: false,
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
            self.pipe.push(start.coordinate);
            *start.connections.first().unwrap()
        };
        let mut former = self.start;

        while position != self.start {
            let pos = self.map.get_mut(&position).unwrap();
            pos.is_loop = true;
            self.pipe.push(pos.coordinate);

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

    fn mark_wall(&mut self) {
        let mut mark_wall = |right: bool| {
            for i in 0..self.pipe.len() {
                let current = self.map.get(&self.pipe[i]).unwrap();
                let former = if i == 0 { self.pipe[self.pipe.len() - 1] } else { self.pipe[i - 1] };

                let wall = match current.symbol_enum {
                    Symbol::Vertical => match (former.0 > current.coordinate.0, right) {
                        (true, true) => (current.coordinate.0, current.coordinate.1 + 1),
                        (true, false) => (current.coordinate.0, current.coordinate.1 - 1),
                        (false, true) => (current.coordinate.0, current.coordinate.1 - 1),
                        (false, false) => (current.coordinate.0, current.coordinate.1 + 1),
                    },
                    Symbol::Horizontal => match (former.1 > current.coordinate.1, right) {
                        (true, true) => (current.coordinate.0 - 1, current.coordinate.1),
                        (true, false) => (current.coordinate.0 + 1, current.coordinate.1),
                        (false, true) => (current.coordinate.0 - 1, current.coordinate.1),
                        (false, false) => (current.coordinate.0 + 1, current.coordinate.1),
                    },
                    Symbol::NorthEast => match (former.0 < current.coordinate.0, right) {
                        (true, true) => (current.coordinate.0, current.coordinate.1 + 1),
                        (true, false) => (current.coordinate.0, current.coordinate.1),
                        (false, true) => (current.coordinate.0, current.coordinate.1),
                        (false, false) => (current.coordinate.0, current.coordinate.1),
                    },
                    Symbol::NorthWest => match (former.1 > current.coordinate.1, right) {
                        (true, true) => (current.coordinate.0, current.coordinate.1),
                        (true, false) => (current.coordinate.0, current.coordinate.1),
                        (false, true) => (current.coordinate.0, current.coordinate.1),
                        (false, false) => (current.coordinate.0, current.coordinate.1),
                    },
                    Symbol::SouthWest => match (former.1 > current.coordinate.1, right) {
                        (true, true) => (current.coordinate.0, current.coordinate.1),
                        (true, false) => (current.coordinate.0, current.coordinate.1),
                        (false, true) => (current.coordinate.0, current.coordinate.1),
                        (false, false) => (current.coordinate.0, current.coordinate.1),
                    },
                    Symbol::SouthEast => match (former.1 > current.coordinate.1, right) {
                        (true, true) => (current.coordinate.0, current.coordinate.1),
                        (true, false) => (current.coordinate.0, current.coordinate.1),
                        (false, true) => (current.coordinate.0, current.coordinate.1),
                        (false, false) => (current.coordinate.0, current.coordinate.1),
                    },
                    Symbol::Ground => unreachable!(),
                    Symbol::Start => unreachable!(),
                };

                self.map.get_mut(&wall).unwrap().is_wall = true;
            }
        };
    }
}

pub fn part_2(grid: &mut Grid) -> u64 {
    let pipe = grid.mark_loop();

    let mut nests = 0;

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
        grid.mark_wall();
        println!("{}", grid);
    }

    #[ignore]
    #[case("ex2.txt")]
    #[case("ex3.txt")]
    #[case("input.txt")]
    fn test_loop(file_name: &str) {
        let mut grid = get_input(file_name);
        grid.mark_loop();
        dbg!(grid.pipe);
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
