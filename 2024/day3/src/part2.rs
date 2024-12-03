#![allow(dead_code)]
use color_eyre::Result;

fn read_file(file: &str) -> Result<String> {
    Ok(std::fs::read_to_string(file)?)
}

fn part2(mut input: &str) -> isize {
    let mut product = 0;

    let mut do_ = true;

    'outer: while !input.is_empty() {
        if input.len() >= 4 && do_ && (&input[0..4] == "mul(") {
            input = &input[4..];

            let mut i = 0;
            for c in input.chars() {
                i += 1;
                if c.is_numeric() {
                    continue;
                }
                if c == ',' {
                    break;
                }
                continue 'outer;
            }
            let a: isize = input[0..i - 1].parse().unwrap();
            input = &input[i..];

            let mut i = 0;
            for c in input.chars() {
                i += 1;
                if c.is_numeric() {
                    continue;
                }
                if c == ')' {
                    break;
                }
                continue 'outer;
            }
            let b: isize = input[0..i - 1].parse().unwrap();
            input = &input[i..];

            product += a * b;
            continue;
        }
        if input.len() >= 5 && &input[0..5] == "don't" {
            do_ = false;
            input = &input[5..];
            continue;
        }
        if input.len() >= 2 && &input[0..2] == "do" {
            do_ = true;
            input = &input[2..];
            continue;
        }
        input = &input[1..]
    }

    product
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

    #[case("ex2.txt", 48)]
    #[case("input.txt", 83595109)]
    fn test_part2(file: &str, expected: isize) {
        let input = read_file(file).unwrap();
        let result = part2(&input);
        assert_eq!(result, expected);
    }
}
