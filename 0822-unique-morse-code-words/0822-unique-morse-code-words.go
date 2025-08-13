import "fmt"

func uniqueMorseRepresentations(words []string) int {
    TABLE := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
    MORSE_MAP := make(map[string]string)
    answer_set := make(map[string]bool)
    var answer int 
    for i:=0;i<len(TABLE);i++{
         s := fmt.Sprintf("%c", 97+i)
         MORSE_MAP[s] = TABLE[i]
    }
    for i:=0;i<len(words);i++ {
        morse_code := ""
        for j:=0;j<len(words[i]);j++{
            morse_code += MORSE_MAP[string(words[i][j])]
        }
        answer_set[morse_code] = true
    }
    for k := range answer_set {
        if answer_set[k] {
            answer ++
        }
    }
    return answer
}
