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
        l, r, check = 1, n, 1
        while check != 0:
            mid = (l + r) / 2
            check = guess(mid)
            if check == 1: l = mid + 1
            else: r = mid - 1
        return int(mid)

print(Solution.guessNumber(Solution, 10))
