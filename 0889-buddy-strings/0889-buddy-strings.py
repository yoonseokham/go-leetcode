class Solution:
    def buddyStrings(self, s: str, goal: str) -> bool:
        if len(s) != len(goal): return False
        if s == goal:
            return not (len(set(s)) == len(s))
        diff_index_list = []
        for i in range(len(s)):
            if goal[i] != s[i]:
                diff_index_list.append(i)
            if len(diff_index_list) > 2: return False
        if len(diff_index_list) == 2:
            return all(
                (
                    s[diff_index_list[0]] == goal[diff_index_list[1]],
                    s[diff_index_list[1]] == goal[diff_index_list[0]],
                    )
            )
        return False
