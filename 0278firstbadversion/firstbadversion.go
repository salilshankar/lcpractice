func firstBadVersion(n int) int {

	if !isBadVersion(n-1) && isBadVersion(n) {
		return n
	}

	start, end := 1, n

	for start < end {
		mid := (start + end) / 2
		if isBadVersion(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}