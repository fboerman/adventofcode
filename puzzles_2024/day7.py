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