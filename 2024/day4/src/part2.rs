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

fn part2(grid: &Grid) -> isize {
    let mut finds = 0;

    for (r, row) in grid.0.iter().enumerate() {
        for (c, cell) in row.iter().enumerate() {
            if cell == &'A' && test_x(grid, r as isize, c as isize) {
                finds += 1;
            }
        }
    }

    finds
}

fn test_x(grid: &Grid, x: isize, y: isize) -> bool {
    let ul = match grid.get(x - 1, y - 1) {
        Some(cell) => cell,
        None => return false,
    };
    let ur = match grid.get(x + 1, y - 1) {
        Some(cell) => cell,
        None => return false,
    };
    let dl = match grid.get(x - 1, y + 1) {
        Some(cell) => cell,
        None => return false,
    };
    let dr = match grid.get(x + 1, y + 1) {
        Some(cell) => cell,
        None => return false,
    };

    let left = format!("{}{}", ul, dr);
    let right = format!("{}{}", ur, dl);

    (left == "MS" || left == "SM") && (right == "MS" || right == "SM")
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

    #[case("ex1.txt", 9)]
    #[case("input.txt", 2011)]
    fn test_part2(file: &str, expected: isize) {
        let input = read_file(file).unwrap();
        let result = part2(&input);
        assert_eq!(result, expected);
    }
}
