import math

def calculateFuel(mass):
    return math.floor(mass / 3) - 2

input = open("./input.txt", "r")
totalFuel = 0
for line in input:
    totalFuel += calculateFuel(int(line))
print(totalFuel)