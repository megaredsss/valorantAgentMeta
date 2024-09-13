package Test

import (
	"reflect"
	"testing"
	"valorantAgentMeta/internal"
)

func TestRemovePercent(t *testing.T) {
	got := internal.Table{internal.Row{
		"Chamber":  "15%",
		"Deadlock": "18%",
	}}
	want := internal.Table{internal.Row{
		"Chamber":  "15",
		"Deadlock": "18",
	}}
	internal.RemovePercent(got)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
