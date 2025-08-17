import collections

class Solution:
    def findLUSlength(self, a: str, b: str) -> int:
        sub_a_set = set()
        sub_b_set = set()
        answer = -1
        for s, sub_set in zip(
            (a, b),
            (sub_a_set, sub_b_set),
        ):
            for i in range(len(s)):
                for j in range(i, len(s)):
                    sub_set.add(s[i:j+1])
        for s in sub_a_set ^ sub_b_set:
            answer = max(len(s), answer)
        return answer
