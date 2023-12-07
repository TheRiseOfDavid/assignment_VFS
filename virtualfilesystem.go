package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

type Folder struct {
	name        string
	description string
}

// VirtualFileSystem 出發點
type VirtualFileSystem struct {
	owners []User
}

func createVirtaulFileSystem() *VirtualFileSystem {
	return &VirtualFileSystem{owners: make([]User, 0)}
}

func isNameValid(str string) bool {
	for _, char := range str {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

type FolderOptions struct {
	userName    string
	folderName  string
	description string
}

func (fs *VirtualFileSystem) selectUser(userName string) (*User, error) {
	for i, owner := range fs.owners {
		if owner.name == userName {
			return &fs.owners[i], nil
		}
	}
	return nil, fmt.Errorf("Error: The [%s] doesn't exist\n", userName)
}

func (fs *VirtualFileSystem) createFolder(folderParma FolderOptions) (string, error) {
	if !isNameValid(folderParma.folderName) {
		return "", fmt.Errorf("Error: The [%s] contain invalid chars\n", folderParma.folderName)
	}

	folderParma.userName = strings.ToLower(folderParma.userName)     //進來程式後都小寫
	folderParma.folderName = strings.ToLower(folderParma.folderName) //進來程式後都小寫

	owner, err := fs.selectUser(folderParma.userName)
	if err != nil {
		return "", err // err 回傳 fmt.Errorf("Error: : The [%s] doesn't exist\n", userName)
	}

	fmt.Print("The len of folders is ", len(owner.folders), "\n")
	/*
		for _, folder := range owner.folders {
			if folder.name == folderParma.folderName {
				fmt.Print(folderParma.folderName, "\n")
			}
		}
	*/

	for _, folder := range owner.folders {
		if folder.name == folderParma.folderName {
			return "", fmt.Errorf("Error: [%s] has already existed\n", folderParma.folderName)
		}
	}

	folder := Folder{name: folderParma.folderName, description: folderParma.description}
	owner.folders = append(owner.folders, folder)
	fmt.Print("update The len of folders is ", len(owner.folders), "\n")
	return fmt.Sprintf("Create [%s] successfully. \n", folderParma.folderName), nil
}

type UnitTestOptions struct {
	isUnitTest int
	reader     io.Reader
}

// 提供line interface 給 VFS
func (fs *VirtualFileSystem) commandShell(utOption UnitTestOptions) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("vfs> ")
		scanner.Scan()
		command := scanner.Text()

		if command == "exit" {
			//if os.Args[1] == "exit" {
			//break
			return
		}

		args := strings.Fields(command)
		//if len(os.Args) < 2 {
		if len(args) < 2 {
			fmt.Println("Invalid command. Usage: command [arguments]")
			continue
			//return

		}

		//switch os.Args[0] {
		switch args[0] {
		case "register":
			//msg, err := fs.registerUser(os.Args[2])
			msg, err := fs.registerUser(args[1])
			if err != nil {
				fmt.Fprint(os.Stdout, err)
				continue
			}
			fmt.Fprint(os.Stdout, msg)

		case "create-folder":
			parma := FolderOptions{userName: args[1], folderName: args[2]}
			if len(args) >= 4 {
				parma.description = args[3]
			}

			msg, err := fs.createFolder(parma)
			if err != nil {
				fmt.Fprint(os.Stdout, err)
				continue
			}
			fmt.Fprint(os.Stdout, msg)
		}
	}
}

func main() {
	fs := createVirtaulFileSystem()
	fs.commandShell(UnitTestOptions{isUnitTest: 0})

}
