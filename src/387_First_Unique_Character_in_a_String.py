class Solution:
    def firstUniqChar(self, s: str) -> int:
        index, i = -1, 0
        for char in s:
            if s.count(char) == 1:
                index = i
                break
            i += 1
        return index

print(Solution.firstUniqChar(Solution, "aabb"))