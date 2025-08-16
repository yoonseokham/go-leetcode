class Solution:
    def isUgly(self, n: int) -> bool:
        divided = True
        while n>1 and divided:
            divided = False
            for num in (2, 3, 5):
                if not n%num:
                    n = n//num
                    divided = True
        return n == 1

        