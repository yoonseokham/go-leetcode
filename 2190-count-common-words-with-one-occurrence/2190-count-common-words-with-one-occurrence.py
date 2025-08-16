import collections

class Solution:
    def countWords(self, words1: List[str], words2: List[str]) -> int:
        w1 = collections.Counter(words1)
        w2 = collections.Counter(words2)
        answer = 0 
        for k in w1.keys():
            if w1[k] == 1 and w2[k] == 1:
                answer += 1
        return answer

