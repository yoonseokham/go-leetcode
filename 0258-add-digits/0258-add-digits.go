import (
    "strconv"
)

func addDigits(num int) int {
    if num/10 == 0 {
        return num
    }
    string_num := strconv.Itoa(num)
    var sum_string_num int
    for i:=0;i<len(string_num);i++{
        temp, _ := strconv.Atoi(string(string_num[i]))
        sum_string_num += temp
    }
    return addDigits(sum_string_num)
}
