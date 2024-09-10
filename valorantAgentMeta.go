package main

import (
	"fmt"
	"valorantAgentMeta/formatTable"
	"valorantAgentMeta/getResult"
	"valorantAgentMeta/tableReader"
)

func main() {

	table := tableReader.ReadTable()

	formatTable.RemovePercent(table)
	var testValMap string
	fmt.Println("Type map name(like this: Bind)")

	_, err := fmt.Scanln(&testValMap)
	if err != nil {
		fmt.Println("Error reading input:", err)
	}

	mapWithIntValue := formatTable.ConvertMap(table, testValMap)
	getResult.SortAndPrintResult(mapWithIntValue)
	fmt.Println("")
	fmt.Println("Scrapping Completed")

}
