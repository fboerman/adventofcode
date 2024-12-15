def get_input(day, test):
    with open(f'puzzle_inputs/day{day}_{"test" if test else "input"}.txt', 'r') as stream:
        return stream.read()

def get_item(M, y, x):
    if y >= M.shape[0] or y < 0:
        return
    if x >= M.shape[1] or x < 0:
        return
    return M[y,x]

def chunks(lst, n):
    #https://stackoverflow.com/a/312464
    """Yield successive n-sized chunks from lst."""
    for i in range(0, len(lst), n):
        yield lst[i:i + n]