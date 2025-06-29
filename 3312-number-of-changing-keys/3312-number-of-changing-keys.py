class Solution:
    def countKeyChanges(self, s: str) -> int:
        previous = ord(s[0])
        change = 0
        for i in range(1,len(s)):
            current = ord(s[i])
            if not (previous == current or abs(current - previous) == 32):
                change += 1
            previous = current
        return change

        