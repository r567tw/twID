package twID

import (
	"testing"
)

func TestVerify(t *testing.T) {
	want := false
	msg := Verify("AAAAAADDDDDD")
	if msg != want {
		t.Fatalf("Error")
	}
}
