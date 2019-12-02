def findOutput(data, noun, verb):
    data[1] = noun
    data[2] = verb

    #Iterate until you find 99
    i = 0
    while data[i] != 99:
        #Get the values in the first 4 cells
        opcode = data[i]
        inputPosition1 = data[i + 1]
        inputPosition2 = data[i + 2]
        outputPosition = data[i + 3]
        
        #Do the operation
        input1 = data[inputPosition1]
        input2 = data[inputPosition2]
        output = input1 + input2 if opcode == 1 else input1 * input2

        #Save the answer and move to the next four numbers
        data[outputPosition] = output
        i += 4
    return data[0]

def solution(path, expectedValue):
    with open(path, "r") as input:
        #Convert file to a list
        dataList = [int(num) for num in input.read().split(",")]
        
    for noun in range(0, 99):
        for verb in range(0, 99):
            #Find output BUT pass the list as a copy (not by reference)
            if(findOutput(dataList[:], noun, verb) == expectedValue):
                return 100 * noun + verb

realSolution = solution("input.txt", 19690720)
assert realSolution == 6429