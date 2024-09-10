package getResult

import (
	"fmt"
	"valorantAgentMeta/pair"
)

func SortAndPrintResult(mapWithIntValue map[string]int) {

	result := pair.SortMapByValues(mapWithIntValue)
	for i := 0; i < 5; i++ {
		fmt.Print(result[i].Key, " ")
	}
}
