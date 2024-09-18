package main

import "testing"

func TestNotification(t *testing.T) {
	expenditure := []int32{2, 3, 4, 2, 3, 6, 8, 4, 5}
	d := int32(5)
	notifications := activityNotifications(expenditure, d)

	if notifications != 2 {
		t.Fatalf("wrong amount of notifications. got %d expect 2", notifications)
	}

}
