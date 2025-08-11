class Solution:
    MONTH_MAP = {
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
    
    def reformatDate(self, date: str) -> str:
        day, month, year = list(date.split())
        day = '0' + day[:-2]
        month = '0' + self.MONTH_MAP[month]
        return f"{year}-{month[-2:]}-{day[-2:]}"
