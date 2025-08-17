func findLUSlength(a string, b string) int {
    if a == b {
        return -1
    }
    sub_a_set := map[string]struct{}{}
    sub_b_set := map[string]struct{}{}

    answer := -1

    for s, sub_set := range map[string]map[string]struct{}{
        a: sub_a_set,
        b: sub_b_set,
    } {
        for i:=0; i < len(s); i++ {
            for j:=i; j < len(s); j++ {
                sub_set[s[i:j+1]] = struct{}{}
            }
        }
    }
    a_b_xor := map[string]struct{}{}
    for k, _ := range sub_a_set {
        a_b_xor[k] = struct{}{}
    }
    for k, _ := range sub_b_set {
        if _, ok := a_b_xor[k]; ok {
            delete(a_b_xor, k)
        } else {
            a_b_xor[k] = struct{}{}
        }
    }
    for k, _ := range a_b_xor {
        answer = max(len(k), answer)
    }
    return answer
}
