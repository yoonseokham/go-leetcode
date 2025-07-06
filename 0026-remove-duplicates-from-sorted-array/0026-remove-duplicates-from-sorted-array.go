import (
    "fmt"
    "sort"
)

func removeDuplicates(nums []int) int {
    no_dup_set := make(map[int]bool)
    temp_array := make([]int, len(nums),len(nums))
    defer clear(no_dup_set)
    for i := 0; i < len(nums); i++ {
        temp_array[i] = 101
        no_dup_set[nums[i]] = true
    }
    i := 0
    for k, v := range no_dup_set {
        if v {
            temp_array[i] = k
            i ++
        }
    }
    sort.Ints(temp_array)
    for i := 0 ;i < len(temp_array); i++ {
        nums[i] = temp_array[i]
    }
    return len(no_dup_set)
}