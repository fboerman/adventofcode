import numpy as np
from .aoc_utils import get_item
import itertools


def day8(data_input, part2=False):
    M = np.array([np.array(list(x)) for x in data_input.split('\n') if x != ''])
    #M_ = M.copy()
    antennas = {}
    antinodes = set()
    for y in range(M.shape[0]):
        for x in range(M.shape[1]):
            cell = get_item(M, y, x)
            if cell != '.' and cell != '#':
                if cell not in antennas:
                    antennas[cell] = []
                antennas[cell].append(np.array([x, y]))
                if part2:
                    antinodes.add((x, y))

    for antenna_list in antennas.values():
        for A, B in itertools.combinations(antenna_list, r=2):
            num_harmonics = 1
            while True:
                option1 = tuple(B + num_harmonics*(B - A))
                option2 = tuple(A - num_harmonics*(B - A))
                flag1 = get_item(M, *reversed(option1))
                flag2 = get_item(M, *reversed(option2))

                if flag1:
                    antinodes.add(option1)
                if flag2:
                    antinodes.add(option2)

                if not part2:
                    break

                if not flag1 and not flag2:
                    break

                num_harmonics += 1

    # for node in antinodes:
    #     M_[node[1], node[0]] = '#'
    # for y in range(M.shape[0]):
    #     for x in range(M.shape[1]):
    #         print(get_item(M_, y, x), end='')
    #     print()
    return len(antinodes)