#![allow(dead_code)]

use std::collections::HashSet;

use color_eyre::{eyre::OptionExt, Result};

#[cfg(windows)]
const NEWLINE: &str = "\r\n";
#[cfg(not(windows))]
const NEWLINE: &str = "\n";

type Moves = Vec<(i8, i8)>;

// (row, left, right)
type Boxes = HashSet<Box>;
type Box = (usize, usize, usize);

type Walls = HashSet<(usize, usize)>;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
struct Coordinate(usize, usize);

impl Coordinate {
    fn set(&mut self, r: usize, c: usize) {
        self.0 = r;
        self.1 = c;
    }

    fn increment(&mut self, mv: (i8, i8)) {
        self.0 = (self.0 as isize + mv.0 as isize) as usize;
        self.1 = (self.1 as isize + mv.1 as isize) as usize;
    }

    fn decrement(&mut self, mv: (i8, i8)) {
        self.0 = (self.0 as isize - mv.0 as isize) as usize;
        self.1 = (self.1 as isize - mv.1 as isize) as usize;
    }
}

fn read_file(file: &str) -> Result<(Moves, Boxes, Walls, Coordinate)> {
    let mut moves = Vec::new();
    let mut boxes = HashSet::new();
    let mut walls = HashSet::new();
    let mut position = Coordinate(0, 0);

    let input = std::fs::read_to_string(file)?;
    let (input_grid, input_moves) = input
        .split_once(&format!("{NEWLINE}{NEWLINE}"))
        .ok_or_eyre("Invalid input")?;

    let mut grid = Vec::new();
    for line in input_grid.lines() {
        let mut row = Vec::new();
        for char in line.chars() {
            match char {
                '.' => row.extend("..".chars()),
                '#' => row.extend("##".chars()),
                'O' => row.extend("[]".chars()),
                '@' => row.extend("@.".chars()),
                _ => unreachable!(),
            }
        }
        grid.push(row);
    }

    for (r, row) in grid.iter().enumerate() {
        for (c, cell) in row.iter().enumerate() {
            match cell {
                '[' => boxes.insert((r, c, c + 1)),
                ']' => false,
                '#' => walls.insert((r, c)),
                '@' => {
                    position.set(r, c);
                    false
                }
                '.' => false,
                _ => unreachable!(),
            };
        }
    }

    let input_moves = input_moves.replace(NEWLINE, "");
    for char in input_moves.chars() {
        match char {
            '<' => moves.push((0, -1)),
            '>' => moves.push((0, 1)),
            '^' => moves.push((-1, 0)),
            'v' => moves.push((1, 0)),
            _ => unreachable!(),
        }
    }

    Ok((moves, boxes, walls, position))
}

fn print(boxes: &Boxes, walls: &Walls, pos: Coordinate) {
    let mut max_r = 0;
    let mut max_c = 0;
    for (row, _, r) in boxes.iter() {
        max_r = max_r.max(*row);
        max_c = max_c.max(*r);
    }
    for (r, c) in walls.iter() {
        max_r = max_r.max(*r);
        max_c = max_c.max(*c);
    }

    for r in 0..=max_r {
        for c in 0..=max_c {
            if (pos.0, pos.1) == (r, c) {
                print!("@");
            } else if walls.contains(&(r, c)) {
                print!("#");
            } else if boxes.contains(&(r, c, c + 1)) {
                print!("[");
            } else if boxes.contains(&(r, c - 1, c)) {
                print!("]");
            } else {
                print!(".");
            }
        }
        println!();
    }
    println!();
}

fn make_move(mv: (i8, i8), boxes: &mut Boxes, walls: &Walls, pos: &mut Coordinate) {
    let mut pos2 = *pos;
    pos2.increment(mv);

    if walls.contains(&(pos2.0, pos2.1)) {
        return;
    }

    // box on the right
    if boxes.contains(&(pos2.0, pos2.1, pos2.1 + 1))
        && !move_box(mv, boxes, walls, (pos2.0, pos2.1, pos2.1 + 1))
    {
        return;
    }

    // box on the left
    if boxes.contains(&(pos2.0, pos2.1 - 1, pos2.1))
        && !move_box(mv, boxes, walls, (pos2.0, pos2.1 - 1, pos2.1))
    {
        return;
    }

    *pos = pos2;
}

fn move_box(mv: (i8, i8), boxes: &mut Boxes, walls: &Walls, bx: Box) -> bool {
    let mut boxes2 = boxes.clone();

    let mut bx2 = bx;
    bx2.0 = (bx2.0 as isize + mv.0 as isize) as usize;
    bx2.1 = (bx2.1 as isize + mv.1 as isize) as usize;
    bx2.2 = (bx2.2 as isize + mv.1 as isize) as usize;

    if walls.contains(&(bx2.0, bx2.1)) {
        return false;
    }
    if walls.contains(&(bx2.0, bx2.2)) {
        return false;
    }

    if mv.1 == 1 || mv.1 == -1 {
        let bx3 = (
            bx2.0,
            (bx2.1 as isize + mv.1 as isize) as usize,
            (bx2.2 as isize + mv.1 as isize) as usize,
        );
        if boxes.contains(&bx3) && !move_box(mv, &mut boxes2, walls, bx3) {
            return false;
        }
    } else {
        // above or below
        if boxes.contains(&bx2) && !move_box(mv, &mut boxes2, walls, bx2) {
            return false;
        }

        // left
        let bx3 = (bx2.0, bx2.1 - 1, bx2.2 - 1);
        if boxes.contains(&bx3) && !move_box(mv, &mut boxes2, walls, bx3) {
            return false;
        }

        // right
        let bx3 = (bx2.0, bx2.1 + 1, bx2.2 + 1);
        if boxes.contains(&bx3) && !move_box(mv, &mut boxes2, walls, bx3) {
            return false;
        }
    }

    boxes2.remove(&bx);
    boxes2.insert(bx2);

    *boxes = boxes2;

    true
}

fn part1(moves: Moves, mut boxes: Boxes, walls: Walls, mut pos: Coordinate) -> usize {
    // print(&boxes, &walls, pos);

    for mv in moves {
        make_move(mv, &mut boxes, &walls, &mut pos);
    }

    let mut sum = 0;
    for bx in boxes {
        sum += bx.0 * 100 + bx.1;
    }

    sum
}

#[cfg(test)]
mod tests {
    use test_case::case;

    use super::*;

    #[case("ex1.txt")]
    #[case("ex3.txt")]
    fn test_read_file(file: &str) {
        let (moves, boxes, walls, pos) = read_file(file).unwrap();
        println!("{:?}", moves);
        println!("{:?}", boxes);
        print(&boxes, &walls, pos);
    }

    // #[case("ex3.txt", 0)]
    #[case("ex1.txt", 9021)]
    #[case("input.txt", 1319212)]
    fn test_part1(file: &str, expected: usize) {
        let (moves, boxes, walls, pos) = read_file(file).unwrap();
        let sum = part1(moves, boxes, walls, pos);
        assert_eq!(sum, expected);
    }
}
