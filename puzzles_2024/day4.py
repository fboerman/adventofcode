import numpy as np
from .aoc_utils import get_item


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