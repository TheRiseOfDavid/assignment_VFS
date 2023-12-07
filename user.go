package main

import (
	"fmt"
	"strings"
)

type User struct {
	name    string
	folders []Folder
}

func (fs *VirtualFileSystem) registerUser(userName string) (string, error) {
	//需要再補一個 Error: The [username] contain invalid chars.
	if !isNameValid(userName) {
		return "", fmt.Errorf("Error: The %s contain invalid chars.\n", userName)
	}

	userName = strings.ToLower(userName) //進來程式後都小寫
	for _, owner := range fs.owners {
		if owner.name == userName {
			//fmt.Fprintf(os.Stdout, "Error: The %v has already existed.\n", owner.name)
			return "", fmt.Errorf("Error: The %s has already existed.\n", owner.name)
		}
	}

	owner := &User{name: userName, folders: make([]Folder, 0)}
	fs.owners = append(fs.owners, *owner)

	//fmt.Fprintf(os.Stdout, "Add %v successfull. \n", owner.name)
	return fmt.Sprintf("Add %v successfull. \n", owner.name), nil
}
