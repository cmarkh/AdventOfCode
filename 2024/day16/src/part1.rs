#![allow(dead_code)]

use color_eyre::{eyre::OptionExt, Result};

fn read_file(file: &str) -> Result<()> {
    let input = std::fs::read_to_string(file)?;
    todo!()
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

    // #[case("ex1.txt", 10092)]
    // #[case("input.txt", 1294459)]
    // fn test_part1(file: &str, expected: usize) {
    //     let (grid, moves) = read_file(file).unwrap();
    //     let sum = part1(grid, moves);
    //     assert_eq!(sum, expected);
    // }
}
