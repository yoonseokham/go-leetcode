class Solution:
    def convertTime(self, current: str, correct: str) -> int:
        answer = 0
        diff_min = Solution.getDiffMin(current, correct)
        for i in [60, 15, 5, 1]:
            operation = diff_min//i
            answer += operation
            diff_min -= i*operation
        return answer

    @staticmethod
    def getDiffMin(current: str, correct: str) -> int:
        current_hour, current_min = [int(i) for i in current.split(":")]
        correct_hour, correct_min = [int(i) for i in correct.split(":")]
        return 60*(correct_hour - current_hour) + correct_min - current_min