from typing import List


class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        nums = sorted(nums)
        i = 0
        for num in nums:
            if i == 0 and num == 1:
                return 0
            if nums[i - 1] != num - 1 and i != 0:
                return num - 1
            i += 1
        return nums[len(nums) - 1] + 1

print(Solution.missingNumber(Solution, [9,6,4,2,3,5,7,0,1]))