func countWords(words1 []string, words2 []string) int {
    w1 := make(map[string]int)
    w2 := make(map[string]int)
    answer := 0
    for _, word1 := range(words1) {
        w1[word1] ++
    }
    for _, word2 := range(words2) {
        w2[word2] ++
    }
    for key,_ := range(w1) {
        if w1[key] == 1 && w2[key] == 1 {
            answer += 1
        }
    }
    return answer
}