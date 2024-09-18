package main

import (
	"log"
	"sort"
)

func maximumToys(prices []int32, k int32) int32 {
	// Write your code here

	sort.Slice(prices, func(i, j int) bool {
		return prices[i] < prices[j]
	})

	toys := 0
	for _, value := range prices {
		if k == 0 {
			break
		}

		if k >= value {
			toys++
			k = k - value
			continue
		}

		break
	}

	return int32(toys)
}

func main() {

	toys := []int32{1, 2, 3, 4}

	boughtToys := maximumToys(toys, 7)

	log.Println(boughtToys)
}
