func isPalindromic(word string) bool {
    for i,j := 0,len(word)-1 ;i<j;i,j=i+1,j-1 {
        if word[i] != word[j] {
            return false
        }
    }
    return true
}

func firstPalindrome(words []string) string {
    for i:=0;i<len(words);i++{
        if isPalindromic(words[i]) {
            return words[i]
        }
    }
    return ""
}