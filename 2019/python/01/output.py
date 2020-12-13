#!/usr/local/bin/python3

def calc_fuel(mass):
    if mass <= 0:
        return abs(mass)
    return (mass // 3 - 2) + calc_fuel(mass // 3 - 2)


def calc_fuel_one_star(mass):
    return mass // 3 - 2


def main():
    fuel = 0
    with open('input', 'r') as f:
        masses = [int(line.strip('\n')) for line in f.readlines()]
    for mass in masses:
        fuel += calc_fuel(mass)
    return fuel


print(main())

# print(calc_fuel(1969))
# print(calc_fuel(100756))