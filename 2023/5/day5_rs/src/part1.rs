use tap::Pipe;

#[derive(Default, Debug)]
pub struct Almanac {
    pub seeds: Vec<i64>,
    pub seed_to_soil: Map,
    pub soil_to_fertilizer: Map,
    pub fertilizer_to_water: Map,
    pub water_to_light: Map,
    pub light_to_temperature: Map,
    pub temperature_to_humidity: Map,
    pub humidity_to_location: Map,
}

#[derive(Default, Debug)]
pub struct Map(Vec<Range>);

#[derive(Default, Debug)]
pub struct Range {
    pub start: i64,
    pub end: i64,
    pub offset: i64,
}

impl Map {
    pub fn destination(&self, source: i64) -> i64 {
        for range in &self.0 {
            if range.start > source {
                return source;
            }
            if range.end >= source {
                return source + range.offset;
            }
        }
        source
    }
}

pub fn parse_input(input: &str) -> Almanac {
    let mut almanac = Almanac::default();

    let construct_map = |map: &mut Map, lines: &&str| {
        for range in lines.split('\n') {
            match range.splitn(3, ' ').collect::<Vec<_>>().as_slice() {
                [dest_start, source_start, len] => {
                    let dest_start = dest_start.parse::<i64>().unwrap();
                    let source_start = source_start.parse::<i64>().unwrap();
                    let len = len.parse::<i64>().unwrap();

                    map.0.push(Range {
                        start: source_start,
                        end: source_start + len - 1,
                        offset: dest_start - source_start,
                    })
                }
                _ => panic!(),
            }
        }
        map.0.sort_by_key(|r| r.start)
    };

    let categories = input.split("\n\n");
    for cat in categories {
        if let Some((_, seeds)) = cat.split_once("seeds: ") {
            almanac.seeds = seeds
                .split_ascii_whitespace()
                .map(|n| n.parse::<i64>().unwrap())
                .collect();
            continue;
        }
        match cat.split(" map:\n").collect::<Vec<_>>().as_slice() {
            ["seed-to-soil", rest] => construct_map(&mut almanac.seed_to_soil, rest),
            ["soil-to-fertilizer", rest] => construct_map(&mut almanac.soil_to_fertilizer, rest),
            ["fertilizer-to-water", rest] => construct_map(&mut almanac.fertilizer_to_water, rest),
            ["water-to-light", rest] => construct_map(&mut almanac.water_to_light, rest),
            ["light-to-temperature", rest] => {
                construct_map(&mut almanac.light_to_temperature, rest)
            }
            ["temperature-to-humidity", rest] => {
                construct_map(&mut almanac.temperature_to_humidity, rest)
            }
            ["humidity-to-location", rest] => {
                construct_map(&mut almanac.humidity_to_location, rest)
            }
            other => panic!("{:?}", other),
        }
    }

    almanac
}

pub fn part1(almanac: Almanac) -> i64 {
    let mut lowest = i64::MAX;

    for seed in almanac.seeds {
        let location = almanac
            .seed_to_soil
            .destination(seed)
            .pipe(|n| almanac.soil_to_fertilizer.destination(n))
            .pipe(|n| almanac.fertilizer_to_water.destination(n))
            .pipe(|n| almanac.water_to_light.destination(n))
            .pipe(|n| almanac.light_to_temperature.destination(n))
            .pipe(|n| almanac.temperature_to_humidity.destination(n))
            .pipe(|n| almanac.humidity_to_location.destination(n));
        lowest = lowest.min(location);
    }

    lowest
}

#[cfg(test)]
mod test {
    use test_case::case;

    use super::*;

    #[ignore]
    #[case("ex1.txt")]
    fn test_input(input_name: &str) {
        let input = std::fs::read_to_string(format!("src/{}", input_name)).unwrap();
        let parse = parse_input(&input);
        dbg!(&parse);
    }

    #[case(0 => 0)]
    #[case(1 => 1)]
    #[case(51 => 53)]
    #[case(96 => 98)]
    #[case(97 => 99)]
    #[case(99 => 51)]
    fn test_seed_to_soil(source: i64) -> i64 {
        let input = std::fs::read_to_string("src/ex1.txt").unwrap();
        let almanac = parse_input(&input);
        almanac.seed_to_soil.destination(source)
    }

    #[case("ex1.txt" => 35)]
    #[case("input.txt" => 175622908)]
    fn test_part1(input_name: &str) -> i64 {
        let input = std::fs::read_to_string(format!("src/{}", input_name)).unwrap();
        let almanac = parse_input(&input);
        part1(almanac)
    }
}
