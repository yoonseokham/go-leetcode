# The isBadVersion API is already defined for you.
# def isBadVersion(version: int) -> bool:

class Solution:
    def minBadSearcher(self,n):
        start = 1
        end = n
        while start < end:
            mid = (start + end) // 2
            if isBadVersion(mid):
                end = mid
            else:
                start = mid
            if end - start == 1:
                return end

    def firstBadVersion(self, n: int) -> int:
        if isBadVersion(1): return 1
        return self.minBadSearcher(n)
        