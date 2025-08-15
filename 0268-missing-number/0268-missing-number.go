func missingNumber(nums []int) int {
    showed := make(map[int]bool)
    for _, num := range nums {
        showed[num] = true
    }
    for i:=0; i<=len(nums);i++{
        if !showed[i] {
            return i
        }
    }
    panic("unreachable")
}