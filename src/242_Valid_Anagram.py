class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        set1 = set(list(s.strip('')))
        for char in set1:
            if s.count(char) != t.count(char):
                return False
        return True

print(Solution.isAnagram(Solution, "aacc", "ccac"))