import "math"
import "fmt"

func countOperations(num1 int, num2 int) int {
    count := 0
    var recursive func(int, int)
    recursive = func (num1 int, num2 int) {
        if num1 == 0 || num2 == 0 {
            return
        }
        temp_num2 := int(math.Min(float64(num1),float64(num2)))
        temp_num1 := int(math.Max(float64(num1),float64(num2))) - temp_num2
        fmt.Println(temp_num1,temp_num2)
        count ++
        recursive(temp_num1,temp_num2)
    }
    recursive(num1, num2)
    return count
}