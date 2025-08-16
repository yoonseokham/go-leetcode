func isUgly(n int) bool {
    isDivided := true
    for ;isDivided && n > 1;{
        isDivided = false
        for _, num := range []int{2,3,5} {
            if n%num == 0 {
                n = n/num
                isDivided = true
            }
        }
    }
    return n == 1
}