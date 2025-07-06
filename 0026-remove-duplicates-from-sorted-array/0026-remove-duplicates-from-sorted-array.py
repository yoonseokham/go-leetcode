class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        no_dup = sorted(list(set(nums)))
        for i in range(len(no_dup)):
            nums[i] = no_dup[i]
        answer = len(no_dup)
        del no_dup
        return answer