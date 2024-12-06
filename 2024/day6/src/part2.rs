#![allow(dead_code)]

use std::collections::HashSet;

use color_eyre::Result;

fn read_file(file: &str) -> Result<Grid> {
    let input = std::fs::read_to_string(file)?;
    let grid: Vec<Vec<char>> = input.lines().map(|line| line.chars().collect()).collect();

    Ok(Grid(grid))
}

#[derive(Debug, Clone)]
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

#[derive(Debug, Clone, PartialEq, Eq, Hash)]
struct Guard {
    row: usize,
    col: usize,
    direction: char,
}

fn next_is_blocked(grid: &Grid, guard: &mut Guard) -> bool {
    match guard.direction {
        '^' => {
            if guard.row == 0 {
                return false;
            }
            grid.get(guard.row - 1, guard.col) != '.'
        }
        '>' => {
            if guard.col == grid.0[0].len() - 1 {
                return false;
            }
            grid.get(guard.row, guard.col + 1) != '.'
        }
        'v' => {
            if guard.row == grid.0.len() - 1 {
                return false;
            }
            grid.get(guard.row + 1, guard.col) != '.'
        }
        '<' => {
            if guard.col == 0 {
                return false;
            }
            grid.get(guard.row, guard.col - 1) != '.'
        }
        _ => unreachable!(),
    }
}

fn turn(guard: &mut Guard) {
    match guard.direction {
        '^' => guard.direction = '>',
        '>' => guard.direction = 'v',
        'v' => guard.direction = '<',
        '<' => guard.direction = '^',
        _ => unreachable!(),
    }
}

// step returns true if off the edge of the map
fn step(grid: &Grid, guard: &mut Guard) -> bool {
    match guard.direction {
        '^' => {
            if guard.row == 0 {
                return true;
            }
            guard.row -= 1;
        }
        '>' => {
            if guard.col == grid.0[0].len() - 1 {
                return true;
            }
            guard.col += 1;
        }
        'v' => {
            if guard.row == grid.0.len() - 1 {
                return true;
            }
            guard.row += 1;
        }
        '<' => {
            if guard.col == 0 {
                return true;
            }
            guard.col -= 1;
        }
        _ => unreachable!(),
    }
    false
}

fn path(grid: &mut Grid, guard: &mut Guard) -> Option<HashSet<Guard>> {
    let mut visited = HashSet::new();
    loop {
        // grid.print(guard);
        // enter();

        if !visited.insert(guard.clone()) {
            return None;
        }
        while next_is_blocked(grid, guard) {
            turn(guard);
        }
        if step(grid, guard) {
            return Some(visited);
        }
    }
}

fn part2(grid: &mut Grid) -> u32 {
    let mut options = 0;

    let og_guard = grid.find_guard();
    grid.0[og_guard.row][og_guard.col] = '.';

    let original = path(grid, &mut og_guard.clone()).unwrap();
    let original = original
        .iter()
        .map(|guard| (guard.row, guard.col))
        .collect::<HashSet<_>>();

    for (row, col) in original {
        if row == og_guard.row && col == og_guard.col {
            continue;
        }
        let mut grid = grid.clone();
        grid.0[row][col] = 'O';
        match path(&mut grid, &mut og_guard.clone()) {
            Some(_) => {}
            None => options += 1,
        }
        dbg!(&options);
    }

    options
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

    #[case("ex1.txt", 6)]
    #[case("input.txt", 1778)]
    fn test_part2(file: &str, expected: u32) {
        let mut grid = read_file(file).unwrap();
        let result = part2(&mut grid);
        assert_eq!(result, expected);
    }
}
