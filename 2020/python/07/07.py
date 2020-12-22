r = []
with open("input.txt") as f:
    for l in f:
        a = l.split(" contain")[0].split(" ")[0] + " " +  l.split(" contain")[0].split(" ")[1]
        b = l.split(" contain")[1].strip().split(", ")
        c = []
        for x in b:
            c.append(x.split(" ")[1] + " " + x.split(" ")[2])
        r.append([a,set(c)])


q = ["shiny gold"]

c = 0

history = set()
while len(q) > 0:
    a = q.pop()
    for n in r:
        if a in n[1] and n[0] not in q and n[0] not in history:
            q.insert(0, n[0])
            history.add(n[0])
            c += 1

print(c)

