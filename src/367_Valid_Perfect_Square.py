class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        return (num ** 0.5).is_integer()

print(Solution.isPerfectSquare(Solution, 16))
