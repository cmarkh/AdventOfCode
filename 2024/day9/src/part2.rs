#![allow(dead_code)]

use color_eyre::{eyre::OptionExt, Result};

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
struct Block {
    start: usize,
    len: usize,
    id: i64,
}

fn print(disk: &[i64]) {
    for n in disk {
        if *n == -1 {
            print!(".");
        } else {
            print!("{}", n);
        }
    }
    println!();
}

fn read_file(file: &str) -> Result<(Vec<i64>, Vec<Block>)> {
    let mut disk = Vec::new();
    let mut blocks = Vec::new();

    let input = std::fs::read_to_string(file)?;
    let input = input.trim();

    for (i, ch) in input.chars().enumerate() {
        if i % 2 == 0 {
            let idx = (i / 2) as i64;
            let num = ch.to_digit(10).ok_or_eyre("Invalid digit")?;
            let block = Block {
                start: disk.len(),
                len: num as usize,
                id: idx,
            };
            blocks.push(block);
            for _ in 0..num {
                disk.push(idx);
            }
        } else {
            let num = ch.to_digit(10).ok_or_eyre("Invalid digit")?;
            for _ in 0..num {
                disk.push(-1);
            }
        }
    }

    Ok((disk, blocks))
}

fn part2(mut disk: Vec<i64>, mut blocks: Vec<Block>) -> i64 {
    let mut checksum = 0;

    while let Some(block) = blocks.pop() {
        // println!("{:?}", block);
        // print(&disk);
        for i in 0..disk.len() {
            if i >= block.start {
                break;
            }
            if disk[i..(i + block.len)].iter().all(|&n| n == -1) {
                disk[i..(i + block.len)].copy_from_slice(&vec![block.id; block.len]);
                disk[block.start..(block.start + block.len)].copy_from_slice(&vec![-1; block.len]);
                break;
            }
        }
    }
    // print(&disk);

    for (i, &n) in disk.iter().enumerate() {
        if n == -1 {
            continue;
        }
        checksum += n * i as i64;
    }

    checksum
}

#[cfg(test)]
mod tests {
    use test_case::case;

    use super::*;

    #[case("ex1.txt")]
    fn test_read_file(file: &str) {
        let (disk, blocks) = read_file(file).unwrap();
        print(&disk);
        println!("{:?}", blocks);
    }

    #[case("ex1.txt", 2858)]
    #[case("input.txt", 6361380647183)]
    fn test_part2(file: &str, expected: i64) {
        let (disk, blocks) = read_file(file).unwrap();
        let checksum = part2(disk, blocks);
        assert_eq!(checksum, expected);
    }
}
