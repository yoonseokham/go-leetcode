class Solution:
    TABLE = [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
    def __init__(self):
        self.MORSE_CODE_TABLE = {}
        self.tableToAlphaMap()

    def tableToAlphaMap(self):
        for i in range(len(self.TABLE)):
            self.MORSE_CODE_TABLE[chr(97+i)] = self.TABLE[i]

    def uniqueMorseRepresentations(self, words: List[str]) -> int:
        answer_set = set()
        for word in words:
            morse_code = ''
            for c in word:
                morse_code += self.MORSE_CODE_TABLE[c]
            answer_set.add(morse_code)
        return len(answer_set)