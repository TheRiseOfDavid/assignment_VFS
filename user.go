package main

import (
	"fmt"
	"strings"
)

type User struct {
	name    string
	folders []Folder
}

func (owner *User) selectFolder(folderName string) (*Folder, error) {
	for i, folder := range owner.folders {
		if folder.name == folderName {
			return &owner.folders[i], nil
		}
	}
	return nil, fmt.Errorf("Error: The [%s] doesn't exist.\n", folderName)
}

func (owner *User) isFolderExists(folderName string) bool {
	for _, folder := range owner.folders {
		if folder.name == folderName {
			return true
		}
	}
	return false
}

func (fs *VirtualFileSystem) isUserExist(userName string) bool {
	for _, owner := range fs.owners {
		if owner.name == userName {
			//fmt.Fprintf(os.Stdout, "Error: The %v has already existed.\n", owner.name)
			return true
		}
	}
	return false
}

func (fs *VirtualFileSystem) selectUser(userName string) (*User, error) {
	for i, owner := range fs.owners {
		if owner.name == userName {
			return &fs.owners[i], nil // 細節: golang iterator 是 copy constructor
		}
	}
	return nil, fmt.Errorf("Error: The [%s] doesn't exist\n", userName)
}

func (fs *VirtualFileSystem) registerUser(userName string) (string, error) {
	//需要再補一個 Error: The [username] contain invalid chars.
	if !isNameValid(userName) {
		return "", fmt.Errorf("Error: The %s contain invalid chars.\n", userName)
	}

	userName = strings.ToLower(userName) //進來程式後都小寫
	if fs.isUserExist(userName) {
		return "", fmt.Errorf("Error: The %s has already existed.\n", userName)
	}

	owner := &User{name: userName, folders: make([]Folder, 0)}
	fs.owners = append(fs.owners, *owner)

	//fmt.Fprintf(os.Stdout, "Add %v successfull. \n", owner.name)
	return fmt.Sprintf("Add %v successfull. \n", owner.name), nil
}
