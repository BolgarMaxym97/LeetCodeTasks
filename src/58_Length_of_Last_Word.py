class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        return len(s.split()[-1])

print(Solution.lengthOfLastWord(Solution, "   fly me   to   the moon  "))
