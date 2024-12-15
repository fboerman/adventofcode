import numpy as np
from .aoc_utils import  get_item
from tqdm import tqdm


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