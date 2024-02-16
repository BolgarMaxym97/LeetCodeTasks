from typing import List


class Solution:
    def findLeastNumOfUniqueInts(self, arr: List[int], k: int) -> int:
        nums_frequency = dict()
        for v in arr:
            if v in nums_frequency:
                nums_frequency[v] += 1
            else:
                nums_frequency[v] = 1
        nums_frequency = sorted(nums_frequency.items(), key=lambda item: item[1])
        for _ in range(len(nums_frequency)):
            if k == 0:
                break
            if nums_frequency[0][1] <= k:
                k -= nums_frequency[0][1]
                del nums_frequency[0]
                continue
            break

        unique_nums = set(items[0] for items in nums_frequency)
        return len(unique_nums)

print(Solution.findLeastNumOfUniqueInts(Solution, [1,2,3], 3))