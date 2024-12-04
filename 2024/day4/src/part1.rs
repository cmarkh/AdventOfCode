#![allow(dead_code)]
use color_eyre::Result;

struct Grid(Vec<Vec<char>>);

impl std::fmt::Display for Grid {
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

impl Grid {
    fn get(&self, x: isize, y: isize) -> Option<&char> {
        match self.0.get(x as usize) {
            Some(row) => row.get(y as usize),
            None => None,
        }
    }
}

fn read_file(file: &str) -> Result<Grid> {
    let input = std::fs::read_to_string(file)?;
    let grid = Grid(input.lines().map(|line| line.chars().collect()).collect());
    Ok(grid)
}

fn part1(grid: &Grid) -> isize {
    let mut finds = 0;

    for (r, row) in grid.0.iter().enumerate() {
        for (c, cell) in row.iter().enumerate() {
            if cell == &'X' {
                finds += test_x(grid, r as isize, c as isize)
            }
        }
    }

    finds
}

fn test_x(grid: &Grid, x: isize, y: isize) -> isize {
    let mut finds = 0;

    if right(grid, x, y) {
        finds += 1;
    }
    if left(grid, x, y) {
        finds += 1;
    }
    if up(grid, x, y) {
        finds += 1;
    }
    if down(grid, x, y) {
        finds += 1;
    }
    if up_right(grid, x, y) {
        finds += 1;
    }
    if down_right(grid, x, y) {
        finds += 1;
    }
    if up_left(grid, x, y) {
        finds += 1;
    }
    if down_left(grid, x, y) {
        finds += 1;
    }

    finds
}

fn right(grid: &Grid, x: isize, y: isize) -> bool {
    if !grid.get(x + 1, y).map_or(false, |cell| cell == &'M') {
        return false;
    }
    if !grid.get(x + 2, y).map_or(false, |cell| cell == &'A') {
        return false;
    }
    if !grid.get(x + 3, y).map_or(false, |cell| cell == &'S') {
        return false;
    }
    true
}

fn left(grid: &Grid, x: isize, y: isize) -> bool {
    if !grid.get(x - 1, y).map_or(false, |cell| cell == &'M') {
        return false;
    }
    if !grid.get(x - 2, y).map_or(false, |cell| cell == &'A') {
        return false;
    }
    if !grid.get(x - 3, y).map_or(false, |cell| cell == &'S') {
        return false;
    }
    true
}

fn up(grid: &Grid, x: isize, y: isize) -> bool {
    if !grid.get(x, y - 1).map_or(false, |cell| cell == &'M') {
        return false;
    }
    if !grid.get(x, y - 2).map_or(false, |cell| cell == &'A') {
        return false;
    }
    if !grid.get(x, y - 3).map_or(false, |cell| cell == &'S') {
        return false;
    }
    true
}

fn down(grid: &Grid, x: isize, y: isize) -> bool {
    if !grid.get(x, y + 1).map_or(false, |cell| cell == &'M') {
        return false;
    }
    if !grid.get(x, y + 2).map_or(false, |cell| cell == &'A') {
        return false;
    }
    if !grid.get(x, y + 3).map_or(false, |cell| cell == &'S') {
        return false;
    }
    true
}

fn up_right(grid: &Grid, x: isize, y: isize) -> bool {
    if !grid.get(x + 1, y - 1).map_or(false, |cell| cell == &'M') {
        return false;
    }
    if !grid.get(x + 2, y - 2).map_or(false, |cell| cell == &'A') {
        return false;
    }
    if !grid.get(x + 3, y - 3).map_or(false, |cell| cell == &'S') {
        return false;
    }
    true
}

fn down_right(grid: &Grid, x: isize, y: isize) -> bool {
    if !grid.get(x + 1, y + 1).map_or(false, |cell| cell == &'M') {
        return false;
    }
    if !grid.get(x + 2, y + 2).map_or(false, |cell| cell == &'A') {
        return false;
    }
    if !grid.get(x + 3, y + 3).map_or(false, |cell| cell == &'S') {
        return false;
    }
    true
}

fn up_left(grid: &Grid, x: isize, y: isize) -> bool {
    if !grid.get(x - 1, y - 1).map_or(false, |cell| cell == &'M') {
        return false;
    }
    if !grid.get(x - 2, y - 2).map_or(false, |cell| cell == &'A') {
        return false;
    }
    if !grid.get(x - 3, y - 3).map_or(false, |cell| cell == &'S') {
        return false;
    }
    true
}

fn down_left(grid: &Grid, x: isize, y: isize) -> bool {
    if !grid.get(x - 1, y + 1).map_or(false, |cell| cell == &'M') {
        return false;
    }
    if !grid.get(x - 2, y + 2).map_or(false, |cell| cell == &'A') {
        return false;
    }
    if !grid.get(x - 3, y + 3).map_or(false, |cell| cell == &'S') {
        return false;
    }
    true
}

#[cfg(test)]
mod tests {
    use test_case::case;

    use super::*;

    #[case("ex1.txt")]
    fn test_read_file(file: &str) {
        let input = read_file(file).unwrap();
        println!("{}", input);
    }

    #[case("ex1.txt", 18)]
    #[case("input.txt", 2618)]
    fn test_part1(file: &str, expected: isize) {
        let input = read_file(file).unwrap();
        let result = part1(&input);
        assert_eq!(result, expected);
    }
}
