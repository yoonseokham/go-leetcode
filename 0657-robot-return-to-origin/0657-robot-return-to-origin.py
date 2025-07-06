import collections

class Solution:
    def judgeCircle(self, moves: str) -> bool:
        move_op = collections.defaultdict(lambda:0)
        for move in moves:
            move_op[move] += 1
        return (move_op["L"] == move_op["R"]) and (move_op["U"] == move_op["D"])
        