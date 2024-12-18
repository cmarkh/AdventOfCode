#![allow(dead_code)]

use color_eyre::{eyre::OptionExt, Result};

#[cfg(windows)]
const NEWLINE: &str = "\r\n";
#[cfg(not(windows))]
const NEWLINE: &str = "\n";

type Moves = Vec<(i8, i8)>;

#[derive(Debug, Clone)]
struct Grid(Vec<Vec<char>>);

impl Grid {
    fn get(&self, pos: Coordinate) -> char {
        self.0[pos.0][pos.1]
    }

    fn set(&mut self, pos: Coordinate, c: char) {
        self.0[pos.0][pos.1] = c;
    }

    fn print(&self, position: Coordinate) {
        for (r, row) in self.0.iter().enumerate() {
            for (c, cell) in row.iter().enumerate() {
                if r == position.0 && c == position.1 {
                    print!("@");
                } else {
                    print!("{}", cell)
                }
            }
            println!();
        }
        println!();
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
struct Coordinate(usize, usize);

impl Coordinate {
    fn set(&mut self, r: usize, c: usize) {
        self.0 = r;
        self.1 = c;
    }

    fn increment(&mut self, mv: (i8, i8)) {
        self.0 = (self.0 as isize + mv.0 as isize) as usize;
        self.1 = (self.1 as isize + mv.1 as isize) as usize;
    }

    fn decrement(&mut self, mv: (i8, i8)) {
        self.0 = (self.0 as isize - mv.0 as isize) as usize;
        self.1 = (self.1 as isize - mv.1 as isize) as usize;
    }
}

fn read_file(file: &str) -> Result<(Grid, Moves)> {
    let input = std::fs::read_to_string(file)?;
    let (input_grid, input_moves) = input
        .split_once(&format!("{NEWLINE}{NEWLINE}"))
        .ok_or_eyre("Invalid input")?;

    let grid = input_grid.lines().map(|l| l.chars().collect()).collect();

    let input_moves = input_moves.replace(NEWLINE, "");
    let moves = input_moves
        .chars()
        .map(|c| match c {
            '<' => (0, -1),
            '>' => (0, 1),
            '^' => (-1, 0),
            'v' => (1, 0),
            _ => unreachable!(),
        })
        .collect();

    Ok((Grid(grid), moves))
}

fn sum_boxes(grid: &Grid) -> usize {
    let mut sum = 0;

    for (r, row) in grid.0.iter().enumerate() {
        for (c, cell) in row.iter().enumerate() {
            if *cell == 'O' {
                sum += r * 100 + c;
            }
        }
    }

    sum
}

fn part1(mut grid: Grid, moves: Moves) -> usize {
    let mut pos = Coordinate(0, 0);
    for (r, row) in grid.0.iter_mut().enumerate() {
        for (c, cell) in row.iter_mut().enumerate() {
            if *cell == '@' {
                pos.set(r, c);
                *cell = '.';
            }
        }
    }

    'mv: for mv in moves {
        // check if move will be valid
        let mut pos2 = pos;
        loop {
            pos2.increment(mv);
            match grid.get(pos2) {
                '#' => continue 'mv,
                'O' => continue,
                '.' => break,
                _ => unreachable!(),
            }
        }

        // execute the move
        while pos2 != pos {
            let pos3 = pos2;
            pos2.decrement(mv);
            grid.set(pos3, grid.get(pos2));
        }
        pos.increment(mv);

        // grid.print(pos);
    }

    sum_boxes(&grid)
}

#[cfg(test)]
mod tests {
    use test_case::case;

    use super::*;

    #[case("ex1.txt")]
    fn test_read_file(file: &str) {
        let input = read_file(file).unwrap();
        println!("{:?}", input);
    }

    #[case("ex1.txt", 10092)]
    #[case("ex2.txt", 2028)]
    #[case("input.txt", 1294459)]
    fn test_part1(file: &str, expected: usize) {
        let (grid, moves) = read_file(file).unwrap();
        let sum = part1(grid, moves);
        assert_eq!(sum, expected);
    }
}
