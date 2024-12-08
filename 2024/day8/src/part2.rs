#![allow(dead_code)]

use std::collections::{HashMap, HashSet};

use color_eyre::Result;

#[derive(Debug)]
struct Grid {
    max_row: i32,
    max_col: i32,
}

#[derive(Debug)]
struct Antennas(HashMap<char, Vec<Point>>);

#[derive(Debug, Clone, Copy, PartialEq, Eq, Hash)]
struct Point {
    row: i32,
    col: i32,
}

#[derive(Debug)]
struct Distance {
    row: i32,
    col: i32,
}

fn read_file(file: &str) -> Result<(Grid, Antennas)> {
    let mut antennas = HashMap::new();

    let input = std::fs::read_to_string(file)?;

    let grid = Grid {
        max_row: input.lines().count() as i32 - 1,
        max_col: input.lines().next().unwrap().len() as i32 - 1,
    };

    for (r, row) in input.lines().enumerate() {
        for (c, ch) in row.chars().enumerate() {
            if ch == '.' {
                continue;
            }
            antennas.entry(ch).or_insert_with(Vec::new).push(Point {
                row: r as i32,
                col: c as i32,
            });
        }
    }

    Ok((grid, Antennas(antennas)))
}

fn find_antinodes(grid: &Grid, antennas: &[Point]) -> HashSet<Point> {
    let mut antinodes = HashSet::new();

    for (i, a) in antennas.iter().enumerate() {
        for (j, b) in antennas.iter().enumerate() {
            if i == j {
                continue;
            }
            let x = a.col - b.col;
            let y = a.row - b.row;

            let mut anti = Point {
                row: a.row + y,
                col: a.col + x,
            };
            while anti.row >= 0
                && anti.row <= grid.max_row
                && anti.col >= 0
                && anti.col <= grid.max_col
            {
                antinodes.insert(anti);
                anti.row += y;
                anti.col += x;
            }

            let mut anti = Point {
                row: a.row - y,
                col: a.col - x,
            };
            while anti.row >= 0
                && anti.row <= grid.max_row
                && anti.col >= 0
                && anti.col <= grid.max_col
            {
                antinodes.insert(anti);
                anti.row -= y;
                anti.col -= x;
            }

            let mut anti = Point {
                row: b.row + y,
                col: b.col + x,
            };
            while anti.row >= 0
                && anti.row <= grid.max_row
                && anti.col >= 0
                && anti.col <= grid.max_col
            {
                antinodes.insert(anti);
                anti.row += y;
                anti.col += x;
            }

            let mut anti = Point {
                row: b.row - y,
                col: b.col - x,
            };
            while anti.row >= 0
                && anti.row <= grid.max_row
                && anti.col >= 0
                && anti.col <= grid.max_col
            {
                antinodes.insert(anti);
                anti.row -= y;
                anti.col -= x;
            }
        }
    }

    antinodes
}

fn distance(a: &Point, b: &Point) -> i32 {
    (a.row - b.row).abs() + (a.col - b.col).abs()
}

fn part2(grid: &Grid, antennas: &Antennas) -> i32 {
    let mut locations = HashSet::new();

    for antennas in antennas.0.values() {
        let antinodes = find_antinodes(grid, antennas);
        // for antinode in &antinodes {
        //     println!("{:?}", antinode);
        // }
        locations.extend(antinodes);
    }

    locations.len() as i32
}

#[cfg(test)]
mod tests {
    use test_case::case;

    use super::*;

    #[case("ex1.txt")]
    fn test_read_file(file: &str) {
        let (grid, antennas) = read_file(file).unwrap();
        println!("{:?}", grid);
        for antenna in antennas.0 {
            println!("{:?}", antenna);
        }
    }

    #[case("ex1.txt", 34)]
    #[case("ex4.txt", 9)]
    #[case("input.txt", 1277)]
    fn test_part2(file: &str, expected: i32) {
        let (grid, antennas) = read_file(file).unwrap();
        let result = part2(&grid, &antennas);
        assert_eq!(result, expected);
    }
}
