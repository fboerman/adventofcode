from puzzles_2024 import *


def get_input(day, test):
    with open(f'puzzle_inputs/day{day}_{"test" if test else "input"}.txt', 'r') as stream:
        return stream.read()


def test_day1():
    assert day1_part1(get_input(1, test=True)) == 11