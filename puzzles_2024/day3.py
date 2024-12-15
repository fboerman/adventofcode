import numpy as np
import re

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