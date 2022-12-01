def findOutput(list, noun, verb):
    list[1] = noun
    list[2] = verb

    #Iterate until you find 99
    i = 0
    while list[i] != 99:
        #Get the values in the first 4 cells
        opcode = list[i]
        inputPosition1 = list[i + 1]
        inputPosition2 = list[i + 2]
        outputPosition = list[i + 3]
        
        #Do the operation
        input1 = list[inputPosition1]
        input2 = list[inputPosition2]
        output = input1 + input2 if opcode == 1 else input1 * input2

        #Save the answer and move to the next four numbers
        list[outputPosition] = output
        i += 4
    return list[0]

def solution(path, noun, verb):
    with open(path, "r") as input:
        #Convert file to a list
        dataList = [int(num) for num in input.read().split(",")]
    return findOutput(dataList, noun, verb)

solutionPart2 = solution("input.txt", 64, 29)
assert solutionPart2 == 19690720

realSolution = solution("input.txt", 12, 2)
assert realSolution == 3716293