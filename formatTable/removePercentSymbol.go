package formatTable

import (
	"strings"
	"valorantAgentMeta/tableReader"
)

func RemovePercent(table tableReader.Table) {
	for i := range table {
		for key, value := range table[i] {
			newValue := strings.ReplaceAll(value, "%", "")
			table[i][key] = newValue
		}
	}
}
