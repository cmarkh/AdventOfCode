import day5


def test_ex1():
    with open("5/ex1.txt") as f:
        inp = f.read()
    seeds = day5.seeds(inp)
    almanac = day5.parse_almanac(inp)
    assert day5.part2(seeds, almanac) == 46


def test_input():
    with open("5/input.txt") as f:
        inp = f.read()
    seeds = day5.seeds(inp)
    almanac = day5.parse_almanac(inp)
    assert day5.part2(seeds, almanac) == 5200543
