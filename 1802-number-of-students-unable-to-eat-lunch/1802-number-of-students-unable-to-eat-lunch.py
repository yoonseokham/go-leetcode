import collections

class Solution:
    def countStudents(self, students: List[int], sandwiches: List[int]) -> int:
        students_dq = collections.deque(students)
        sandwiches_dq = collections.deque(sandwiches)
        preference = collections.Counter(students_dq)
        while sandwiches_dq and preference[sandwiches_dq[0]]>0:
            if students_dq[0] == sandwiches_dq[0]:
                preference[sandwiches_dq[0]] -= 1
                students_dq.popleft()
                sandwiches_dq.popleft()
            else:
                students_dq.rotate(-1)
        return len(students_dq)
