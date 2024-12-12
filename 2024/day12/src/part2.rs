#![allow(dead_code)]

use std::collections::{HashMap, HashSet};

use color_eyre::Result;

type Grid = Vec<Vec<char>>;

fn part2(grid: &Grid) -> u64 {
    let mut price = 0;

    let mut history: HashSet<(usize, usize)> = HashSet::new();
    let mut groups: HashMap<(char, (usize, usize)), u64> = HashMap::new();
    let mut perimeters: HashMap<(char, (usize, usize)), u64> = HashMap::new();

    for (r, row) in grid.iter().enumerate() {
        for (c, cell) in row.iter().enumerate() {
            let mut queue = vec![(r, c)];
            while let Some(pos) = queue.pop() {
                if !history.insert(pos) {
                    continue;
                }

                // println!("{:?}: {:?}", cell, pos);

                *groups.entry((*cell, (r, c))).or_insert(0) += 1;

                // up
                if pos.0 > 0 {
                    if grid[pos.0 - 1][pos.1] == *cell {
                        queue.push((pos.0 - 1, pos.1));
                    } else {
                        *perimeters.entry((*cell, (r, c))).or_insert(0) += 1;
                    }
                } else {
                    *perimeters.entry((*cell, (r, c))).or_insert(0) += 1;
                }

                // down
                if pos.0 < grid.len() - 1 {
                    if grid[pos.0 + 1][pos.1] == *cell {
                        queue.push((pos.0 + 1, pos.1));
                    } else {
                        *perimeters.entry((*cell, (r, c))).or_insert(0) += 1;
                    }
                } else {
                    *perimeters.entry((*cell, (r, c))).or_insert(0) += 1;
                }

                // left
                if pos.1 > 0 {
                    if grid[pos.0][pos.1 - 1] == *cell {
                        queue.push((pos.0, pos.1 - 1));
                    } else {
                        *perimeters.entry((*cell, (r, c))).or_insert(0) += 1;
                    }
                } else {
                    *perimeters.entry((*cell, (r, c))).or_insert(0) += 1;
                }

                // right
                if pos.1 < row.len() - 1 {
                    if grid[pos.0][pos.1 + 1] == *cell {
                        queue.push((pos.0, pos.1 + 1));
                    } else {
                        *perimeters.entry((*cell, (r, c))).or_insert(0) += 1;
                    }
                } else {
                    *perimeters.entry((*cell, (r, c))).or_insert(0) += 1;
                }
            }
            // println!("break\n");
        }
    }

    for group in groups.keys() {
        price += groups[group] * perimeters[group];
        // println!("{:?}: {}*{}", group, groups[group], perimeters[group]);
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

    #[case("ex1.txt", 140)]
    #[case("ex2.txt", 772)]
    #[case("ex3.txt", 1930)]
    #[case("input.txt", 1359028)]
    fn test_part2(file: &str, expected: u64) {
        let grid = read_file(file).unwrap();
        let price = part2(&grid);
        assert_eq!(price, expected);
    }
}
