class Solution:
    def addDigits(self, num: int) -> int:
        string_num = str(num)
        if len(string_num) == 1:
            return num
        return self.addDigits(sum(int(i) for i in string_num))
