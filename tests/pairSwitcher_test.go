package Test

import (
	"reflect"
	"testing"
	"valorantAgentMeta/internal"
)

func TestSortByValues(t *testing.T) {
	data := map[string]int{
		"Deadlock": 35,
		"Chamber":  70,
	}

	want := internal.PairList{internal.Pair{Key: "Chamber", Value: 70}, internal.Pair{Key: "Deadlock", Value: 35}}
	got := internal.SortMapByValues(data)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
