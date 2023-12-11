package virtualfilesystem

import "testing"

func TestIsNameValidForTrue(t *testing.T) {
	output := isNameValid("david")
	expected := true
	if output != expected {
		t.Errorf("IsNameValidForTrue \nreturned %t\nexpected %t", output, expected)
	}
}

func TestIsNameValidForFalse(t *testing.T) {
	output := isNameValid("david大")
	expected := false
	if output != expected {
		t.Errorf("IsNameValidForTrue \nreturned %t\nexpected %t", output, expected)
	}
}

func TestIsLengthExcessiveForTrue(t *testing.T) {
	output := isLengthExcessive("david", 3)
	expected := true
	if output != expected {
		t.Errorf("IsNameValidForTrue \nreturned %t\nexpected %t", output, expected)
	}
}

func TestIsLengthExcessiveForFalse(t *testing.T) {
	output := isLengthExcessive("david", 10)
	expected := false
	if output != expected {
		t.Errorf("IsNameValidForTrue \nreturned %t\nexpected %t", output, expected)
	}
}
