class Solution:
    def addDigitsRecursive(self, num: int) -> int:
        if len(str(num)) == 1:
            return num
        sum_nums = sum(int(i) for i in list(str(num).strip('')))
        return self.addDigits(self, sum_nums)

    def addDigitsIteration(self, num: int) -> int:
        sum_nums = sum(int(i) for i in list(str(num).strip('')));
        while len(str(sum_nums)) > 1:
            sum_nums = sum(int(i) for i in list(str(sum_nums).strip('')))
        return sum_nums

print(Solution.addDigitsIteration(Solution, 38))
