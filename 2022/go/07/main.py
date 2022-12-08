import sys
import pprint

pp = pprint.PrettyPrinter(width=41, compact=True)

def main():
    stack = []
    dirs = {}
    with open(sys.argv[1]) as file:
        for line in file:
            line = line.strip()
            if line[:4] == "$ cd":
                if line[5:] != "..":
                    dirs[stack[-1]]['dirs'][line[5:]] = {"size":0, "dirs":{}}
                    stack.append(line[5:])
                else:
                    stack.pop()
            elif line[:4] == "dir":
                dirs[stack[-1]][dirs][line[5:]] = {"size":0, "dirs":{}}
            elif line[0].isnumeric():
                dirs[stack[-1]]["size"] += int(line[0])


if __name__ == "__main__":
    main()
