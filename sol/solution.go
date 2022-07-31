package sol

import "sort"

/**
 * Definition of Interval:
 * type Interval struct {
 *     Start, End int
 * }
 */
/**
 * @param intervals: an array of meeting time intervals
 * @return: the minimum number of conference rooms required
 */
func MinMeetingRooms(intervals []*Interval) int {
	count, result := 0, 0
	n := len(intervals)
	start, end := make([]int, n), make([]int, n)
	for pos := 0; pos < n; pos++ {
		start[pos] = intervals[pos].Start
		end[pos] = intervals[pos].End
	}
	sort.Ints(start)
	sort.Ints(end)
	pStart, pEnd := 0, 0
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for pStart < n {
		if start[pStart] < end[pEnd] {
			pStart++
			count++
		} else {
			pEnd++
			count--
		}
		result = max(result, count)
	}
	return result
}
