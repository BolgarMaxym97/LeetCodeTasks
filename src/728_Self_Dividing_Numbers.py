from typing import List


class Solution:
    def selfDividingNumbers(self, left: int, right: int) -> List[int]:
        result = []
        for i in range(left, right + 1):
            isSelfDividing = True
            for char in str(i):
                if int(char) == 0 or i % int(char) != 0:
                    isSelfDividing = False
                    break
            if isSelfDividing:
                result.append(i)
        return result

print(Solution.selfDividingNumbers(Solution, 1, 22))