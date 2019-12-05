def check(num):
    code = str(num)
    substringLength2 = False
    currentSubstring = 0
    substringChar = 'X'
    isIncreasing = True
    for i in range(0, len(code)):
        if(i < len(code) - 1 and code[i] > code[i + 1]):
            isIncreasing = False 
        if(substringChar == code[i]):  
            currentSubstring += 1
        else:
            if(currentSubstring == 2):
                substringLength2 = True
            substringChar = code[i]
            currentSubstring = 1
    
    if(currentSubstring == 2):
        substringLength2 = True
    return substringLength2 and isIncreasing

def solution(lowerBound, upperBound):
    solutions = 0
    for num in range(lowerBound, upperBound):
        if check(num):
            solutions += 1
    return solutions


assert check(112233) == True
assert check(123444) == False
assert check(111122) == True
realSolution = solution(206938, 679128)
assert realSolution == 1133
print(realSolution)

