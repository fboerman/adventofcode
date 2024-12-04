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


def day2(data_input, part):
    safe_count = 0

    if part != 1:
        shift = 1
    else:
        shift = 0

    for line in data_input.split('\n'):
        if line == '':
            continue
        A = np.array(line.split(' ')).astype(int)

        safe = False
        for i in range(len(A)):
            A_i = np.array(A.tolist()[:i] + A.tolist()[i+shift:])
            A_diff = np.diff(A_i, 1)
            safe = (
                np.signbit(A_diff).all() or
                np.invert(np.signbit(A_diff)).all()
            ) and \
            (
                (np.abs(A_diff) < 4).all() and
                (np.abs(A_diff) >= 1).all()
            )
            if safe or shift == 0:
                break
        safe_count += int(safe)
    return safe_count
