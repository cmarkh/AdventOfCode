def part1(input: str) -> int:
    total = 0

    for line in input.splitlines():
        str_num = ""
        for ch in line:
            if ch.isnumeric():
                str_num += ch
                break
        for ch in line[::-1]:
            if ch.isnumeric():
                str_num += ch
                break
        total += int(str_num)

    return total


NUMS = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}


def part2(input: str) -> int:
    total = 0

    for line in input.splitlines():
        nums = get_nums(line)

        num = nums[0][1] + nums[-1][1]
        total += int(num)

    return total


def get_nums(line: str) -> list:
    nums = []

    for i, ch in enumerate(line):
        if ch.isnumeric():
            nums.append((i, ch))

    for word, num in NUMS.items():
        idx = line.find(word)
        if idx != -1:
            nums.append((idx, str(num)))
        idx = line.rfind(word)
        if idx != -1:
            nums.append((idx, str(num)))

    return sorted(nums, key=lambda x: x[0])
