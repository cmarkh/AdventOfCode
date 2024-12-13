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
            let mut queue = vec![(r, c)];
            while let Some(pos) = queue.pop() {
                if !history.insert(pos) {
                    continue;
                }

                groups.entry((*cell, (r, c))).or_default().push(pos);

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
            // println!("break\n");
        }
    }

    for group in &mut groups.values_mut() {
        let mut perimeter = 0;

        let rs = group.iter().map(|(r, _)| r).collect::<HashSet<_>>();
        for r in rs {
            dbg!(r);
            let mut row = group
                .iter()
                .filter(|(rr, _)| r == rr)
                .map(|(_, c)| c)
                .collect::<Vec<_>>();
            row.sort();
            for i in 1..row.len() {
                if *row[i] != (row[i - 1] + 1) {
                    perimeter += 1;
                }
            }
            perimeter += 2;
        }
        let cs = group.iter().map(|(_, c)| c).collect::<HashSet<_>>();
        for c in cs {
            let mut col = group
                .iter()
                .filter(|(_, cc)| c == cc)
                .map(|(r, _)| r)
                .collect::<Vec<_>>();
            col.sort();
            for i in 1..col.len() {
                if col[i] != col[i - 1] {
                    perimeter += 1;
                }
            }
            perimeter += 2;
        }

        dbg!(perimeter);

        price += perimeter * group.len();
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
    #[case("input.txt", 1359028)]
    fn test_part2(file: &str, expected: usize) {
        let grid = read_file(file).unwrap();
        let price = part2(&grid);
        assert_eq!(price, expected);
    }
}
