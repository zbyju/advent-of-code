class Point:
    def __init__(self, x, y):
        self.x = x
        self.y = y

class Edge:
    def __init__(self, start, end):
        self.start = start
        self.end = end

class Graph:
    def __init__(self):
        edges = {}
    def addEdge(self, start, direction, length):
        if direction == 'U':
            edge = Edge(start, Point(start.x, start.y + length))
        elif direction == 'D':
            edge = Edge(start, Point(start.x, start.y - length))
        elif direction == 'R':
            edge = Edge(start, Point(start.x + length, start.y))
        elif direction == 'L':
            edge = Edge(start, Point(start.x - length, start.y))
        self.edges.add(edge)

def addEdges(wire, commands):
    instructions = commands.split(',')
    currentPosition = Point(0, 0)
    for instruction in instructions:
        wire.addEdge(currentPosition, instruction[0], int(instruction[1:]))

def solution(path):
    with open('input.txt', "r") as commands:
        commands1, commands2 = commands.read().splitlines()

    wire1 = Graph()
    wire2 = Graph()

    addEdges(wire1, commands1)
    addEdges(wire2, commands2)
    return 0


realSolution = solution("input.txt")
'''
sampleSolution1 = solution("sample1.txt")
print(sampleSolution1)
assert sampleSolution1 == 159
sampleSolution2 = solution("sample2.txt")
print(sampleSolution2)
assert sampleSolution2 == 135
print(realSolution)
'''