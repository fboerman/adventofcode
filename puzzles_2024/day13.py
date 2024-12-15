from ortools.linear_solver import pywraplp
import re


def day13(data_input, part2=False):
    machines = []
    for i, machine_config in enumerate(data_input.split('\n\n')):
        machine_config = machine_config.split('\n')
        A = [int(x) for x in re.findall(r'\d+', machine_config[0])]
        B = [int(x) for x in re.findall(r'\d+', machine_config[1])]
        prize = [int(x) for x in re.findall(r'\d+', machine_config[2])]
        if part2:
            prize[0] += 10000000000000
            prize[1] += 10000000000000

        solver = pywraplp.Solver('adventofcode day13', pywraplp.Solver.SAT_INTEGER_PROGRAMMING)
        variables = {
            'A': solver.NumVar(0, 100 if not part2 else solver.infinity(), 'A'),
            'B': solver.NumVar(0, 100 if not part2 else solver.infinity(), 'B')
        }
        solver.Add(variables['A'] * A[0] + variables['B'] * B[0] == prize[0], name='X')
        solver.Add(variables['A'] * A[1] + variables['B'] * B[1] == prize[1], name='Y')

        solver.Minimize(3*variables['A']+variables['B'])

        status = solver.Solve()
        if status == pywraplp.Solver.OPTIMAL:
            machines.append(int(solver.Objective().Value()))
        # print(f'Machine {i}')
        # if status == pywraplp.Solver.OPTIMAL:
        #     print(int(solver.Objective().Value()))
        #     print(int(variables['A'].solution_value()), int(variables['B'].solution_value()))
        # else:
        #     print("no solution")


    return sum(machines)