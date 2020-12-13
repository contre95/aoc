#!/usr/local/bin/python3
import csv


def opcode_one(list, first, second, third):
    list[third] = list[first] + list[second]


def opcode_two(list, first, second, third):
    list[third] = list[first] * list[second]


opcodes = {1: opcode_one, 2: opcode_two, 99: 'exit'}


def read_input():
    with open('input', 'r') as f:
        lista = list(csv.reader(f))
    return list(map(int, lista[0]))


def output(noun, verb):
    intcode = read_input()
    intcode[1] = noun
    intcode[2] = verb
    i = 0
    while i < len(intcode):
        if intcode[i] in opcodes:
            if intcode[i] == 99:
                return intcode[0]
            else:
                opcodes[intcode[i]](intcode, intcode[i + 1], intcode[i + 2], intcode[i + 3])
        i += 4
    return intcode[0]


def main():
    # First part
    print(output(12, 2))
    # Second part
    for noun in range(99):
        for verb in range(99):
            if output(noun, verb) == 19690720:
                print(f"100 * {noun} + {verb} = {100 * noun + verb}")


main()
