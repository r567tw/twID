package twID

import (
	"testing"
)

func TestVerifyFalse_1(t *testing.T) {
	want := false
	msg := Verify("AAAAAADDDDDD")
	if msg != want {
		t.Fatalf("Error")
	}
}

func TestVerifyTrue_1(t *testing.T) {
	want := true
	msg := Verify("a123456789")
	if msg != want {
		t.Fatalf("Error")
	}
}

func TestVerifyFalse_2(t *testing.T) {
	want := true
	msg := Verify("a123456788")
	if msg != want {
		t.Fatalf("Error")
	}
}
