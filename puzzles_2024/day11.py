from functools import lru_cache


def day11(data_input, part2=False):
    stones = [int(x) for x in data_input.split('\n')[0].split(' ')]

    @lru_cache(maxsize=None)
    def blink(n, to_go):
        if to_go == 0:
            return 1
        if n == 0:
            return blink(1, to_go-1)
        elif len(str(n)) % 2 == 0:
            return blink(int(str(n)[:len(str(n)) // 2]), to_go-1) + \
                    blink(int(str(n)[len(str(n)) // 2:]), to_go-1)
        else:
            return blink(n*2024,to_go-1)

    return sum([blink(n, 25 if not part2 else 75) for n in stones])