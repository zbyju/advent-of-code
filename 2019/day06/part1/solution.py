from collections import defaultdict

class Graph:
    def __init__(self):
        self.orbits = {} # String -> Int (Name -> number of orbits)
        self.edges = defaultdict(list)  # String -> List(String) (Name -> List of neighbours)

    def addEdge(self, line):
        tmp = line.split(')')
        self.edges[tmp[0]].append(tmp[1])
    
    def checkSum(self, start):
        visited = set()
        queue = [] 
        checksum = 0

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
                    checksum += childOrbits

        return checksum

def solution(path):
    graph = Graph()
    with open(path, "r") as input:
        lines = [line.strip('\n') for line in input.readlines()]
        for line in lines:
            graph.addEdge(line)
        return graph.checkSum("COM")


sampleAnswer1 = solution("sample1.txt")
assert sampleAnswer1 == 42

sampleAnswer2 = solution("sample2.txt")
assert sampleAnswer2 == 42

realAnswer = solution("input.txt")
assert realAnswer == 241064