from typing import List


class Solution:
    def flipAndInvertImage(self, image: List[List[int]]) -> List[List[int]]:
        return [[1 if item == 0 else 0 for item in reversed(items)] for items in image]

print(Solution.flipAndInvertImage(Solution, [[1,1,0],[1,0,1],[0,0,0]]))