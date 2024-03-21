#![allow(dead_code)]

use std::{collections::HashSet, fmt};

struct Galaxies(HashSet<(usize, usize)>);

type Galaxy = (usize, usize);

fn parse_input(input: &str) -> Galaxies {
    let mut galaxies = HashSet::new();

    for (r, line) in input.lines().enumerate() {
        for (c, ch) in line.chars().enumerate() {
            if ch == '#' {
                galaxies.insert((r, c));
            }
        }
    }

    Galaxies(galaxies)
}

impl fmt::Display for Galaxies {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let (r, c) = self.max();

        for r in 0..=r {
            for c in 0..=c {
                if self.0.contains(&(r, c)) {
                    write!(f, "#")?;
                } else {
                    write!(f, ".")?;
                }
            }
            writeln!(f)?;
        }

        Ok(())
    }
}

impl Galaxies {
    fn max(&self) -> (usize, usize) {
        let (mut r, mut c) = (0, 0);

        for galaxy in &self.0 {
            if galaxy.0 > r {
                r = galaxy.0
            }
            if galaxy.1 > c {
                c = galaxy.1
            }
        }

        (r, c)
    }
}

#[cfg(test)]
mod test {
    use test_case::case;

    use super::*;

    pub fn get_input(file_name: &str) -> Galaxies {
        let input = std::fs::read_to_string(format!("src/{}", file_name)).unwrap();
        parse_input(&input)
    }

    #[ignore]
    #[case("ex1.txt")]
    #[case("input.txt")]
    fn test_input(file: &str) {
        let galaxies = get_input(file);
        println!("{}", galaxies);
    }
}
