pub mod almanac;
pub mod seeds;

#[timed::timed]
pub fn part2(seeds: &seeds::Seeds, almanac: &almanac::Almanac) -> i64 {
    (0..i64::MAX)
        .find(|loc| seeds.contains(almanac.seed(*loc)))
        .unwrap()
}

#[cfg(test)]
mod test {
    use super::*;
    use test_case::case;

    #[case("ex1.txt" => 46)]
    #[case("input.txt" => 5200543)]
    fn test_part2(input: &str) -> i64 {
        let input = std::fs::read_to_string(format!("src/{}", input)).unwrap();
        let seeds = seeds::parse_seeds(&input);
        let almanac = almanac::parse_almanac(&input);

        part2(&seeds, &almanac)
    }
}
