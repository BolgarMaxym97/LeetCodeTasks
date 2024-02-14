from typing import List


class Solution:
    def rearrangeArray(self, nums: List[int]) -> List[int]:
        result = []
        positive_nums = [n for n in nums if n > 0]
        negative_nums = [n for n in nums if n < 0]
        i = 0
        for p in positive_nums:
            result.append(p)
            if negative_nums[i]:
                result.append(negative_nums[i])
            i += 1
        return result

print(Solution.rearrangeArray(Solution, [-1, 1]))