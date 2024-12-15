import numpy as np


def day1_parse(data_input):
    return np.array([x.split('   ') for x in data_input.split('\n') if x != ''])\
        .astype(int)


def day1_part1(data_input):
    M = day1_parse(data_input)
    M[:, 0].sort()
    M[:, 1].sort()
    return np.sum(np.abs(M[:, 0] - M[:, 1]))


def day1_part2(data_input):
    M = day1_parse(data_input)
    #https://stackoverflow.com/a/28663910
    occurences = dict(
        zip(
            *np.unique(M[:, 1], return_counts=True)
        )
    )

    return np.sum([occurences.get(x, 0)*x for x in M[:, 0]])
