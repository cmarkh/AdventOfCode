#![allow(dead_code)]

use color_eyre::{eyre::OptionExt, Result};

type Instructions = Vec<(u32, u32)>;
type Update = Vec<u32>;
type Updates = Vec<Update>;

fn read_file(file: &str) -> Result<(Instructions, Updates)> {
    let mut instructions = Instructions::new();
    let mut updates = Updates::new();

    let input = std::fs::read_to_string(file)?;
    let input = input.replace("\r\n", "\n");
    let (split_instructions, split_updates) = input.split_once("\n\n").ok_or_eyre("no split 1")?;

    for line in split_instructions.lines() {
        let (l, r) = line.split_once("|").ok_or_eyre("no split 2")?;
        let l = l.trim().parse()?;
        let r = r.trim().parse()?;
        instructions.push((l, r));
    }

    for line in split_updates.lines() {
        let update = line
            .split(',')
            .map(|n| n.trim().parse::<u32>())
            .collect::<Result<_, _>>()?;
        updates.push(update);
    }

    Ok((instructions, updates))
}

fn ordered(instructions: &Instructions, update: &Update) -> bool {
    for (l, r) in instructions {
        for (i, u) in update.iter().enumerate() {
            if u == l {
                for before in update.iter().take(i) {
                    if before == r {
                        return false;
                    }
                }
            }
            if u == r {
                for after in update.iter().skip(i + 1) {
                    if after == l {
                        return false;
                    }
                }
            }
        }
    }
    true
}

fn reorder(instructions: &Instructions, update: &mut Update) {
    for (l, r) in instructions {
        for (i, u) in update.iter().enumerate() {
            if u == l {
                for sub_i in 0..i {
                    if update[sub_i] == *r {
                        update.remove(sub_i);
                        update.insert(i, *r);
                        return reorder(instructions, update);
                    }
                }
            }
            if u == r {
                for sub_i in (i + 1)..update.len() {
                    if update[sub_i] == *l {
                        update.remove(sub_i);
                        update.insert(i, *l);
                        return reorder(instructions, update);
                    }
                }
            }
        }
    }
}

fn part2(instructions: &Instructions, updates: Updates) -> u32 {
    let mut sum = 0;

    for mut update in updates {
        if !ordered(instructions, &update) {
            reorder(instructions, &mut update);
            let idx = update.len() / 2;
            let middle = update[idx];
            sum += middle;
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
        let (instructions, updates) = read_file(file).unwrap();
        println!("{:?}\n{:?}", instructions, updates);
    }

    #[case("ex1.txt", 123)]
    #[case("input.txt", 4719)]
    fn test_part2(file: &str, expected: u32) {
        let (instructions, updates) = read_file(file).unwrap();
        let result = part2(&instructions, updates);
        assert_eq!(result, expected);
    }
}
