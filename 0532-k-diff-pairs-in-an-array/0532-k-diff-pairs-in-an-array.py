import collections

class Solution(object):
    def findPairs(self, nums, k):
        kv_dict = collections.Counter(nums)
        answer = []
        for num in nums:
            kv_dict[num] = kv_dict[num] - 1
            if kv_dict[num+k]:
                answer.append((num,num+k))
            kv_dict[num] = kv_dict[num] + 1
        print(answer)
        return len(set(answer))

        
        