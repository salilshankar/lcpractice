func removeCoveredIntervals(intervals [][]int) int {
	if len(intervals) <= 1 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}

		return intervals[i][0] < intervals[j][0]
	})

	count := 0
	current := 0

	for _, arr := range intervals {
		if current < arr[1] {
			count++
			current = arr[1]
		}
	}

	return count

}