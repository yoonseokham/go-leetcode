import (
    "fmt"
    "strconv"
)

func addDigits(num int) int {
    string_num := strconv.Itoa(num)
    if len(string_num) == 1 {
        return num
    }
    var sum_string_num int
    for i:=0;i<len(string_num);i++{
        temp, _ := strconv.Atoi(string(string_num[i]))
        sum_string_num += temp
    }
    return addDigits(sum_string_num)
}
