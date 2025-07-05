func minimumOperations(grid [][]int) int {
    operation := 0
    var diff int
    for i := 0; i<len(grid[0]); i++ {
        for j := 1; j<len(grid); j++ {
            if grid[j][i] <= grid[j-1][i] {
                diff = grid[j-1][i] - grid[j][i]
                operation += diff + 1
                grid[j][i] = grid[j-1][i] + 1
            }
        }
    }
    return operation
}
