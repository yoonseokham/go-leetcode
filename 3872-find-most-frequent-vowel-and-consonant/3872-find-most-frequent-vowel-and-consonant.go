func maxFreqSum(s string) int {
    v_max := 0
    c_max := 0
    vowels := map[byte]bool{
        'a': true, 
        'e': true,
        'i': true,
        'o': true,
        'u':true,
    }
    counter := make(map[byte]int)
    for i:=0 ;i<len(s);i++{
        counter[s[i]]++
    }
    for k,v := range counter {
        if vowels[k] {
            v_max = max(v, v_max)
        } else {
            c_max = max(v, c_max)
        }
    }
    return c_max + v_max
}
