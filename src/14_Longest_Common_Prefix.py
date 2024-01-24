from typing import List

class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        min_string = min(strs, key=len)
        min_string_len = len(min_string)
        for i in range(min_string_len)[::-1]:
            sub_string = min_string[:i+1]
            if sum(sub_string in s[:i+1] for s in strs) == len(strs):
                return sub_string
        return ""

print(Solution.longestCommonPrefix(Solution, strs=["dog","racecar","car"]))