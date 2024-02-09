from typing import List


class Solution:
    def largeGroupPositions(self, s: str) -> List[List[int]]:
        result, current_char, i = [], [], 0

        for char in s:
            if char not in current_char:
                if len(current_char) and current_char[1] >= 3:
                    result.append([current_char[2], i - 1])
                current_char = [char, 1, i]
                i += 1
                continue
            current_char[1] += 1
            i += 1
        return result if len(result) > 0 else [[0, len(s) - 1]] if current_char[1] >= 3 else []

print(Solution.largeGroupPositions(Solution, 'babaaaabbb'))