import sys

a = { "U": [0,1], "L": [-1,0], "R": [1,0], "D": [0,-1] }

with open(f"./{sys.argv[1]}") as f:
    lines = f.read().strip().split("\n")
    ih, jh, it, jt  = 0,0,0,0
    positions = set()
    for line in lines:
        direc, module = line.split(" ")
        for _ in range(0,int(module)):
            ih += a[direc][0]
            jh += a[direc][1]
            touching = abs(jt-jh) <= 1 and abs(it-ih) <= 1
            if not touching:
                if it!=ih:
                    it += (ih - it) / abs(ih-it)
                if  jt!=jh:
                    jt += (jh - jt) / abs(jh-jt)
            positions.add(f"{int(it)}-{int(jt)}")

print(len(positions))
