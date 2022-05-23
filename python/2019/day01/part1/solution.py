import math

def calculateFuel(mass):
    return math.floor(mass / 3) - 2

def solution(path):
    totalFuel = 0
    with open(path, "r") as input:
        for line in input:
            totalFuel += calculateFuel(int(line))
    return totalFuel

realSolution = solution("input.txt")
assert realSolution == 3389778