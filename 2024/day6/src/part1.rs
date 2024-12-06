#![allow(dead_code)]

use std::collections::HashSet;

use color_eyre::Result;

struct Grid(Vec<Vec<char>>);

impl Grid {
    fn get(&self, r: usize, c: usize) -> char {
        match self.0.get(r) {
            Some(row) => *row.get(c).unwrap(),
            None => unreachable!(),
        }
    }

    fn find_guard(&self) -> Guard {
        for (r, row) in self.0.iter().enumerate() {
            for (c, col) in row.iter().enumerate() {
                if *col == '^' || *col == 'v' || *col == '<' || *col == '>' {
                    return Guard {
                        row: r,
                        col: c,
                        direction: *col,
                    };
                }
            }
        }
        unreachable!()
    }

    fn print(&self, guard: &Guard) {
        for (r, row) in self.0.iter().enumerate() {
            for (c, cell) in row.iter().enumerate() {
                if r == guard.row && c == guard.col {
                    print!("{}", guard.direction);
                } else {
                    print!("{}", cell);
                }
            }
            println!();
        }
        println!();
    }
}

#[derive(Debug)]
struct Guard {
    row: usize,
    col: usize,
    direction: char,
}

impl Guard {
    fn turn_and_step(&mut self) {
        match self.direction {
            '^' => {
                self.direction = '>';
                self.col += 1;
            }
            '>' => {
                self.direction = 'v';
                self.row += 1;
            }
            'v' => {
                self.direction = '<';
                self.col -= 1;
            }
            '<' => {
                self.direction = '^';
                self.row -= 1;
            }
            _ => unreachable!(),
        }
    }
}

fn read_file(file: &str) -> Result<Grid> {
    let input = std::fs::read_to_string(file)?;
    let grid = input.lines().map(|line| line.chars().collect()).collect();
    Ok(Grid(grid))
}

fn step(grid: &Grid, guard: &mut Guard) {
    match guard.direction {
        '^' => {
            if guard.row == 0 {
                guard.turn_and_step();
            } else if grid.get(guard.row - 1, guard.col) == '.' {
                guard.row -= 1;
            } else {
                guard.turn_and_step();
            }
        }
        '>' => {
            if grid.get(guard.row, guard.col + 1) == '.' {
                guard.col += 1;
            } else {
                guard.turn_and_step();
            }
        }
        'v' => {
            if grid.get(guard.row + 1, guard.col) == '.' {
                guard.row += 1;
            } else {
                guard.turn_and_step();
            }
        }
        '<' => {
            if guard.col == 0 {
                guard.turn_and_step();
            } else if grid.get(guard.row, guard.col - 1) == '.' {
                guard.col -= 1;
            } else {
                guard.turn_and_step();
            }
        }
        _ => unreachable!(),
    }
}

fn done(grid: &Grid, guard: &mut Guard) -> bool {
    match guard.direction {
        '^' => {
            if guard.row == 0 {
                return true;
            }
        }
        '>' => {
            if guard.col == grid.0[0].len() - 1 {
                return true;
            }
        }
        'v' => {
            if guard.row == grid.0.len() - 1 {
                return true;
            }
        }
        '<' => {
            if guard.col == 0 {
                return true;
            }
        }
        _ => unreachable!(),
    }
    false
}

fn part1(grid: &mut Grid) -> u32 {
    let mut visited = HashSet::new();

    let mut guard = grid.find_guard();
    grid.0[guard.row][guard.col] = '.';

    loop {
        // grid.print(&guard);
        // println!("------------------");
        // enter();
        visited.insert((guard.row, guard.col));
        if done(grid, &mut guard) {
            break;
        }
        step(grid, &mut guard);
    }

    visited.len() as u32
}

fn enter() {
    println!("Press Enter to continue...");
    let mut input = String::new();
    std::io::Write::flush(&mut std::io::stdout()).unwrap(); // Ensure the prompt is displayed immediately
    std::io::stdin().read_line(&mut input).unwrap();
    println!();
}

#[cfg(test)]
mod tests {
    use test_case::case;

    use super::*;

    #[case("ex1.txt")]
    fn test_read_file(file: &str) {
        let grid = read_file(file).unwrap();
        let guard = grid.find_guard();
        grid.print(&guard);
    }

    #[case("ex1.txt", 41)]
    #[case("input.txt", 5964)]
    fn test_part1(file: &str, expected: u32) {
        let mut grid = read_file(file).unwrap();
        let result = part1(&mut grid);
        assert_eq!(result, expected);
    }
}
