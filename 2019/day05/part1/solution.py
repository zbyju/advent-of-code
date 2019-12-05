def findOutput(data, noun, verb):
    data[1] = noun
    data[2] = verb

    #Iterate until you find 99
    i = 0
    while data[i] != 99:
        opcode = data[i]
        inputPosition1 = data[i + 1]
        inputPosition2 = data[i + 2]
        outputPosition = data[i + 3]
        
        input1 = data[inputPosition1]
        input2 = data[inputPosition2]
        output = input1 + input2 if opcode == 1 else input1 * input2

        data[outputPosition] = output
        i += 4
    return data[0]

def solution(path, expectedValue):
    with open(path, "r") as input:
        dataList = [int(num) for num in input.read().split(",")]
        
    for noun in range(0, 99):
        for verb in range(0, 99):
            if(findOutput(dataList[:], noun, verb) == expectedValue):
                return 100 * noun + verb

