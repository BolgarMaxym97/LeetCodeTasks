class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False

        reversed = int(str(x)[::-1])
        return reversed == x

print(Solution.isPalindrome(Solution, 121))