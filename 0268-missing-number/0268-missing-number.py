class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        showed = {i:False for i in range(len(nums)+1)}
        for num in nums:
            showed[num] = True
        for k, v in showed.items():
            if not v:
                return k
