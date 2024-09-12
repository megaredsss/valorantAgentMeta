package internal

import (
	"strconv"
)

var saveMap map[string]string

// ConvertMap converting map[string]string to [string]int
// delete Valorant Map from map and return new map
func ConvertMap(table Table, testValMap string) map[string]int {
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
