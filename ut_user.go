package virtualfilesystem

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

func TestUser(t *testing.T) {

	fs := createVirtaulFileSystem()
	commandShell(fs)
	cmd := exec.Command("register", "user1")
	actualOutput, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	expectedOutput := []byte("Add user1 successfull. \n")

	if !bytes.Equal(expectedOutput, actualOutput) {
		t.Errorf("Expected output:\n%s\nActual output:\n%s", expectedOutput, actualOutput)
	}

}
