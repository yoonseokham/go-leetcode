class Solution:
    def sumOddLengthSubarrays(self, arr: List[int]) -> int:
        answer = 0
        sum_list = [i for i in arr]
        for i in range(1, len(arr)):
            sum_list[i] = sum_list[i-1] + arr[i]
        for length in range(1,len(arr)+1,2):
            for j in range(len(arr)-length+1):
                start = j
                end = j + length - 1
                if not start:
                    answer += sum_list[end]
                else:
                    answer += (sum_list[end] -  sum_list[start-1])
        return answer
        