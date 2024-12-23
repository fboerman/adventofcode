from puzzles_2024 import *
from puzzles_2024.aoc_utils import get_input


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


def test_day7():
    assert day7(get_input(7, test=True), part2=False) == 3749
    assert day7(get_input(7, test=True), part2=True) == 11387


def test_day8():
    assert day8(get_input(8, test=True), part2=False) == 14
    assert day8(get_input(8, test=True), part2=True) == 34


def test_day9():
    assert day9_part1(get_input(9, test=True)) == 1928
    assert day9_part2(get_input(9, test=True)) == 2858


def test_day10():
    assert day10(get_input(10, test=True), part2=False) == 36
    assert day10(get_input(10, test=True), part2=True) == 81


def test_day11():
    assert day11(get_input(11, test=True), part2=False) == 55312


def test_day13():
    assert day13(get_input(13, test=True), part2=False) == 480