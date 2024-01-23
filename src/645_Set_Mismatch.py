from typing import List

class Solution:
    def findErrorNums(self, nums: List[int]) -> List[int]:
        correctArr = list(range(1, len(nums) + 1))
        print(correctArr)
        result = []
        for x in set(nums):
            if nums.count(x) > 1:
                result.append(x)
                result.append(list(set(correctArr) - set(nums))[0])
        return result

print(Solution.findErrorNums(Solution, [2,2]));