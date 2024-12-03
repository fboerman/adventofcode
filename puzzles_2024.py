import numpy as np


def day1_part1(data_input):
    M = np.array([x.split('   ') for x in data_input.split('\n') if x != ''])\
        .astype(int)
    M[:, 0].sort()
    M[:, 1].sort()

    return np.sum(np.abs(M[:, 0] - M[:, 1]))


def day1_part2(data_input):
    #https://stackoverflow.com/a/28663910
    pass