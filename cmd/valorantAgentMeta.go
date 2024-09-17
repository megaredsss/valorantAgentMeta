package main

import (
	"fmt"
	"valorantAgentMeta/internal"
)

var Map string

func main() {

	table := internal.ReadTable()

	internal.RemovePercent(table)

	fmt.Println("Type map name(like this: Bind)")

	// get Map from user
	_, err := fmt.Scanln(&Map)
	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	mapWithIntValue := internal.ConvertMap(table, Map)

	internal.SortAndPrintResult(mapWithIntValue)

}
