from puzzles_2024 import *


def get_input(day, test):
    with open(f'puzzle_inputs/day{day}_{"test" if test else "input"}.txt', 'r') as stream:
        return stream.read()


def test_day1():
    assert day1_part1(get_input(1, test=True)) == 11
    assert day1_part2(get_input(1, test=True)) == 31


def test_day2():
    assert day2(get_input(2, test=True), 1) == 2
    assert day2(get_input(2, test=True), 2) == 4


def test_day3():
    assert day3_part1(get_input(3, test=True)) == 161
    assert day3_part2(get_input(3, test=True)) == 48


def test_day4():
    assert day4_part1(get_input(4, test=True)) == 18
    assert day4_part2(get_input(4, test=True)) == 9


def test_day5():
    assert day5_part1(get_input(5, test=True)) == 143


def test_day6():
    assert day6(get_input(6, test=True)) == (41, 6)
