#![allow(dead_code)]

use color_eyre::{eyre::OptionExt, Result};

type Disk = Vec<Cell>;

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
enum Cell {
    FreeSpace,
    Block(u64),
}

fn print(disk: &Disk) {
    for cell in disk {
        match cell {
            Cell::FreeSpace => print!("."),
            Cell::Block(idx) => print!("{}", idx),
        }
    }
    println!();
}

fn read_file(file: &str) -> Result<Disk> {
    let mut disk = Disk::new();

    let input = std::fs::read_to_string(file)?;
    let input = input.trim();

    for (i, ch) in input.chars().enumerate() {
        if i % 2 == 0 {
            let idx = (i / 2) as u64;
            let num = ch.to_digit(10).ok_or_eyre("Invalid digit")?;
            disk.extend(vec![Cell::Block(idx); num as usize]);
        } else {
            let num = ch.to_digit(10).ok_or_eyre("Invalid digit")?;
            disk.extend(vec![Cell::FreeSpace; num as usize]);
        }
    }

    Ok(disk)
}

fn part1(disk: &mut Disk) -> u64 {
    let mut checksum = 0;

    let empties = disk
        .iter()
        .enumerate()
        .filter(|(_, &cell)| cell == Cell::FreeSpace)
        .map(|(i, _)| i)
        .collect::<Vec<_>>();

    for empty in empties {
        // print(disk);

        let mut cell = disk.pop().unwrap();
        while cell == Cell::FreeSpace {
            cell = disk.pop().unwrap();
        }

        if empty >= disk.len() {
            disk.push(cell);
            break;
        }

        disk[empty] = cell;
    }

    // print(disk);

    for (i, cell) in disk.iter().enumerate() {
        match cell {
            Cell::Block(val) => checksum += i as u64 * *val,
            Cell::FreeSpace => unreachable!(),
        }
    }

    checksum
}

#[cfg(test)]
mod tests {
    use test_case::case;

    use super::*;

    #[case("ex1.txt")]
    fn test_read_file(file: &str) {
        let disk = read_file(file).unwrap();
        print(&disk);
    }

    #[case("ex1.txt", 1928)]
    #[case("input.txt", 6337367222422)]
    fn test_part1(file: &str, expected: u64) {
        let mut disk = read_file(file).unwrap();
        let checksum = part1(&mut disk);
        assert_eq!(checksum, expected);
    }
}
