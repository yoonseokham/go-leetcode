import collections

class Solution:
    def maxFreqSum(self, s: str) -> int:
        counted = collections.Counter(s)
        vowel = set(('a', 'e', 'i', 'o', 'u'))
        answer = 0
        v_max = 0
        c_max = 0
        for k, v in counted.items():
            if k in vowel:
                v_max = max(v_max, v)
            else:
                c_max = max(c_max, v)
        return v_max + c_max
