package virtualfilesystem

import (
	"VFS/user"
	"reflect"
	"testing"
)

func TestIsUserExistsForTrue(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	output := fs.isUserExists("david")

	expected := true
	if output != expected {
		t.Errorf("TestIsUserExistsForTrue \nreturned %t\nexpected %t", output, expected)
	}

}

func TestIsUserExistsForFalse(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	output := fs.isUserExists("david5")

	expected := false
	if output != expected {
		t.Errorf("TestIsUserExistsForTrue \nreturned %t\nexpected %t", output, expected)
	}
}

func TestSelectUserForSuccess(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	owner, _ := fs.selectUser("david")

	output := reflect.TypeOf(owner)
	expected := reflect.TypeOf(&user.User{})
	if output != expected {
		t.Errorf("TestSelectUserForSuccess \nreturned %s\nexpected %s", output, expected)
	}
}

func TestSelectUserForError(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	_, output := fs.selectUser("david2")

	expected := "Error: The [david2] doesn't exist.\n"
	if output.Error() != expected {
		t.Errorf("TestSelectUserForError \nreturned %s\nexpected %s", output, expected)
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
