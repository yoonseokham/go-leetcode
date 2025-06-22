type Pair struct {
    X, Y int
}

func findPairs(nums []int, k int) int {
    kv_map := make(map[int]int)
    seen := make(map[Pair]bool)
    count := 0
    for i:=0;i<len(nums);i++{
        kv_map[nums[i]] += 1
    }
    for i:=0;i<len(nums);i++{
        kv_map[nums[i]] -= 1
        if kv_map[nums[i]+k] > 0 {
            if !seen[Pair{nums[i], nums[i]+k}]{
                count += 1
                seen[Pair{nums[i], nums[i]+k}] = true
            }
        }
        kv_map[nums[i]] += 1
    }
    return count
}