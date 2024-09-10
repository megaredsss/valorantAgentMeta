package formatTable

import (
	"strconv"
	"valorantAgentMeta/tableReader"
)

var saveMap map[string]string

func ConvertMap(table tableReader.Table, testValMap string) map[string]int {
	mapWithIntValue := make(map[string]int)
	for i := range table {
		for key, value := range table[i] {
			if key == "Map" && value == testValMap {
				delete(table[i], "Map")
				saveMap = table[i]
			}
		}

	}
	for key, value := range saveMap {
		mapWithIntValue[key], _ = strconv.Atoi(value)
	}
	return mapWithIntValue
}
