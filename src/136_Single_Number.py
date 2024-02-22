from typing import List


class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        nums_frequency = dict()
        for v in nums:
            if v in nums_frequency:
                if nums_frequency[v] > 1:
                    continue
                nums_frequency[v] += 1
            else:
                nums_frequency[v] = 1
        return list(nums_frequency.keys())[list(nums_frequency.values()).index(1)]

print(Solution.singleNumber(Solution, [4,1,2,1,2]))
