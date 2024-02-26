from typing import List


class Solution:
    def findWordsContaining(self, words: List[str], x: str) -> List[int]:
        result, i = [], 0

        for word in words:
            if x in word:
                result.append(i)
            i += 1

        return result

print(Solution.findWordsContaining(Solution, words = ["leet","code"], x = "e"))
