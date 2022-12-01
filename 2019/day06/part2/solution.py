from collections import defaultdict

class Graph:
    def __init__(self):
        self.orbits = {} # String -> Int (Name -> number of orbits)
        self.edges = defaultdict(list)  # String -> List(String) (Name -> List of neighbours)

    def addEdge(self, line):
        tmp = line.split(')')
        self.edges[tmp[0]].append(tmp[1])
        self.edges[tmp[1]].append(tmp[0])
    
    def distance(self, start, end):
        visited = set()
        queue = []

        queue.append(start)
        self.orbits[start] = 0

        while(queue):
            current = queue.pop(0)

            if not current in visited:
                visited.add(current)
                for child in self.edges[current]:
                    queue.append(child)
                    childOrbits = self.orbits[current] + 1
                    self.orbits[child] = childOrbits
                    if child == end:
                        return self.orbits[end] - 2

def solution(path):
    graph = Graph()
    with open(path, "r") as input:
        lines = [line.strip('\n') for line in input.readlines()]
        for line in lines:
            graph.addEdge(line)
        return graph.distance("YOU", "SAN")


sampleAnswer1 = solution("sample1.txt")
assert sampleAnswer1 == 4

sampleAnswer2 = solution("sample2.txt")
assert sampleAnswer2 == 4

realAnswer = solution("input.txt")
assert realAnswer == 418