import os
from datetime import date

if __name__ == '__main__':
    for d in range(1, date.today().day + 1):
        fname = os.path.join('puzzle_inputs', f'day{d}_input.txt')
        if not os.path.exists(fname):
            command = f'./aoc-cli download -d{d} --input-only -i {fname}'
            os.system(command)