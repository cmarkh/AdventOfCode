#![allow(dead_code)]
use color_eyre::Result;

fn read_file(file: &str) -> Result<String> {
    Ok(std::fs::read_to_string(file)?)
}

fn part1(input: &str) -> isize {
    let mut product = 0;

    let muls = input.split("mul(").collect::<Vec<&str>>();
    for mul in muls {
        let mul = match mul.split_once(")") {
            Some((mul, _)) => mul,
            None => continue,
        };
        let (a, b) = match mul.split_once(",") {
            Some((a, b)) => (a, b),
            None => continue,
        };
        let a: isize = match a.parse() {
            Ok(a) => a,
            Err(_) => continue,
        };
        let b: isize = match b.parse() {
            Ok(b) => b,
            Err(_) => continue,
        };
        product += a * b;
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

    #[case("ex1.txt", 161)]
    #[case("input.txt", 161289189)]
    fn test_part1(file: &str, expected: isize) {
        let input = read_file(file).unwrap();
        let result = part1(&input);
        assert_eq!(result, expected);
    }
}
