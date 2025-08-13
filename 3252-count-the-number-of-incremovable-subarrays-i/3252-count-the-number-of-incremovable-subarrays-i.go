func getRemovedNums(nums []int, start int, end int) []int {
    copied := make([]int, len(nums))
    copy(copied, nums)
    new_nums := append(copied[:start], copied[end:]...)
    return new_nums
}

func isIncreased(nums []int) bool {
    temp := 0
    for i:=0;i<len(nums);i++{
        if temp >= nums[i] {
            return false
        }
        temp = nums[i]
    }
    return true
}

func incremovableSubarrayCount(nums []int) int {
    var answer int
    for i:=0;i<=len(nums);i++{
        for j:=i+1;j<=len(nums);j++{
            if isIncreased(getRemovedNums(nums, i, j)) {
                answer += 1
            }     
        }
    }
    return answer
}
