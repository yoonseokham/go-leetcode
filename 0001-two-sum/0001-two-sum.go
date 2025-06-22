import "fmt"

func twoSum(nums []int, target int) []int {
    value_index_list_map := make(map[int][]int)
    answer := make([]int,2)
    for i:=0;i<len(nums);i++{
        value_index_list_map[nums[i]] = append(value_index_list_map[nums[i]],i)
    }
    fmt.Println(value_index_list_map)
    for i:=0;i<len(nums);i++ {
        if 2*nums[i] == target {
            if len(value_index_list_map[nums[i]]) ==2 {
                return value_index_list_map[nums[i]]
            }
        } else {
            if len(value_index_list_map[target-nums[i]])>0 {
                answer[0] = value_index_list_map[target-nums[i]][0]
                answer[1] = value_index_list_map[nums[i]][0]
            }
        }

    }
    return answer
}