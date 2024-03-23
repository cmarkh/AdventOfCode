#![allow(dead_code)]

use std::collections::HashSet;

#[derive(Debug)]
struct Coordinates {
    coords: Vec<Coordinate>,
    rows: u64,
    cols: u64,
    emtpy_rows: HashSet<u64>,
    empty_cols: HashSet<u64>,
}

type Coordinate = (u64, u64);

fn parse_input(input: &str) -> Coordinates {
    let mut coords = Vec::new();

    for (r, line) in input.lines().enumerate() {
        for (c, ch) in line.chars().enumerate() {
            if ch == '#' {
                coords.push((r as u64, c as u64));
            }
        }
    }

    Coordinates::new(coords)
}

impl Coordinates {
    fn new(coords: Vec<Coordinate>) -> Self {
        let rows = coords.iter().max_by_key(|coord| coord.0).unwrap().0;
        let cols = coords.iter().max_by_key(|coord| coord.1).unwrap().1;
        let emtpy_rows = Self::empty_rows(&coords, rows);
        let empty_cols = Self::empty_cols(&coords, cols);

        Coordinates { coords, rows, cols, emtpy_rows, empty_cols }
    }

    fn empty_rows(coords: &[Coordinate], row_count: u64) -> HashSet<u64> {
        let mut empties = HashSet::new();

        let mut not_empties = HashSet::new();
        for coord in coords {
            not_empties.insert(coord.0);
        }

        for r in 0..row_count {
            if !not_empties.contains(&r) {
                empties.insert(r);
            }
        }

        empties
    }

    fn empty_cols(coords: &[Coordinate], col_count: u64) -> HashSet<u64> {
        let mut empties = HashSet::new();

        let mut not_empties = HashSet::new();
        for coord in coords {
            not_empties.insert(coord.1);
        }

        for c in 0..col_count {
            if !not_empties.contains(&c) {
                empties.insert(c);
            }
        }

        empties
    }

    fn pairs(&self) -> Vec<((u64, u64), (u64, u64))> {
        let mut pairs = Vec::new();

        for i in 0..self.coords.len() {
            for l in (i + 1)..self.coords.len() {
                pairs.push((self.coords[i], self.coords[l]))
            }
        }

        pairs
    }

    fn shortest_path(&self, a: (u64, u64), b: (u64, u64), empty_size: u64) -> u64 {
        let (a_r, a_c) = a;
        let (b_r, b_c) = b;

        let max_r = a_r.max(b_r);
        let min_r = a_r.min(b_r);

        let max_c = a_c.max(b_c);
        let min_c = a_c.min(b_c);

        let mut dist_r = max_r - min_r;
        for r in (min_r + 1)..max_r {
            if self.emtpy_rows.contains(&r) {
                dist_r += empty_size - 1;
            }
        }

        let mut dist_c = max_c - min_c;
        for c in (min_c + 1)..max_c {
            if self.empty_cols.contains(&c) {
                dist_c += empty_size - 1;
            }
        }

        dist_r + dist_c
    }
}

fn part_2(coordinates: Coordinates, empty_size: u64) -> u64 {
    let pairs = coordinates.pairs();

    let mut sum = 0;
    for pair in &pairs {
        sum += coordinates.shortest_path(pair.0, pair.1, empty_size);
    }

    sum
}

#[cfg(test)]
mod test {

    use test_case::case;

    use super::*;

    pub fn get_input(file_name: &str) -> Coordinates {
        let input = std::fs::read_to_string(format!("src/{}", file_name)).unwrap();
        parse_input(&input)
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn test_input(file: &str) {
        let galaxies = get_input(file);
        println!("{:?}", galaxies);
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn print_galaxies(file: &str) {
        let coordinates = get_input(file);
        for coord in &coordinates.coords {
            println!("{:?}", coord);
        }
        dbg!(&coordinates.coords.len());
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn print_pairs(file: &str) {
        let coordinates = get_input(file);
        let pairs = coordinates.pairs();
        for pair in &pairs {
            println!("{:?}", pair);
        }
        dbg!(pairs.len());
    }

    #[case("ex1.txt", 2 => 374)]
    #[case("ex1.txt", 10 => 1030)]
    #[case("ex1.txt", 100 => 8410)]
    #[case("input.txt", 1_000_000 => 553083047914)]
    fn test_part_2(file: &str, empty_size: u64) -> u64 {
        let coordinates = get_input(file);
        part_2(coordinates, empty_size)
    }
}
