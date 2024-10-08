package internal

import (
	"fmt"
)

// SortAndPrintResult sorting map and printing 5 best agents on the Valorant Map
func SortAndPrintResult(mapWithIntValue map[string]int) {

	result := SortMapByValues(mapWithIntValue)
	for i := 0; i < 5; i++ {
		fmt.Print(result[i].Key, " ")
	}
}
