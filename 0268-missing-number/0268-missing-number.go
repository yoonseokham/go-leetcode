func missingNumber(nums []int) int {
    showed := make(map[int]struct{})
    for _, num := range nums {
        showed[num] = struct{}{}
    }
    for i:=0; i<=len(nums);i++{
        if _, ok := showed[i]; !ok {
            return i
        }
    }
    panic("unreachable")
}
