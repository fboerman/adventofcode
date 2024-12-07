import numpy as np
import re


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


def day3_part1(data_input):
    return np.sum(
        [int(x[0])*int(x[1]) for x in re.findall(r'mul\(([0-9]{1,3}),([0-9]{1,3})\)', data_input.replace('\n',''))]
    )


def day3_part2(data_input):
    result = 0
    enabled = True
    for instruction in re.findall(r"(mul\([0-9]{1,3},[0-9]{1,3}\))|(don't\(\))|(do\(\))", data_input.replace('\n', '')):
        if instruction[0] != '' and enabled:
            result += np.sum([int(x[0])*int(x[1]) for x in re.findall(r'([0-9]{1,3}),([0-9]{1,3})', instruction[0])])
        elif instruction[1] != '':
            enabled = False
        elif instruction[2] != '':
            enabled = True
    return result


def get_item(M, y, x):
    if y >= M.shape[0] or y < 0:
        return
    if x >= M.shape[1] or x < 0:
        return
    return M[y,x]


def day4_part1(data_input):
    xmas_count = 0

    M = np.array([list(l) for l in data_input.split('\n') if l != ''])
    for y in range(M.shape[0]):
        for x in range(M.shape[1]):
            if get_item(M, y, x) == 'X':
                directions = {}
                x_ = 0
                y_ = 0
                for a in ['up', 'down', '']:
                    if a == 'up':
                        y_ = - 1
                    elif a == 'down':
                        y_ = 1
                    else:
                        y_ = 0
                    for b in ['left', 'right', '']:
                        if a == '' and b == '':
                            continue
                        if b == 'left':
                            x_ = -1
                        elif b == 'right':
                            x_ = 1
                        else:
                            x_ = 0
                        directions[f'{b}{a}'] = ''.join([get_item(M, y+(y_*i), x+(x_*i)) or '' for i in range(4)])
                # for direction, string in directions.items():
                #     if string == 'XMAS':
                #         print(x, y, direction)
                xmas_count += sum([x == 'XMAS' for x in directions.values()])

    return xmas_count


def day4_part2(data_input):
    xmas_count = 0

    M = np.array([list(l) for l in data_input.split('\n') if l != ''])
    for y in range(M.shape[0]):
        for x in range(M.shape[1]):
            if get_item(M, y, x) == 'A':
                w1 = ''.join([
                    get_item(M, y-1, x-1) or '', 'A', get_item(M, y+1, x+1) or ''
                ])
                w2 = ''.join([
                    get_item(M, y - 1, x + 1) or '', 'A', get_item(M, y + 1, x - 1) or ''
                ])
                if w1 in ['SAM', 'MAS'] and w2 in ['SAM', 'MAS']:
                    xmas_count += 1

    return xmas_count


def day5_part1(data_input):
    result = 0

    rules_input, updates_input = data_input.split('\n\n')
    rules_before = {}
    rules_after = {}
    for line in rules_input.split('\n'):
        A, B = line.split('|')

        if A not in rules_before:
            rules_before[A] = []
        rules_before[A].append(B)

        if B not in rules_after:
            rules_after[B] = []
        rules_after[B].append(A)

    for i, update in enumerate(updates_input.split('\n')):
        if update == '':
            continue
        L = update.split(',')
        P = {k: v for v, k in enumerate(L)}
        correct_flag = True
        for n in P.keys():
            if n in rules_before:
                if not all([
                    P[n] < P[x]
                    for x in rules_before[n]
                    if x in P
                ]):
                    correct_flag = False
                    break

            if n in rules_after:
                if not all([
                    P[n] > P[x]
                    for x in rules_after[n]
                    if x in P
                ]):
                    correct_flag = False
                    break
        if correct_flag:
            #print(update)
            result += int(L[int((len(L)-1)/2)])

    return result
