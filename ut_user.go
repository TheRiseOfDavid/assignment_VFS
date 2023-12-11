package main

// import (
// 	"bytes"
// 	"fmt"
// 	"os/exec"
// 	"strings"
// 	"testing"
// )

// func TestUser(t *testing.T) {
// 	input := "register user1\n"

// 	fs := createVirtaulFileSystem()
// 	fs.commandShell(UnitTestOptions{isUnitTest: 1, reader: strings.NewReader(input)})
// 	cmd := exec.Command("register", "user1")
// 	actualOutput, err := cmd.CombinedOutput()

// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		return
// 	}

// 	expectedOutput := []byte("Add user1 successfull. \n")

// 	if !bytes.Equal(expectedOutput, actualOutput) {
// 		t.Errorf("Expected output:\n%s\nActual output:\n%s", expectedOutput, actualOutput)
// 	}

// }
