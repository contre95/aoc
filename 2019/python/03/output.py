#!python3


def read_input():
    cables = []
    with open('input', 'r') as f:
        for line in f.readlines():
            line = line.strip('\n')
            cables.append(list(line.split(',')))
    return cables


def circuit(cable):
    # global new_pos
    c = set()
    last_pos = (0, 0)
    for m in cable:
        if m[0] == 'R':
            for x in range(abs(int(m[1::]))):
                new_pos = (last_pos[0] + 1, last_pos[1])
                c.add(new_pos)
                last_pos = new_pos
        elif m[0] == 'L':
            for x in range(abs(int(m[1::]))):
                new_pos = (last_pos[0] - 1, last_pos[1])
                c.add(new_pos)
                last_pos = new_pos
        elif m[0] == 'U':
            for x in range(abs(int(m[1::]))):
                new_pos = (last_pos[0], last_pos[1] + 1)
                c.add(new_pos)
                last_pos = new_pos
        elif m[0] == 'D':
            for x in range(abs(int(m[1::]))):
                new_pos = (last_pos[0], last_pos[1] - 1)
                c.add(new_pos)
                last_pos = new_pos
    return c


def main():
    cables = read_input()
    c1 = circuit(cables[0])
    c2 = circuit(cables[1])
    print(min(list(map(lambda tup: abs(tup[0]) + abs(tup[1]), c1.intersection(c2)))))


main()
