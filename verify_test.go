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

func TestVerifyFalse_3(t *testing.T) {
	want := false
	msg := Verify("a1234,5678")
	if msg != want {
		t.Fatalf("Error")
	}
}

func TestVerifyFalse_4(t *testing.T) {
	want := false
	msg := Verify("a12/u2014")
	if msg != want {
		t.Fatalf("Error")
	}
}
