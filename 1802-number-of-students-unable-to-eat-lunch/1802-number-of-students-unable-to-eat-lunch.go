func countStudents(students []int, sandwiches []int) int {
    var temp int
    var loop int
    for ;len(sandwiches)>0&&loop<=len(sandwiches); {
        if students[0] == sandwiches[0] {
            students = students[1:]
            sandwiches = sandwiches[1:]
            loop = 0
        } else {
           temp = students[0]
           students = students[1:]
           students = append(students, temp)
           loop += 1
        }
    }
    return len(students)
}
