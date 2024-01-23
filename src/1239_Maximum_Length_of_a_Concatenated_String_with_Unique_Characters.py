from typing import List

class Solution:
    def maxLength(self, arr: List[str]) -> int:
        if len(arr) == 1:
            return len(arr[0])

        firstLongest = max(arr, key=len)
        arr.remove(firstLongest)
        secondLongest = max(arr, key=len)
        return len(firstLongest + secondLongest)

print(Solution.maxLength(Solution, ["cha","r","act","ers"]))