def inputParser(path:str) -> list:
    lines = []
    with open(path) as f:
        for l in f:
            lines.append(l[:-2])
    return lines



def main():
    levels =  inputParser("input2.txt") # Matrix
    lv = 0
    i = 0
    cant_levels = len(levels)
    largo_lv = len(levels[0]) 
    print(levels[0]) 
    print(f"{cant_levels} - {largo_lv}")
    trees = 0
    while True:
        print(f"{lv} - {i}")
        print(levels[lv])
        lv= lv + 1
        if  lv >= cant_levels:
            return trees
        i += 3
        if i >= largo_lv:
            i = i - largo_lv
        if levels[lv][i] == "#":
            trees = trees + 1
            print("#")

if __name__ == "__main__" :
    print(main())
