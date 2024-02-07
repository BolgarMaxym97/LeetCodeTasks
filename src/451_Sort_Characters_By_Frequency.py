class Solution:
    def frequencySort(self, s: str) -> str:
        char_frequency = {}
        for char in s:
            char_frequency[char] = char_frequency[char] + 1 if char in char_frequency else 1
        char_frequency = dict(reversed(sorted(char_frequency.items(), key=lambda item: item[1])))
        result = ""

        for char in char_frequency:
            result += char * char_frequency[char]
        return result

print(Solution.frequencySort(Solution, "tree"))