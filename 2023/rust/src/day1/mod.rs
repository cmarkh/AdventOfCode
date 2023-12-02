#![allow(dead_code)]

pub fn part1(input: &str) -> i32 {
    let mut sum = 0;
    for line in input.lines() {
        let mut line_num = "".to_string();
        for char in line.chars() {
            if char.is_numeric() {
                line_num.push(char);
                break;
            }
        }
        for char in line.chars().rev() {
            if char.is_numeric() {
                line_num.push(char);
                break;
            }
        }
        let line_num = line_num.parse::<i32>().unwrap();
        sum += line_num;
    }
    sum
}

lazy_static::lazy_static! {
    static ref NUMBERS: std::collections::HashMap<String, i32> = {
        let mut set = std::collections::HashMap::new();
        set.insert("one".to_string(), 1);
        set.insert("two".to_string(), 2);
        set.insert("three".to_string(), 3);
        set.insert("four".to_string(), 4);
        set.insert("five".to_string(), 5);
        set.insert("six".to_string(), 6);
        set.insert("seven".to_string(), 7);
        set.insert("eight".to_string(), 8);
        set.insert("nine".to_string(), 9);
        set
    };
}

pub fn part2(input: &str) -> i32 {
    let mut sum = 0;

    for line in input.lines() {
        let nums = get_nums(line);
        let string = format!("{}{}", nums[0].1, nums[nums.len() - 1].1);
        sum += string.parse::<i32>().unwrap();
    }

    sum
}

fn get_nums(line: &str) -> Vec<(usize, char)> {
    let mut nums = Vec::new();

    for (i, ch) in line.char_indices() {
        if ch.is_numeric() {
            nums.push((i, ch))
        }
    }

    for (strng, n) in NUMBERS.iter() {
        match line.find(strng) {
            Some(i) => nums.push((i, char::from_digit(*n as u32, 10).unwrap())),
            None => continue,
        }
        match line.rfind(strng) {
            Some(i) => nums.push((i, char::from_digit(*n as u32, 10).unwrap())),
            None => continue,
        }
    }

    nums.sort_by_key(|tup| tup.0);
    nums
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        let case = std::fs::read_to_string("src/day1/input1.txt").unwrap();
        let ans = part1(&case);
        println!("{}", ans);
        assert_eq!(part1(&case), 56049);
    }

    #[test]
    fn test_part2() {
        let case = std::fs::read_to_string("src/day1/input1.txt").unwrap();
        let ans = part2(&case);
        println!("{}", ans);
        assert_eq!(ans, 54530);
    }

    #[test]
    fn ex1() {
        let case = std::fs::read_to_string("src/day1/ex1.txt").unwrap();
        assert_eq!(part1(&case), 142);
    }

    #[test]
    fn ex1_2() {
        let case = std::fs::read_to_string("src/day1/ex1.txt").unwrap();
        assert_eq!(part2(&case), 142);
    }

    #[test]
    fn ex2() {
        let case = std::fs::read_to_string("src/day1/ex2.txt").unwrap();
        let ans = part2(&case);
        assert_eq!(ans, 281);
    }

    #[test]
    #[ignore]
    fn test_get_nums() {
        let case = std::fs::read_to_string("src/day1/ex2.txt").unwrap();
        let line = case.lines().last().unwrap();
        let nums = get_nums(line);
        for num in nums {
            println!("{:?}", num);
        }
    }
}
