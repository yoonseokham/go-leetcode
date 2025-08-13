class Solution:
    def getRemovedNums(self, nums, removed_start, removed_end):
        new_nums = nums[:removed_start] + nums[removed_end:]
        return new_nums
    
    def isIncrease(self, new_nums) -> bool:
        temp = 0
        for num in new_nums:
            if num <= temp: return False
            temp = num
        print(new_nums)
        return True

    def incremovableSubarrayCount(self, nums: List[int]) -> int:
        answer = 0
        for i in range(len(nums)+1):
            for j in range(i+1,len(nums)+1):
               if self.isIncrease(self.getRemovedNums(nums, i, j)):
                answer += 1 
        return answer
