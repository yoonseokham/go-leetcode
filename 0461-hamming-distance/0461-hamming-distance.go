func hammingDistance(x int, y int) int {
    answer := 0
    x_b := getBinary(x)
    y_b := getBinary(y)
    for i:= 0;i < 32; i++ {
        if x_b[i] != y_b[i] {
            answer++
        }
    }
    return answer
}

func getBinary(n int) [32]int {
    var binary [32]int
    for i:=0;n>0;i++ {
        binary[i] = n%2
        n = int(n/2)
    }
    return binary
}
