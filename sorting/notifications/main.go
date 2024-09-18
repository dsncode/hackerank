package main

import (
	"cmp"
	"slices"
)

func computeMedian(ref []int32) int32 {

	slices.SortFunc(ref, func(i, j int32) int {
		return cmp.Compare(i, j)
	})

	if len(ref)%2 != 0 {
		return ref[len(ref)/2]
	}
	return (ref[len(ref)/2] + ref[len(ref)/2+1]) / 2
}

func activityNotifications(expenditure []int32, d int32) int32 {

	var notifications int32 = 0

	for i, value := range expenditure {

		if int32(i)-d < 0 {
			continue
		}

		referenceDays := expenditure[i-int(d) : i]
		median := computeMedian(referenceDays)

		threshold := median * 2

		if value >= threshold {
			//log.Println("expenses: ", value, " >= threshold:", threshold, "ref: ", referenceDays)
			notifications++
		}

	}

	return notifications
}

func main() {

}
