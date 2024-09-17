package main

import (
	"log"
	"testing"
)

func TestRotate(t *testing.T) {
	i := []int32{1, 2, 3, 4, 5}

	left := rotLeft(i, 1)

	log.Println(left)
}
