import "math"

func countKeyChanges(s string) int {
    previous := int(s[0])
    change_count :=0
    var current int
    for i:=1;i<len(s);i++ {
        current = int(s[i])
        if previous != current && int(math.Abs(float64(previous- current))) != 32 {
            change_count += 1
        }
        previous = current
    }
    return change_count
}