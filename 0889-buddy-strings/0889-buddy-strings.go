func buddyStrings(s string, goal string) bool {
    if len(s) != len(goal) {return false}
    if s == goal {
        s_dict := make(map[byte]bool)
        for i:=0;i<len(s);i++{
            s_dict[s[i]] = true
        }
        return len(s_dict) != len(s)
    }
    diffIndexs := make([]int,0,2)
    for i:= 0;i<len(s);i++{
        if s[i] != goal[i] {
            diffIndexs = append(diffIndexs,i)
        }
        if len(diffIndexs) > 2 {return false}
    }
    if len(diffIndexs) == 1 {return false}
    isOkay := (s[diffIndexs[0]] == goal[diffIndexs[1]] && s[diffIndexs[1]] == goal[diffIndexs[0]])
    if !isOkay {return false}
    return true
}
