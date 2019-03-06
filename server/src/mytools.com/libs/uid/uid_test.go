package uid

import "testing"

func Test_UIDGen(t *testing.T) {
	record := make(map[int64]bool)
	for i := 0; i < 3000000; i++ {
		id := Gen()
		if _, ok := record[id]; ok {
			t.FailNow()
		}
		record[id] = true
	}
}
