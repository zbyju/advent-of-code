import math

def calculateFuel(mass):
    return max(math.floor(mass / 3) - 2, 0)

def solution(path):
    totalFuel = 0
    with open(path, "r") as input:
        for line in input:
            mass = int(line)
            while(mass > 0):
                fuel = calculateFuel(mass)
                totalFuel += fuel
                mass = fuel
    return totalFuel

realSolution = solution("input.txt")
assert realSolution == 5081802