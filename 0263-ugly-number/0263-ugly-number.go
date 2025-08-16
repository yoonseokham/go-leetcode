func isUgly(n int) bool {
   for ;n>1;{
    isDivided := false
    for _, num := range []int{2,3,5}{
        if n%num == 0 {
            n = n/num
            isDivided = true
            }
        }
    if !isDivided {
        break
        }
    }
    if n != 1 {
        return false
    }
    return true
}