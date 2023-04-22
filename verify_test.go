package twID

import (
	"testing"
)

func TestVerifyTrue_Corrent(t *testing.T) {
	want := true
	msg := Verify("a123456789")
	if msg != want {
		t.Fatalf("Error")
	}
}

func TestVerifyFalse_FormatErr(t *testing.T) {
	want := false
	msg := Verify("AAAAAADDDDDD")
	if msg != want {
		t.Fatalf("Error")
	}
}

func TestVerifyFalse_CheckNumErr(t *testing.T) {
	want := false
	msg := Verify("a123456788")
	if msg != want {
		t.Fatalf("Error")
	}
}

func TestVerifyFalse_OtherLanguageErr(t *testing.T) {
	want := false
	msg := Verify("輸入中文不行的")
	if msg != want {
		t.Fatalf("Error")
	}
}
