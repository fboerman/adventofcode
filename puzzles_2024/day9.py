import numpy as np


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