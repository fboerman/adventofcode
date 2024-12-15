

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