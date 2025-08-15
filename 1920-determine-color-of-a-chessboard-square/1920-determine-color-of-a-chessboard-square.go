func squareIsWhite(coordinates string) bool {
    x, y := int(coordinates[0])%2, int(coordinates[1])%2
    if x == 1 && y == 1 {
       return false
    }
    if x != 1 && y != 1 {
        return false
    }
    return true
}
