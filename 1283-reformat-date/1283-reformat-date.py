class Solution:
    MONTH_MAP = {
        "Jan": "01",
        "Feb": "02",
        "Mar": "03",
        "Apr": "04",
        "May": "05",
        "Jun": "06",
        "Jul": "07",
        "Aug": "08",
        "Sep": "09",
        "Oct": "10",
        "Nov": "11",
        "Dec": "12",
        }
    def addZero(self, n: str) -> str:
        return '0' + n if len(n) == 1 else n
    
    def reformatDate(self, date: str) -> str:
        day, month, year = list(date.split())
        day = day[:-2]
        month = self.MONTH_MAP[month]
        return f"{year}-{month}-{self.addZero(day)}"
        