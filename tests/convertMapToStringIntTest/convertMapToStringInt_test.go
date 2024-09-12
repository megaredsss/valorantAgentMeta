package convertMapToStringIntTest

import (
	"reflect"
	"testing"
	"valorantAgentMeta/internal"
)

func TestConvertMap(t *testing.T) {
	stringStringMap := internal.Table{internal.Row{
		"Map":      "Bind",
		"Deadlock": "5",
		"Raze":     "60",
	}}
	want := map[string]int{
		"Deadlock": 5,
		"Raze":     60,
	}
	result := internal.ConvertMap(stringStringMap, "Bind")
	eq := reflect.DeepEqual(result, want)
	if !eq {
		t.Fatalf("convert map failed, got: %v, want: %v", result, want)
	}
}
