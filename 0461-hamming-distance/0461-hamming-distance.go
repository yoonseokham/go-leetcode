import "fmt"
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

func getBinary(n int) []int {
    binary := make([]int, 32, 32)
    i := 0
    for {
        binary[i] = n%2
        i++
        n = int(n/2)
        if n == 0 {
            break
        }
    }
    return binary
}
