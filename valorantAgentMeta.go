package main

import (
	"fmt"
	"valorantAgentMeta/formatTable"
	"valorantAgentMeta/getResult"
	"valorantAgentMeta/tableReader"
)

var testValMap string

func main() {

	table := tableReader.ReadTable()

	formatTable.RemovePercent(table)

	fmt.Println("Type map name(like this: Bind)")

	// get Map from user
	_, err := fmt.Scanln(&testValMap)
	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	mapWithIntValue := formatTable.ConvertMap(table, testValMap)

	getResult.SortAndPrintResult(mapWithIntValue)

}
