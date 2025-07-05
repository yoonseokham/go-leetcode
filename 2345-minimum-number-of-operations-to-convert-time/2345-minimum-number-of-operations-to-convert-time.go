import (
    	"strings"
        "strconv"
    )

func convertTime(current string, correct string) int {
    minDiff := getDiffMinTime(current, correct)
    div_list := [4]int{60, 15, 5, 1}
    answer := 0
    var operation int
    for i:=0;i<len(div_list);i++{
        operation = minDiff/div_list[i]
        answer += operation
        minDiff -= operation*div_list[i]
    }
    return answer
}

func getDiffMinTime(current string, correct string) int {
    current_parts := strings.Split(current, ":")
    current_hour, _ := strconv.Atoi(current_parts[0])
    current_min, _ := strconv.Atoi(current_parts[1])
    correct_parts := strings.Split(correct, ":")
    correct_hour, _ := strconv.Atoi(correct_parts[0])
    correct_min, _ := strconv.Atoi(correct_parts[1])
    return 60*(correct_hour-current_hour) + correct_min - current_min
}