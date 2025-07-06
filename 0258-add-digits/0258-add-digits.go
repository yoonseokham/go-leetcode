import (
    "fmt"
    "strconv"
)

func addDigits(num int) int {
    string_num := strconv.Itoa(num)
    var sum_string_num int
    if len(string_num) == 1 {
        return num
    }
    for i:=0;i<len(string_num);i++{
        temp, _ := strconv.Atoi(string(string_num[i]))
        sum_string_num += temp
    }
    return addDigits(sum_string_num)
}