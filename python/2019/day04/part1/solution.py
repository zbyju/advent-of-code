def check(num):
    code = str(num)
    hasDouble = False
    isIncreasing = True
    for i in range(0, len(code)):
        if(i < len(code) - 1):
            if(code[i] == code[i + 1]):
                hasDouble = True
            elif(code[i] > code[i + 1]):
                isIncreasing = False
        
    return hasDouble and isIncreasing

def solution(lowerBound, upperBound):
    solutions = 0
    for num in range(lowerBound, upperBound):
        if check(num):
            solutions += 1
    return solutions


assert check(111111) == True
assert check(223450) == False
assert check(123789) == False
realSolution = solution(206938, 679128)
assert realSolution == 1653
print(realSolution)

