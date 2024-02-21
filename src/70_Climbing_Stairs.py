class Solution:
    def climbStairs(self, n: int) -> int:
        if n <= 2:
            return n

        prev, current = 1, 2
        for i in range(3, n + 1):
            prev, current = current, prev + current

        return current


print(Solution.climbStairs(Solution, 3))
