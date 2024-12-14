import numpy as np
import re
from tqdm import tqdm
import itertools


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


def day6(data_input):
    M = np.array([np.array(list(x)) for x in data_input.split('\n') if x != ''])
    # find current position
    start_x = None
    start_y = None
    for y_ in range(M.shape[0]):
        for x_ in range(M.shape[1]):
            if get_item(M, y_, x_) == '^':
                start_x = x_
                start_y = y_
                break
        if start_x is not None:
            break

    directions = [
        (0, -1), #x, y step
        (1,  0),
        (0,  1),
        (-1, 0)
    ]
    # recursive make the walk
    def walk_until_next(M, x, y, dir_i, visited_cells):
        while True:
            step_x, step_y = directions[dir_i]
            next_step = get_item(M, y + step_y, x + step_x)
            if next_step is None:
                return visited_cells
            elif next_step == '#':
                #print(visited_cells)
                return walk_until_next(M, x, y, (dir_i + 1)%4, visited_cells)
            else:
                y += step_y
                x += step_x
                if (x, y, dir_i) in visited_cells:
                    return
                visited_cells.append((x, y, dir_i))
    visited_cells_original = walk_until_next(M, start_x, start_y, 0, [(start_x, start_y)])

    part1 = len(set([(x[0], x[1]) for x in visited_cells_original]))
    part2 = 0
    visited_cells_original = set([(x[0], x[1]) for x in visited_cells_original[1:]])

    for cell in tqdm(visited_cells_original):
        M_ = M.copy()
        M_[cell[1], cell[0]] = '#'
        visited_cells_loop = walk_until_next(M_, start_x, start_y, 0, [(start_x, start_y)])
        if visited_cells_loop is None:
            part2 += 1

    return part1, part2


def day7(data_input, part2=False):
    result = 0
    if part2:
        ops_list = ['*', '+', '||']
    else:
        ops_list = ['*', '+']

    def solve(target, target_so_far, remaining, operator):
        if operator == '*':
            target_so_far *= remaining[0]
        elif operator == '+':
            target_so_far += remaining[0]
        elif operator == '||':
            target_so_far = int(str(target_so_far)+str(remaining[0]))

        if target_so_far == target and len(remaining) == 1:
            return True
        if len(remaining) == 1 or target_so_far > target:
            return False

        for new_op in ops_list:
            if solve(target, target_so_far, remaining[1:], new_op):
                return True
        return False

    for equation in data_input.split('\n'):
        if equation == '':
            continue
        parts = [int(x) for x in equation.split(': ')[1].split(' ')]
        target = int(equation.split(': ')[0])
        for start_operator in ops_list:
            if solve(
                target,
                parts[0],
                parts[1:],
                start_operator
            ):
                result += target
                break

    return result


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


def chunks(lst, n):
    #https://stackoverflow.com/a/312464
    """Yield successive n-sized chunks from lst."""
    for i in range(0, len(lst), n):
        yield lst[i:i + n]


def day9_part1(data_input):
    data_input = [int(x) for x in data_input.split('\n')[0]]
    checksum = 0

    fileblocks = []
    last_position = 0
    for i in range(0, len(data_input), 2):
        b = data_input[i]
        fileblocks.append({
            'id': int(i/2),
            'num': b
        })

        try:
            last_position += b + data_input[i+1]
        except IndexError:
            # there was no space anymore
            break



    def print_blocks(fileblocks):
        for block in fileblocks:
            print("".join([str(x) for x in [block['id']] * block['num']]), end='')
        print()

    working_block = fileblocks.pop()
    shift = 1
    for freespace_i in range(1, len(data_input), 2):
        freespace = data_input[freespace_i]
        while freespace != 0:
            if working_block['num'] <= freespace:
                # move the whole block
                fileblocks.insert(shift, working_block)
                freespace -= working_block['num']
                working_block = fileblocks.pop()
            else:
                # split and move the block
                working_block['num'] -= freespace
                fileblocks.insert(shift, {
                    'id': working_block['id'],
                    'num': freespace
                })
                freespace = 0
            #print(shift)
            #print_blocks(fileblocks)
            shift += 1
        #print(shift)
        #print_blocks(fileblocks)
        shift += 1
    fileblocks.append(working_block)

    last_position = 0
    def calc_checksum(file_id, num, last_position):
        return file_id * np.sum(np.array(last_position + np.array(range(num))))
    for block in fileblocks:
        checksum += calc_checksum(block['id'], block['num'], last_position)
        last_position += block['num']
    return checksum


def day9_part2(data_input):
    def print_blocks(fileblocks):
        for block in fileblocks:
            print("".join([str(x if x is not None else '.') for x in [block['id']] * block['num']]), end='')
        print()

    data_input = [int(x) for x in data_input.split('\n')[0]]
    checksum = 0

    fileblocks = []

    for i in range(0, len(data_input), 2):
        fileblocks.append({
            'id': int(i/2),
            'num': data_input[i]
        })
        try:
            fileblocks.append({
                'id': None,
                'num': data_input[i+1]
            })
        except IndexError:
            break

    for block in reversed(fileblocks):
        if block['id'] is None:
            continue
        for i, freespace in enumerate(fileblocks):
            if freespace == block:
                break
            if freespace['id'] is not None or freespace['num'] < block['num']:
                continue
            freespace['num'] -= block['num']
            fileblocks.insert(i, dict(block))
            block['id'] = None
            #print_blocks(fileblocks)
            break

    #print_blocks(fileblocks)

    last_position = 0
    def calc_checksum(file_id, num, last_position):
        return file_id * np.sum(np.array(last_position + np.array(range(num))))
    for block in fileblocks:
        if block['id'] is None:
            last_position += block['num']
            continue
        checksum += calc_checksum(block['id'], block['num'], last_position)
        last_position += block['num']
    return checksum

