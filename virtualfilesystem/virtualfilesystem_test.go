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

func TestCommandShellForSuccess(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	msg, _ := fs.scannerCommand("register david")
	expected := "Add [david] successfully.\n"
	if msg != expected {
		t.Errorf("TestCommandShellForSuccess \nreturned %s\nexpected %s", msg, expected)
	}
}

func TestCommandShellForErrorWithLeastCommand(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, err := fs.scannerCommand("register")
	expected := "Invalid command because miss argument(s). Usage: command [register]\n"
	if err.Error() != expected {
		t.Errorf("TestCommandShellForErrorWithLeastCommand \nreturned %s\nexpected %s", err.Error(), expected)
	}
}

func TestCommandShellForErrorWithInvalidChar(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, err := fs.scannerCommand("register david大")
	expected := "Invalid command. Usage: command [david大]\n"
	if err.Error() != expected {
		t.Errorf("TestCommandShellForErrorWithInvalidChar \nreturned %s\nexpected %s", err.Error(), expected)
	}
}

func TestCommandShellForErrorWithInvalidCommand(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, err := fs.scannerCommand("re david")
	expected := "Error: Unrecognized command [re]\n"
	if err.Error() != expected {
		t.Errorf("TestCommandShellForErrorWithInvalidCommand \nreturned %s\nexpected %s", err.Error(), expected)
	}
}

func TestCommandShellForListFileErrorWithIncorrectFlags(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, err := fs.scannerCommand("list-files user1 folder1 --sort-name a")
	expected := "Usage: list files [username] [foldername] [--sort-name|--sort-created] [asc|desc]\n"
	if err.Error() != expected {
		t.Errorf("TestCommandShellForErrorWithInvalidCommand \nreturned %s\nexpected %s", err.Error(), expected)
	}
}

func TestCommandShellForListFolderErrorWithIncorrectFlags(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, err := fs.scannerCommand("list-folders user1 folder1 --sort-name a")
	expected := "Usage: list files [username] [foldername] [--sort-name|--sort-created] [asc|desc]\n"
	if err.Error() != expected {
		t.Errorf("TestCommandShellForErrorWithInvalidCommand \nreturned %s\nexpected %s", err.Error(), expected)
	}
}
