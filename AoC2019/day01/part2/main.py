import math

def calculateFuel(mass):
    return max(math.floor(mass / 3) - 2, 0)

input = open("./input.txt", "r")
totalFuel = 0
for line in input:
    mass = int(line)
    while(mass > 0):
        fuel = calculateFuel(mass)
        totalFuel += fuel
        mass = fuel
print(totalFuel)