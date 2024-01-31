class Solution:
    def addBinary(self, a: str, b: str) -> str:
        binSum = bin(int(a, 2) + int(b, 2))
        return str(binSum).replace("0b", "")

print(Solution.addBinary(Solution, "1010", "1011"))