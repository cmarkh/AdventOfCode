import day5


def test_ex1():
    with open("ex1.txt") as f:
        inp = f.read()
    seeds = day5.seeds(inp)
    almanac = day5.parse_almanac(inp)
    assert day5.part2(seeds, almanac) == 46


def test_input():
    with open("input.txt") as f:
        inp = f.read()
    seeds = day5.seeds(inp)
    almanac = day5.parse_almanac(inp)
    assert day5.part2(seeds, almanac) == 5200543
