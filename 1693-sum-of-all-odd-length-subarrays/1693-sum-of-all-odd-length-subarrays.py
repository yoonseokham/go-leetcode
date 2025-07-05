class Solution:
    def sumOddLengthSubarrays(self, arr: List[int]) -> int:
        sum_list = list(range(len(arr)))
        sum_list[0] = arr[0]
        for i in range(1, len(arr)):
            sum_list[i] = sum_list[i-1] + arr[i]
        answer = sum_list[-1] 
        for length in range(3,len(arr)+1,2):
            for j in range(len(arr)):
                start = j
                end = j + length - 1
                if end >= len(arr):
                    break
                if not start:
                    answer += sum_list[end]
                else:
                    answer += (sum_list[end] -  sum_list[start-1])
        return answer
        
# 1 5 7 12 15
# sum(1~3) = 
# sum_list[3] - sum_list[0]
# sum(2~4)
# sum_list[4] - sum_list[1]