use tap::Pipe;

#[derive(Default, Debug)]
pub struct Almanac {
    pub soil_to_seed: Map,
    pub fertilizer_to_soil: Map,
    pub water_to_fertilizer: Map,
    pub light_to_water: Map,
    pub temperature_to_light: Map,
    pub humidity_to_temperature: Map,
    pub location_to_humidity: Map,
}

#[derive(Default, Debug)]
pub struct Map(Vec<Range>);

#[derive(Default, Debug, Clone, Copy)]
pub struct Range {
    pub start: i64,
    pub end: i64,
    pub offset: i64,
}

impl Map {
    pub fn source(&self, dest: i64) -> i64 {
        for range in &self.0 {
            if range.start > dest {
                return dest;
            }
            if range.end >= dest {
                return dest + range.offset;
            }
        }
        dest
    }
}

impl Almanac {
    pub fn seed(&self, location: i64) -> i64 {
        self.location_to_humidity
            .source(location)
            .pipe(|n| self.humidity_to_temperature.source(n))
            .pipe(|n| self.temperature_to_light.source(n))
            .pipe(|n| self.light_to_water.source(n))
            .pipe(|n| self.water_to_fertilizer.source(n))
            .pipe(|n| self.fertilizer_to_soil.source(n))
            .pipe(|n| self.soil_to_seed.source(n))
    }
}

pub fn parse_almanac(input: &str) -> Almanac {
    let mut almanac = Almanac::default();

    let construct_map = |map: &mut Map, lines: &&str| {
        for range in lines.split('\n') {
            match range.splitn(3, ' ').collect::<Vec<_>>().as_slice() {
                [dest_start, source_start, len] => {
                    let dest_start = dest_start.parse::<i64>().unwrap();
                    let source_start = source_start.parse::<i64>().unwrap();
                    let len = len.parse::<i64>().unwrap();

                    map.0.push(Range {
                        start: dest_start,
                        end: dest_start + len - 1,
                        offset: source_start - dest_start,
                    })
                }
                _ => panic!(),
            }
        }
        map.0.sort_by_key(|r| r.start)
    };

    let categories = input.split("\n\n");
    for cat in categories {
        match cat.split(" map:\n").collect::<Vec<_>>().as_slice() {
            ["seed-to-soil", rest] => construct_map(&mut almanac.soil_to_seed, rest),
            ["soil-to-fertilizer", rest] => construct_map(&mut almanac.fertilizer_to_soil, rest),
            ["fertilizer-to-water", rest] => construct_map(&mut almanac.water_to_fertilizer, rest),
            ["water-to-light", rest] => construct_map(&mut almanac.light_to_water, rest),
            ["light-to-temperature", rest] => {
                construct_map(&mut almanac.temperature_to_light, rest)
            }
            ["temperature-to-humidity", rest] => {
                construct_map(&mut almanac.humidity_to_temperature, rest)
            }
            ["humidity-to-location", rest] => {
                construct_map(&mut almanac.location_to_humidity, rest)
            }
            [_] => (),
            other => panic!("{:?}", other),
        }
    }

    almanac
}

#[cfg(test)]
mod test {

    use super::*;
    use test_case::case;

    #[ignore]
    #[case("ex1.txt")]
    fn test_parse_almanac(input: &str) {
        let input = std::fs::read_to_string(format!("src/day5/{}", input)).unwrap();
        let almanac = parse_almanac(&input);
        dbg!(almanac);
    }

    #[case(0 => 0)]
    #[case(1 => 1)]
    #[case(53 => 51)]
    #[case(98 => 96)]
    #[case(99 => 97)]
    #[case(51 => 99)]
    fn test_soil_to_seed(source: i64) -> i64 {
        let input = std::fs::read_to_string("src/day5/ex1.txt").unwrap();
        let almanac = parse_almanac(&input);
        almanac.soil_to_seed.source(source)
    }

    #[case(46 => 82)]
    fn test_location_to_seed(location: i64) -> i64 {
        let input = std::fs::read_to_string("src/day5/ex1.txt").unwrap();
        let almanac = parse_almanac(&input);
        almanac.seed(location)
    }
}
