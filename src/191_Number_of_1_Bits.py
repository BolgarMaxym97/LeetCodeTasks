class Solution:
    def hammingWeight(self, n: int) -> int:
      return (bin(n)[2:]).count('1')

print(Solution.hammingWeight(Solution, 11111111111111111111111111111101))