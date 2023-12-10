use tap::Pipe;

use super::Almanac;

pub fn part1(almanac: Almanac) -> u32 {
    let mut lowest = u32::MAX;

    for seed in &almanac.seeds {
        dbg!(seed);
        let location = almanac
            .seed_to_soil
            .get(seed)
            .unwrap_or(seed)
            .pipe(|n| dbg!(n))
            .pipe(|n| almanac.soil_to_fertilizer.get(n).unwrap_or(n))
            .pipe(|n| dbg!(n))
            .pipe(|n| almanac.fertilizer_to_water.get(n).unwrap_or(n))
            .pipe(|n| almanac.water_to_light.get(n).unwrap_or(n))
            .pipe(|n| almanac.light_to_temperature.get(n).unwrap_or(n))
            .pipe(|n| almanac.temperature_to_humidity.get(n).unwrap_or(n))
            .pipe(|n| almanac.humidity_to_location.get(n).unwrap_or(n));
        lowest = lowest.min(*location);
    }

    lowest
}

#[cfg(test)]
mod test {
    use test_case::case;

    use super::*;
    use crate::day5;

    #[case("ex1.txt" => 35)]
    #[case("input.txt" => 1)]
    fn test_part1(input_name: &str) -> u32 {
        let input = std::fs::read_to_string(format!("src/day5/{}", input_name)).unwrap();
        let almanac = day5::parse_input(&input);
        part1(almanac)
    }
}
