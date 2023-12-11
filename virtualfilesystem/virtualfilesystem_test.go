package virtualfilesystem

import (
	"reflect"
	"testing"
)

// CreateVirtaulFileSystem
func TestCreateVirtaulFileSystem(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	output := reflect.TypeOf(fs)
	expected := reflect.TypeOf(VirtualFileSystem{})
	if output == expected {
		t.Errorf("RegisterUserWithSuccess \nreturned %s\nexpected %s", output, expected)
	}

}

func TestRegisterUserWithSuccess(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	output, _ := fs.registerUser("david")

	expected := "Add [david] successfully.\n"
	if output != expected {
		t.Errorf("RegisterUserWithSuccess \nreturned %s\nexpected %s", output, expected)
	}

}

func TestRegisterUserWithErrorUsernameExisted(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, err := fs.registerUser("david")
	_, err = fs.registerUser("david")

	expected := "Error: The [david] has already existed.\n"
	if err.Error() != expected {
		t.Errorf("RegisterUserWithErrorUsernameExisted \nreturned %v\nexpected %s", err, expected)
	}

}

func TestRegisterUserWithErrorUsernameInvalid(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, err := fs.registerUser("david大")

	expected := "Error: The [david大] contain invalid chars.\n"
	if err.Error() != expected {
		t.Errorf("RegisterUserWithErrorUsernameExisted \nreturned %v\nexpected %s", err, expected)
	}

}
