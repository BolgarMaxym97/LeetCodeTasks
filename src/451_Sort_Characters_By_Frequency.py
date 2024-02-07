class Solution:
    def frequencySort(self, s: str) -> str:
        char_frequency = {}
        for char in s:
            char_frequency[char] = char_frequency.get(char, 0) + 1
        char_frequency = sorted(char_frequency.items(), key=lambda item: item[1], reverse=True)
        result = ""

        for char in char_frequency:
            result += char[0] * char[1]
        return result

print(Solution.frequencySort(Solution, "tree"))