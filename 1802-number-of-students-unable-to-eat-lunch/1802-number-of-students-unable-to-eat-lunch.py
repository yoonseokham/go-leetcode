import collections

class Solution:
    def countStudents(self, students: List[int], sandwiches: List[int]) -> int:
        students_dq = collections.deque(students)
        sandwiches_dq = collections.deque(sandwiches)
        loop = 0
        while sandwiches_dq:
            if students_dq[0] == sandwiches_dq[0]:
                students_dq.popleft()
                sandwiches_dq.popleft()
                loop = 0
            else:
                temp = students_dq[0]
                students_dq.popleft()
                students_dq.append(temp)
                loop += 1
                if loop >= len(students_dq):
                    break
        return len(students_dq)
