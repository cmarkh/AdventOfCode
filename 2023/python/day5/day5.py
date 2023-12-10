from dataclasses import dataclass


@dataclass
class SeedRange:
    start: int
    end: int


@dataclass
class Seeds:
    seeds: list[SeedRange]

    def contains(self, seed: int) -> bool:
        for rng in self.seeds:
            if seed < rng.start:
                return False
            if seed <= rng.end:
                return True
        return False


def seeds(inp: str) -> Seeds:
    seeds = Seeds([])

    line = inp.splitlines()[0]
    line = line.strip("seeds: ")
    nums = line.split(" ")
    for i in range(0, len(nums), 2):
        start = int(nums[i])
        end = int(nums[i + 1]) + start
        seeds.seeds.append(SeedRange(start, end))

    seeds.seeds = sorted(seeds.seeds, key=lambda x: x.start)
    return seeds


@dataclass
class Range:
    start: int
    end: int
    offset: int


@dataclass
class Map:
    rng: list[Range]

    def source(self, dest: int) -> int:
        for rng in self.rng:
            if rng.start > dest:
                return dest
            if rng.end >= dest:
                return dest + rng.offset
        return dest


@dataclass
class Almanac:
    soil_to_seed: Map
    fertilizer_to_soil: Map
    water_to_fertilizer: Map
    light_to_water: Map
    temperature_to_light: Map
    humidity_to_temperature: Map
    location_to_humidity: Map

    def seed(self, location):
        hum = self.location_to_humidity.source(location)
        temp = self.humidity_to_temperature.source(hum)
        light = self.temperature_to_light.source(temp)
        water = self.light_to_water.source(light)
        fert = self.water_to_fertilizer.source(water)
        soil = self.fertilizer_to_soil.source(fert)
        seed = self.soil_to_seed.source(soil)
        return seed


def parse_almanac(inp: str) -> Almanac:
    almanac = Almanac(Map([]), Map([]), Map([]), Map([]), Map([]), Map([]), Map([]))

    def construct_map(map: Map, lines: str):
        for rng in lines.split("\n"):
            rng = rng.split(" ")
            dest_start = int(rng[0])
            source_start = int(rng[1])
            length = int(rng[2])

            map.rng.append(
                Range(dest_start, dest_start + length - 1, source_start - dest_start)
            )
        map.rng.sort(key=lambda x: x.start)

    categories = inp.split("\n\n")
    for cat in categories:
        cat = cat.split(" map:\n")
        if cat[0] == "seed-to-soil":
            construct_map(almanac.soil_to_seed, cat[1])
        elif cat[0] == "soil-to-fertilizer":
            construct_map(almanac.fertilizer_to_soil, cat[1])
        elif cat[0] == "fertilizer-to-water":
            construct_map(almanac.water_to_fertilizer, cat[1])
        elif cat[0] == "water-to-light":
            construct_map(almanac.light_to_water, cat[1])
        elif cat[0] == "light-to-temperature":
            construct_map(almanac.temperature_to_light, cat[1])
        elif cat[0] == "temperature-to-humidity":
            construct_map(almanac.humidity_to_temperature, cat[1])
        elif cat[0] == "humidity-to-location":
            construct_map(almanac.location_to_humidity, cat[1])

    return almanac


def part2(seeds: Seeds, almanac: Almanac) -> int:
    for loc in range(0, 2**64):
        if seeds.contains(almanac.seed(loc)):
            return loc
    return -1
