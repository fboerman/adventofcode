import numpy as np

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