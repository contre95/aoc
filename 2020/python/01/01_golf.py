lines = []
with open('./intput.txt') as f:
    for l in f:
        lines.append(l)

r = [int(i)*int(j) for i in lines for j in lines if int(i)+int(j) ==2020 and i!=j]
print(r)
        
