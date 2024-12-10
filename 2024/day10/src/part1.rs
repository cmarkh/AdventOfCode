#![allow(dead_code)]

use std::{collections::HashSet, fmt::Display};

use color_eyre::{eyre::OptionExt, Result};

struct Grid(Vec<Vec<u8>>);

impl Display for Grid {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for row in &self.0 {
            for cell in row {
                write!(f, "{}", cell)?;
            }
            writeln!(f)?;
        }
        Ok(())
    }
}

fn trailheads(grid: &Grid) -> Vec<(usize, usize)> {
    let mut trailheads: Vec<(usize, usize)> = Vec::new();
    for (r, row) in grid.0.iter().enumerate() {
        for (c, cell) in row.iter().enumerate() {
            if *cell == 0 {
                trailheads.push((r, c));
            }
        }
    }
    trailheads
}

fn follow_trail(grid: &Grid, start: (usize, usize)) -> usize {
    let mut summits = HashSet::new();

    let mut queue = vec![start];
    let mut history = HashSet::new();

    while let Some((r, c)) = queue.pop() {
        if !history.insert((r, c)) {
            continue;
        }

        let spot = grid.0[r][c];
        if spot == 9 {
            summits.insert((r, c));
            continue;
        }

        if r > 0 && grid.0[r - 1][c] == spot + 1 {
            queue.push((r - 1, c));
        }
        if r < grid.0.len() - 1 && grid.0[r + 1][c] == spot + 1 {
            queue.push((r + 1, c));
        }
        if c > 0 && grid.0[r][c - 1] == spot + 1 {
            queue.push((r, c - 1));
        }
        if c < grid.0[0].len() - 1 && grid.0[r][c + 1] == spot + 1 {
            queue.push((r, c + 1));
        }
    }

    summits.len()
}

fn part1(grid: &Grid) -> usize {
    let mut sum = 0;

    let trailheads = trailheads(grid);
    for trailhead in trailheads {
        sum += follow_trail(grid, trailhead);
    }

    sum
}

fn read_file(file: &str) -> Result<Grid> {
    let mut grid = Vec::new();

    let input = std::fs::read_to_string(file)?;
    for line in input.lines() {
        let mut row = Vec::new();
        for char in line.chars() {
            let n: u8 = char.to_digit(10).ok_or_eyre("Invalid digit")? as u8;
            row.push(n);
        }
        grid.push(row);
    }

    Ok(Grid(grid))
}

#[cfg(test)]
mod tests {
    use test_case::case;

    use super::*;

    #[case("ex1.txt")]
    fn test_read_file(file: &str) {
        let grid = read_file(file).unwrap();
        println!("{grid}");
    }

    #[case("ex1.txt", 1)]
    #[case("ex2.txt", 36)]
    #[case("input.txt", 557)]
    fn test_part1(file: &str, expected: usize) {
        let grid = read_file(file).unwrap();
        let summits = part1(&grid);
        assert_eq!(summits, expected);
    }
}
