from . import day1


def test_part1():
    with open("1/input.txt") as f:
        inp = f.read()
    ans = day1.part1(inp)
    assert ans == 56049


def test_part2():
    with open("1/input.txt") as f:
        inp = f.read()
    ans = day1.part2(inp)
    assert ans == 54530


def test_ex1():
    with open("1/ex1.txt") as f:
        inp = f.read()
    ans = day1.part1(inp)
    assert ans == 142


def test_ex2():
    with open("1/ex2.txt") as f:
        inp = f.read()
    ans = day1.part2(inp)
    assert ans == 281
