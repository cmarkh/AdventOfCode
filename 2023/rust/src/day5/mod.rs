pub mod part1;
pub mod part2;

use std::collections::HashMap;

#[derive(Default, Debug)]
pub struct Almanac {
    pub seeds: Vec<u32>,
    pub seed_to_soil: HashMap<u32, u32>,
    pub soil_to_fertilizer: HashMap<u32, u32>,
    pub fertilizer_to_water: HashMap<u32, u32>,
    pub water_to_light: HashMap<u32, u32>,
    pub light_to_temperature: HashMap<u32, u32>,
    pub temperature_to_humidity: HashMap<u32, u32>,
    pub humidity_to_location: HashMap<u32, u32>,
}

pub fn parse_input(input: &str) -> Almanac {
    let mut almanac = Almanac::default();

    let construct_map = |map: &mut HashMap<u32, u32>, lines: &&str| {
        for range in lines.split('\n') {
            match range.splitn(3, ' ').collect::<Vec<_>>().as_slice() {
                [dest_start, source_start, len] => {
                    let dest_start = dest_start.parse::<u32>().unwrap();
                    let source_start = source_start.parse::<u32>().unwrap();
                    let len = len.parse::<u32>().unwrap();

                    for i in 0..len {
                        map.insert(source_start + i, dest_start + i);
                    }
                }
                _ => panic!(),
            }
        }
    };

    let categories = input.split("\n\n");
    for cat in categories {
        if let Some((_, seeds)) = cat.split_once("seeds: ") {
            almanac.seeds = seeds
                .split_ascii_whitespace()
                .map(|n| n.parse::<u32>().unwrap())
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

#[cfg(test)]
mod test {
    use test_case::case;

    use super::parse_input;

    #[case("ex1.txt")]
    fn test_input(input_name: &str) {
        let input = std::fs::read_to_string(format!("src/day5/{}", input_name)).unwrap();
        let parse = parse_input(&input);
        dbg!(&parse);
    }
}
