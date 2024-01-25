from typing import List


class Solution:
    def maxFrequencyElements(self, nums: List[int]) -> int:
        nums_frequency = {v: nums.count(v) for v in nums}
        max_frequency = max(nums_frequency.values())
        nums_with_max_frequency = [k for k, v in nums_frequency.items() if v == max_frequency]
        return sum(nums.count(i) for i in nums_with_max_frequency)


print(Solution.maxFrequencyElements(Solution, [1, 2, 3, 4, 5]))
