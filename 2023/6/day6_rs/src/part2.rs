use std::error::Error;

#[derive(Debug, Clone, Copy)]
pub struct Record {
    pub time: u64,
    pub distance: u64,
}

pub fn parse_input(input: &str) -> Result<Record, Box<dyn Error>> {
    let mut input = input.lines();
    let time = input
        .next()
        .ok_or("missing")?
        .trim_start_matches("Time:")
        .replace(' ', "")
        .parse::<u64>()?;
    let distance = input
        .next()
        .ok_or("missing")?
        .trim_start_matches("Distance:")
        .replace(' ', "")
        .parse::<u64>()?;
    Ok(Record { time, distance })
}

pub fn part2(record: &Record) -> Result<u64, Box<dyn Error>> {
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

    Ok(ways)
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
        let record = parse_input(&input).unwrap();
        dbg!(record);
    }

    #[case("ex.txt" => 71503)]
    #[case("input.txt" => 32607562)]
    fn test_part2(input: &str) -> u64 {
        let input = std::fs::read_to_string(format!("src/{}", input)).unwrap();
        let records = parse_input(&input).unwrap();
        part2(&records).unwrap()
    }
}
