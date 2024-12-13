#![allow(dead_code)]

use std::collections::{HashMap, HashSet};

use color_eyre::Result;

type Grid = Vec<Vec<char>>;

fn part2(grid: &Grid) -> usize {
    let mut price = 0;

    let mut history: HashSet<(usize, usize)> = HashSet::new();
    #[allow(clippy::type_complexity)]
    let mut groups: HashMap<(char, (usize, usize)), Vec<(usize, usize)>> = HashMap::new();

    for (r, row) in grid.iter().enumerate() {
        for (c, cell) in row.iter().enumerate() {
            let mut perimeter = 0;
            let mut queue = vec![(r, c)];
            while let Some(pos) = queue.pop() {
                if !history.insert(pos) {
                    continue;
                }

                groups.entry((*cell, (r, c))).or_default().push(pos);

                // perimeter
                {
                    // upper left
                    if ((pos.0 == 0 || grid[pos.0 - 1][pos.1] != *cell)
                        && (pos.1 == 0 || grid[pos.0][pos.1 - 1] != *cell))
                        || (pos.0 != 0
                            && pos.1 != 0
                            && grid[pos.0 - 1][pos.1] == *cell
                            && grid[pos.0][pos.1 - 1] == *cell
                            && grid[pos.0 - 1][pos.1 - 1] != *cell)
                    {
                        perimeter += 1;
                    }

                    // upper right
                    if ((pos.0 == 0 || grid[pos.0 - 1][pos.1] != *cell)
                        && (pos.1 == row.len() - 1 || grid[pos.0][pos.1 + 1] != *cell))
                        || (pos.0 != 0
                            && pos.1 != row.len() - 1
                            && grid[pos.0 - 1][pos.1] == *cell
                            && grid[pos.0][pos.1 + 1] == *cell
                            && grid[pos.0 - 1][pos.1 + 1] != *cell)
                    {
                        perimeter += 1;
                    }

                    // lower left
                    if ((pos.0 == grid.len() - 1 || grid[pos.0 + 1][pos.1] != *cell)
                        && (pos.1 == 0 || grid[pos.0][pos.1 - 1] != *cell))
                        || (pos.0 != grid.len() - 1
                            && pos.1 != 0
                            && grid[pos.0 + 1][pos.1] == *cell
                            && grid[pos.0][pos.1 - 1] == *cell
                            && grid[pos.0 + 1][pos.1 - 1] != *cell)
                    {
                        perimeter += 1;
                    }

                    // lower right
                    if ((pos.0 == grid.len() - 1 || grid[pos.0 + 1][pos.1] != *cell)
                        && (pos.1 == row.len() - 1 || grid[pos.0][pos.1 + 1] != *cell))
                        || (pos.0 != grid.len() - 1
                            && pos.1 != row.len() - 1
                            && grid[pos.0 + 1][pos.1] == *cell
                            && grid[pos.0][pos.1 + 1] == *cell
                            && grid[pos.0 + 1][pos.1 + 1] != *cell)
                    {
                        perimeter += 1;
                    }
                }

                // area
                {
                    // up
                    if pos.0 > 0 && grid[pos.0 - 1][pos.1] == *cell {
                        queue.push((pos.0 - 1, pos.1));
                    }

                    // down
                    if pos.0 < grid.len() - 1 && grid[pos.0 + 1][pos.1] == *cell {
                        queue.push((pos.0 + 1, pos.1));
                    }

                    // left
                    if pos.1 > 0 && grid[pos.0][pos.1 - 1] == *cell {
                        queue.push((pos.0, pos.1 - 1));
                    }

                    // right
                    if pos.1 < row.len() - 1 && grid[pos.0][pos.1 + 1] == *cell {
                        queue.push((pos.0, pos.1 + 1));
                    }
                }
            }

            if let Some(group) = groups.get(&(*cell, (r, c))) {
                price += group.len() * perimeter;
                // println!("{} {:?} {}*{}", grid[r][c], group, group.len(), perimeter);
            }

            // println!("break\n");
        }
    }

    price
}

fn read_file(file: &str) -> Result<Grid> {
    let input = std::fs::read_to_string(file)?;
    let grid = input.lines().map(|line| line.chars().collect()).collect();
    Ok(grid)
}

#[cfg(test)]
mod tests {
    use test_case::case;

    use super::*;

    #[case("ex1.txt")]
    fn test_read_file(file: &str) {
        let stones = read_file(file).unwrap();
        println!("{:?}", stones);
    }

    #[case("ex1.txt", 80)]
    #[case("ex3.txt", 1206)]
    #[case("ex4.txt", 236)]
    #[case("ex5.txt", 368)]
    #[case("input.txt", 839780)]
    fn test_part2(file: &str, expected: usize) {
        let grid = read_file(file).unwrap();
        let price = part2(&grid);
        assert_eq!(price, expected);
    }
}
