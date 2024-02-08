from math import sqrt

class Solution:


    def numSquares(self, n: int) -> int:
        perfectSquares = [i for i in range(1, n) if sqrt(i).is_integer()]

        result = []
        while n > 0:
            n -= max(perfectSquares)
            if n == 0:
                result.append(max(perfectSquares))
                break
            if n > 0:
                result.append(max(perfectSquares))
            if n < 0:
                n += max(perfectSquares)
                perfectSquares.pop(perfectSquares.index(max(perfectSquares)))

        return len(result)


print(Solution.numSquares(Solution, 12))