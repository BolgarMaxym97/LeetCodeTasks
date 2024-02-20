# The guess API is already defined for you.
# @param num, your guess
# @return -1 if num is higher than the picked number
#          1 if num is lower than the picked number
#          otherwise return 0
# def guess(num: int) -> int:

def guess(num: int) -> int:
    if num == 8:
        return 0
    if num > 8:
        return -1
    return 1
class Solution:
    def guessNumber(self, n: int) -> int:
        if guess(n) == 0:
            return n
        while True:
            mid = n // 2
            if guess(mid) == 0:
                return mid
            elif guess(mid) == 1:
                n = int((mid + n))
            else:
                n = int((n - mid))
        return 0

print(Solution.guessNumber(Solution, 10))
