from math import sqrt

class Solution:
    def isPowerOfTwo(self, n: int) -> bool:
        while True:
            if n == 0:
                return False
            if n == 1:
                return True
            if n % 2 == 0:
                n //= 2
            else:
                return False

print(Solution.isPowerOfTwo(Solution, 16))