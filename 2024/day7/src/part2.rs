#![allow(dead_code)]

use color_eyre::{eyre::OptionExt, Result};

type Equations = Vec<Equation>;

#[derive(Clone)]
struct Equation {
    answer: i64,
    digits: Vec<i64>,
}

impl std::fmt::Display for Equation {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}: {:?}", self.answer, self.digits)
    }
}

fn read_file(file: &str) -> Result<Equations> {
    let mut equations = Equations::new();

    let input = std::fs::read_to_string(file)?;

    for line in input.lines() {
        let (answer, digits) = line.split_once(':').ok_or_eyre("Invalid input")?;
        let answer: i64 = answer.trim().parse()?;
        let digits: Vec<i64> = digits
            .split(' ')
            .filter(|&s| !s.is_empty())
            .map(|s| s.parse())
            .collect::<Result<_, _>>()?;
        equations.push(Equation { answer, digits });
    }

    Ok(equations)
}

fn test_equation(equation: &Equation) -> bool {
    let mut queue: Vec<(Equation, Vec<char>)> = vec![(equation.clone(), Vec::new())];

    while let Some(item) = queue.pop() {
        // println!("{} {:?}", item.0, item.1);

        let mut result = item.0.digits[0];
        for (i, digit) in item.0.digits.iter().enumerate() {
            if i == 0 {
                continue;
            }
            if i > item.1.len() {
                break;
            }
            result = operation(result, *digit, item.1[i - 1]);
        }
        if result > equation.answer {
            continue;
        }
        if item.1.len() == item.0.digits.len() - 1 {
            if result == equation.answer {
                return true;
            }
            continue;
        }

        let mut item2 = item.clone();
        item2.1.push('+');
        queue.push(item2);

        let mut item2 = item.clone();
        item2.1.push('*');
        queue.push(item2);

        let mut item2 = item.clone();
        item2.1.push('|');
        queue.push(item2);
    }

    false
}

fn operation(a: i64, b: i64, operator: char) -> i64 {
    match operator {
        '+' => a + b,
        '*' => a * b,
        '|' => format!("{}{}", a, b).parse().unwrap(),
        _ => unreachable!("Invalid operator"),
    }
}

fn part2(equations: &Equations) -> i64 {
    let mut sum = 0;

    for equation in equations {
        if test_equation(equation) {
            sum += equation.answer;
        }
    }

    sum
}

#[cfg(test)]
mod tests {
    use test_case::case;

    use super::*;

    #[case("ex1.txt")]
    fn test_read_file(file: &str) {
        let equations = read_file(file).unwrap();
        for equation in equations {
            println!("{equation}");
        }
    }

    #[case("ex1.txt", 11387)]
    #[case("input.txt", 37598910447546)]
    fn test_part2(file: &str, expected: i64) {
        let equations = read_file(file).unwrap();
        let result = part2(&equations);
        assert_eq!(result, expected);
    }
}
