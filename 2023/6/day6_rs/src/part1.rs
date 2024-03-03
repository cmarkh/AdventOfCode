use std::error::Error;

#[derive(Debug, Clone, Copy)]
pub struct Record {
    pub time: u32,
    pub distance: u32,
}

pub fn parse_input(input: &str) -> Result<Vec<Record>, Box<dyn Error>> {
    let mut records = Vec::new();

    let mut input = input.lines();
    let times = input
        .next()
        .ok_or("missing")?
        .trim_start_matches("Time:")
        .split_ascii_whitespace();
    let distances = input
        .next()
        .ok_or("missing")?
        .trim_start_matches("Distance:")
        .split_ascii_whitespace();

    for (t, d) in times.zip(distances) {
        records.push(Record {
            time: t.parse()?,
            distance: d.parse()?,
        })
    }

    Ok(records)
}

pub fn part1(records: &[Record]) -> Result<u32, Box<dyn Error>> {
    let mut product = 1;

    for record in records {
        let mut ways = 0;
        for speed in 0..record.time {
            let time_remaining = record.time - speed;
            let distance = time_remaining * speed;
            if distance > record.distance {
                ways += 1;
            } else if ways != 0 {
                break;
            }
        }
        product *= ways;
    }

    Ok(product)
}

#[cfg(test)]
mod test {
    use super::*;

    use test_case::case;

    #[ignore]
    #[case("ex.txt")]
    #[case("input.txt")]
    fn test_parse_intput(input: &str) {
        let input = std::fs::read_to_string(format!("src/{}", input)).unwrap();
        let records = parse_input(&input).unwrap();
        dbg!(records);
    }

    #[case("ex.txt" => 288)]
    #[case("input.txt" => 503424)]
    fn test_part1(input: &str) -> u32 {
        let input = std::fs::read_to_string(format!("src/{}", input)).unwrap();
        let records = parse_input(&input).unwrap();
        part1(&records).unwrap()
    }
}
