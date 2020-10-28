package dummy

import "testing"

func TestDummy(t *testing.T) {
	if Dummy() != 42 {
		t.Error("Unexpected value")
	}
}
