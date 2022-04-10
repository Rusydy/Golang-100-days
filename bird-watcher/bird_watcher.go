package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	res := 0
	for i := 0; i < len(birdsPerDay); i++ {
		res += birdsPerDay[i]
	}
	return res
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	max := week * 7
	min := max - 7
	return TotalBirdCount(birdsPerDay[min:max])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, v := range birdsPerDay {
		if i == 0 || i%2 == 0 {
			birdsPerDay[i] = v + 1
		}
	}
	return birdsPerDay
}
