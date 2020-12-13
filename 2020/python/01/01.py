from itertools import combinations
from math import prod

def read_input_to_list(file_path):
    reports = []
    with open(file_path) as f:
        for line in f.readlines():
            reports.append(int(line[:-1]))
    return reports


"""
Encontrar todos los pares de números que sumen 2020 y devolver su producto
"""
def v1():
#a + b == 2020
    numeros = list(set(read_input_to_list("input.txt")))
    productos = []
    count = 0
    for i in range(0,len(numeros)):
        for j in range(0,len(numeros)):
            count+=1
            if numeros[i] + numeros[j] == 2020:
                productos.append(numeros[i]*numeros[j])
    return f"V1: {count} -- {len(numeros)}"

def v2():
#a + b == 2020
    numeros = list(set(read_input_to_list("input.txt")))
    productos = []
    count = 0
    for i in range(0,len(numeros)):
        for j in range(i+1,len(numeros)):
            count+=1
            if numeros[i] + numeros[j] == 2020:
                productos.append(numeros[i]*numeros[j])
    return f"V2: {count} -- {len(numeros)}"

def v3():
#a + b == 2020
    count = 0
    productos = []
    numeros = set(read_input_to_list("input.txt"))
    for n in numeros:
        count+=1
        if 2020-n in numeros:
            productos.append(n*(2020-n))
    return f"V3: {count} -- {len(numeros)}"

def v1p2():
    lines = read_input_to_list("input.txt")
    r = [int(i)*int(j)*int(k) for i in lines for j in lines for k in lines if int(i)+int(j)+int(k) == 2020]
    return r

def v2p2():
    lines = read_input_to_list("input.txt")
    for c in combinations(lines, 3):
        if sum(c) ==  2020:
            return prod(c)


print(v3())
print(v2())
print(v1())
print(v1p2())
print(v2p2())




