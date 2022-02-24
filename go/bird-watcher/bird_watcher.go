package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	birdNumber := 0
	for _, v := range birdsPerDay {
		birdNumber += v
	}
	return birdNumber
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	return TotalBirdCount(birdsPerDay[(week-1)*7 : week*7])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for b := range birdsPerDay {
		if b%2 == 0 {
			birdsPerDay[b]++
		}
	}
	return birdsPerDay
}
