from typing import List

class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        majority_count = len(nums) / 2
        nums_set = set(nums)
        print(nums_set)
        for num in nums_set:
            if nums.count(num) >= majority_count:
                return num
        return 0

print(Solution.majorityElement(Solution, [2,2,1,1,1,2,2]))