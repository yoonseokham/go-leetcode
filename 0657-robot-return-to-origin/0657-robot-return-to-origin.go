func judgeCircle(moves string) bool {
    move_op := make(map[string]int)
    for i:=0; i<len(moves); i++{
        move_op[string(moves[i])] += 1
    }
    return (move_op["U"] == move_op["D"]) && (move_op["L"] == move_op["R"])
}