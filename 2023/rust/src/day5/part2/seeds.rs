#[derive(Default, Debug)]
pub struct Seeds(Vec<SeedRange>);

#[derive(Default, Debug)]
pub struct SeedRange {
    pub start: i64,
    pub end: i64,
}

impl Seeds {
    pub fn contains(&self, seed: i64) -> bool {
        for range in &self.0 {
            if seed < range.start {
                return false;
            }
            if seed <= range.end {
                return true;
            }
        }
        false
    }
}

pub fn parse_seeds(input: &str) -> Seeds {
    let mut seeds = Seeds::default();

    let mut nums = input
        .lines()
        .next()
        .unwrap()
        .strip_prefix("seeds: ")
        .unwrap()
        .split(' ');

    while let Some(num) = nums.next() {
        let start = num.parse::<i64>().unwrap();
        let end = start + nums.next().unwrap().parse::<i64>().unwrap();
        seeds.0.push(SeedRange { start, end })
    }

    seeds.0.sort_by_key(|s| s.start);
    seeds
}

#[cfg(test)]
mod test {
    use super::*;
    use test_case::case;

    #[ignore]
    #[case("ex1.txt")]
    fn test_parse_seeds(input: &str) {
        let input = std::fs::read_to_string(format!("src/day5/{}", &input)).unwrap();
        let seeds = parse_seeds(&input);
        dbg!(seeds);
    }

    #[case(79 => true )]
    #[case(85 => true )]
    #[case(90 => true )]
    #[case(60 => true )]
    #[case(67 => true )]
    #[case(10 => false )]
    #[case(20 => false )]
    #[case(100 => false )]
    fn test_seeds_contains(seed: i64) -> bool {
        let input = std::fs::read_to_string("src/day5/ex1.txt").unwrap();
        let seeds = parse_seeds(&input);
        seeds.contains(seed)
    }
}
