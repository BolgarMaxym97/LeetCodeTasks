class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        current, max = "", 0

        for char in s:
            if char not in current:
                current += char
                max = len(current) if len(current) > max else max
            else:
                current = current[current.index(char) + 1:] + char
        return max


print(Solution.lengthOfLongestSubstring(Solution, "dvdf"))
