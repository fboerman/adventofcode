import numpy as np
from .aoc_utils import get_item

def day10(data_input, part2=False):
    M = np.array([np.array(list(x)).astype(int) for x in data_input.split('\n') if x != ''])
    trailheads = []

    def walk_trail(M, y, x, reachable, rating):
        n = get_item(M, y, x)
        if n == 9:
            reachable.add((y, x))
            rating += 1
            return reachable, rating
        for y_, x_ in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            n_ = get_item(M, y + y_, x + x_)
            if n_ is None:
                continue
            if n_ - n == 1:
                reachable, rating = walk_trail(M, y + y_, x + x_, reachable, rating)
        return reachable, rating

    for y in range(M.shape[0]):
        for x in range(M.shape[1]):
            cell = get_item(M, y, x)
            if cell == 0:
                trailheads.append(walk_trail(M, y, x, set(), 0))
    if not part2:
        return sum([len(x[0]) for x in trailheads])
    else:
        return sum([x[1] for x in trailheads])