/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func lowerBound(n int) int {
    start := 1
    end := n
    for;start<end;{
        mid := int((start + end)/2)
        if isBadVersion(mid) {
            end = mid
        } else {
            start = mid
        }
        if end - start == 1 {
            break
        }
    }
    return end
}
func firstBadVersion(n int) int {
    if isBadVersion(1) {
        return 1
    }
    return lowerBound(n)
}