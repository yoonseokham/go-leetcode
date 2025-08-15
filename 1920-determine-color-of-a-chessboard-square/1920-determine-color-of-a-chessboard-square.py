class Solution:
    def squareIsWhite(self, coordinates: str) -> bool:
        x, y = ord(coordinates[0])%2, int(coordinates[1])%2
        if x and y:
            return False
        if not x and not y:
            return False
        return True
