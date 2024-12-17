#![allow(dead_code)]

use color_eyre::{eyre::OptionExt, Result};
use regex::Regex;

fn read_file(file: &str) -> Result<Vec<Robot>> {
    let mut robots = Vec::new();

    let input = std::fs::read_to_string(file)?;

    let re = Regex::new(r"p=(\d+),(\d+)\s+v=(-?\d+),(-?\d+)").unwrap();

    for line in input.lines() {
        let line = re.captures(line).ok_or_eyre("Invalid line")?;
        let x = line[1].parse()?;
        let y = line[2].parse()?;
        let v_x = line[3].parse()?;
        let v_y = line[4].parse()?;
        robots.push(Robot { x, y, v_x, v_y });
    }

    Ok(robots)
}

fn print(robots: &[Robot], len_x: i64, len_y: i64) {
    for y in 0..len_y {
        for x in 0..len_x {
            let mut count = 0;
            for robot in robots {
                if robot.x == x && robot.y == y {
                    count += 1;
                }
            }
            if count == 0 {
                print!(".");
            } else {
                print!("{count}");
            }
        }
        println!();
    }
}

#[derive(Debug)]
struct Robot {
    x: i64,
    y: i64,
    v_x: i64,
    v_y: i64,
}

impl Robot {
    fn step(&mut self, len_x: i64, len_y: i64) {
        self.x += self.v_x;
        if self.x >= len_x {
            self.x -= len_x;
        }
        if self.x < 0 {
            self.x += len_x;
        }

        self.y += self.v_y;
        if self.y >= len_y {
            self.y -= len_y;
        }
        if self.y < 0 {
            self.y += len_y;
        }
    }
}

fn part1(mut robots: Vec<Robot>, len_x: i64, len_y: i64) -> i64 {
    for robot in robots.iter_mut() {
        for _ in 0..100 {
            robot.step(len_x, len_y);
        }
    }

    let half_x = len_x / 2;
    let half_y = len_y / 2;

    let mut quad_1 = 0;
    for robot in robots.iter() {
        if robot.x < half_x && robot.y < half_y {
            quad_1 += 1;
        }
    }

    let mut quad_2 = 0;
    for robot in robots.iter() {
        if robot.x > half_x && robot.y < half_y {
            quad_2 += 1;
        }
    }

    let mut quad_3 = 0;
    for robot in robots.iter() {
        if robot.x < half_x && robot.y > half_y {
            quad_3 += 1;
        }
    }

    let mut quad_4 = 0;
    for robot in robots.iter() {
        if robot.x > half_x && robot.y > half_y {
            quad_4 += 1;
        }
    }

    for robot in robots.iter() {
        println!("{:?}", robot);
    }
    print(&robots, len_x, len_y);

    quad_1 * quad_2 * quad_3 * quad_4
}

#[cfg(test)]
mod tests {
    use test_case::case;

    use super::*;

    #[case("ex1.txt")]
    fn test_read_file(file: &str) {
        let input = read_file(file).unwrap();
        println!("{:?}", input);
    }

    #[case("ex1.txt", 11, 7, 12)]
    #[case("input.txt", 101, 103, 231221760)]
    fn test_part1(file: &str, len_x: i64, len_y: i64, expected: i64) {
        let robots = read_file(file).unwrap();
        let safety = part1(robots, len_x, len_y);
        assert_eq!(safety, expected);
    }
}
