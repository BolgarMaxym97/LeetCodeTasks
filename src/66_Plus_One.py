from typing import List


class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        number = int(''.join(map(str, digits)))
        increased_number_str = str(number + 1)
        return list(map(int, list(increased_number_str.strip(''))))

print(Solution.plusOne(Solution, [9]))