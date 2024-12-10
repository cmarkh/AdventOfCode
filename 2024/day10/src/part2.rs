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

    let mut queue = vec![vec![start]];
    let mut history = HashSet::new();

    while let Some(path) = queue.pop() {
        // println!("{:?}", path);
        if !history.insert(path.clone()) {
            continue;
        }

        let &(r, c) = path.last().unwrap();

        let spot = grid.0[r][c];
        if spot == 9 {
            summits.insert(path);
            continue;
        }

        if r > 0 && grid.0[r - 1][c] == spot + 1 {
            let mut path2 = path.clone();
            path2.push((r - 1, c));
            queue.push(path2);
        }
        if r < grid.0.len() - 1 && grid.0[r + 1][c] == spot + 1 {
            let mut path2 = path.clone();
            path2.push((r + 1, c));
            queue.push(path2);
        }
        if c > 0 && grid.0[r][c - 1] == spot + 1 {
            let mut path2 = path.clone();
            path2.push((r, c - 1));
            queue.push(path2);
        }
        if c < grid.0[0].len() - 1 && grid.0[r][c + 1] == spot + 1 {
            let mut path2 = path.clone();
            path2.push((r, c + 1));
            queue.push(path2);
        }
    }

    summits.len()
}

fn part2(grid: &Grid) -> usize {
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

    #[case("ex2.txt", 81)]
    #[case("input.txt", 1062)]
    fn test_part2(file: &str, expected: usize) {
        let grid = read_file(file).unwrap();
        let summits = part2(&grid);
        assert_eq!(summits, expected);
    }
}
