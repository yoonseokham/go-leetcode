import "fmt"
import "strings"

func reformatDate(date string) string {
    MONTH_MAP := map[string]string {
        "Jan": "1", 
        "Feb": "2",
        "Mar": "3", 
        "Apr": "4", 
        "May": "5", 
        "Jun": "6", 
        "Jul": "7", 
        "Aug": "8", 
        "Sep": "9", 
        "Oct": "10", 
        "Nov": "11", 
        "Dec": "12",
    }
    date_str := strings.Split(date, " ")
    day := "0" + date_str[0][:len(date_str[0])-2]
    month := "0" + MONTH_MAP[date_str[1]]
    year := date_str[2]
    return fmt.Sprintf("%s-%s-%s", year, month[len(month)-2:], day[len(day)-2:])
}
