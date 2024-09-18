package main

import (
	"fmt"
)

func checkWords(magazine []string, note []string) bool {

	db := make(map[string]int)

	for _, value := range magazine {
		count, exists := db[value]
		if exists {
			db[value] = count + 1
			continue
		}
		db[value] = 1
	}

	for _, word := range note {

		count, exists := db[word]
		if !exists {
			return false
		}

		if count == 1 {
			delete(db, word)
			continue
		}

		db[word] = count - 1
	}

	return true
}

func checkMagazine(magazine []string, note []string) {
	if checkWords(magazine, note) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func main() {

	magazine := []string{"give", "me", "one", "grand", "today", "night"}
	note := []string{"give", "one", "grand", "today"}

	checkMagazine(magazine, note)

}
