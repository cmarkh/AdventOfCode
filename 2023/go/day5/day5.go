package day5

import (
	"bufio"
	"fmt"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start  int64
	End    int64
	Offset int64
}

type Map struct {
	Ranges []Range
}

func (m *Map) Source(dest int64) int64 {
	for _, rng := range m.Ranges {
		if rng.Start > dest {
			return dest
		}
		if rng.End >= dest {
			return dest + rng.Offset
		}
	}
	return dest
}

type Almanac struct {
	SoilToSeed            Map
	FertilizerToSoil      Map
	WaterToFertilizer     Map
	LightToWater          Map
	TemperatureToLight    Map
	HumidityToTemperature Map
	LocationToHumidity    Map
}

func (a *Almanac) Seed(location int64) int64 {
	loc := a.LocationToHumidity.Source(location)
	loc = a.HumidityToTemperature.Source(loc)
	loc = a.TemperatureToLight.Source(loc)
	loc = a.LightToWater.Source(loc)
	loc = a.WaterToFertilizer.Source(loc)
	loc = a.FertilizerToSoil.Source(loc)
	loc = a.SoilToSeed.Source(loc)
	return loc
}

func ConstructMap(m *Map, lines string) {
	scanner := bufio.NewScanner(strings.NewReader(lines))
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) == 3 {
			destStart, _ := strconv.ParseInt(parts[0], 10, 64)
			sourceStart, _ := strconv.ParseInt(parts[1], 10, 64)
			length, _ := strconv.ParseInt(parts[2], 10, 64)
			m.Ranges = append(m.Ranges, Range{Start: destStart, End: destStart + length - 1, Offset: sourceStart - destStart})
		} else {
			panic("Invalid range format")
		}
	}
	slices.SortFunc(m.Ranges, func(a, b Range) int {
		if a.Start < b.Start {
			return -1
		} else if a.Start > b.Start {
			return 1
		} else {
			return 0
		}
	})
}

func ParseAlmanac(input string) Almanac {
	almanac := Almanac{}

	categories := strings.Split(input, "\n\n")
	for _, cat := range categories {
		parts := strings.SplitN(cat, " map:\n", 2)
		if len(parts) == 2 {
			switch parts[0] {
			case "seed-to-soil":
				ConstructMap(&almanac.SoilToSeed, parts[1])
			case "soil-to-fertilizer":
				ConstructMap(&almanac.FertilizerToSoil, parts[1])
			case "fertilizer-to-water":
				ConstructMap(&almanac.WaterToFertilizer, parts[1])
			case "water-to-light":
				ConstructMap(&almanac.LightToWater, parts[1])
			case "light-to-temperature":
				ConstructMap(&almanac.TemperatureToLight, parts[1])
			case "temperature-to-humidity":
				ConstructMap(&almanac.HumidityToTemperature, parts[1])
			case "humidity-to-location":
				ConstructMap(&almanac.LocationToHumidity, parts[1])
			default:
				panic("Unknown category")
			}
		}
	}

	return almanac

}

type SeedRange struct {
	Start int64
	End   int64
}

type Seeds struct {
	Ranges []SeedRange
}

func (s *Seeds) Contains(seed int64) bool {
	for _, r := range s.Ranges {
		if seed < r.Start {
			return false
		}
		if seed <= r.End {
			return true
		}
	}
	return false
}

func ParseSeeds(input string) (*Seeds, error) {
	seeds := &Seeds{}

	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("no lines in input")
	}

	seedStr := strings.TrimPrefix(lines[0], "seeds: ")
	nums := strings.Fields(seedStr)

	for i := 0; i < len(nums); i += 2 {
		start, err := strconv.ParseInt(nums[i], 10, 64)
		if err != nil {
			return nil, err
		}
		end, err := strconv.ParseInt(nums[i+1], 10, 64)
		if err != nil {
			return nil, err
		}
		seeds.Ranges = append(seeds.Ranges, SeedRange{Start: start, End: start + end})
	}

	sort.Slice(seeds.Ranges, func(i, j int) bool {
		return seeds.Ranges[i].Start < seeds.Ranges[j].Start
	})

	return seeds, nil
}

func Part2(seeds Seeds, almanac Almanac) int64 {
	for i := int64(0); i < int64(math.Pow(6, 64)); i++ {
		if seeds.Contains(almanac.Seed(i)) {
			return i
		}
	}
	return -1
}
