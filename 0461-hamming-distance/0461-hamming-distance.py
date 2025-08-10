class Solution:
    def getBinary(self, n :int) -> list[int]:
        binary = [0 for i in range(32)]
        i = 0
        while n > 0:
            binary[i] = n%2
            n = n//2
            i += 1
        return binary

    def hammingDistance(self, x: int, y: int) -> int:
        answer = 0
        for i,j in zip(self.getBinary(x), self.getBinary(y)):
            if i != j:
                answer += 1
        return answer
