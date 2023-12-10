package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

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

func isLengthExcessive(args string, limit int) bool {
	if len(args) > limit {
		return true
	}
	return false
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
			fmt.Errorf("Invalid command. Usage: command [%s]\n", args[0])
			continue
			//return

		}

		lengthFlag := true // 檢查字串長度是否超出 100 字元(文件要求)
		for _, arg := range args {
			if isLengthExcessive(arg, 100) || !isNameValid(arg) {
				fmt.Errorf("Invalid command. Usage: command [%s]\n", arg)
				lengthFlag = true
				continue
			}
		}
		if lengthFlag == true {
			continue
		}

		var msg string = ""
		var err error = nil

		//switch os.Args[0] {
		switch args[0] {
		case "register":
			//msg, err := fs.registerUser(os.Args[2])
			msg, err = fs.registerUser(args[1])

		case "create-folder":
			parma := FolderOptions{userName: args[1], folderName: args[2]}
			if len(args) >= 4 {
				parma.description = args[3]
			}
			msg, err = fs.createFolder(parma)

		case "delete-folder":
			parma := FolderOptions{userName: args[1], folderName: args[2]}
			msg, err = fs.deleteFolder(parma)
		case "list-folders":
			parma := FolderOptions{userName: args[1], sortCriteria: "--sort-name", sortby: "asc"}
			if len(args) >= 3 {
				parma.sortCriteria = args[2]
			}
			if len(args) >= 4 {
				parma.sortby = args[3]
			}
			msg, err = fs.listFolders(parma)

		case "rename-folder":
			parma := FolderOptions{userName: args[1], folderName: args[2], newFolderName: args[3]}
			msg, err = fs.renameFolder(parma)
		case "create-file":
			parma := FileOptions{FolderOptions: FolderOptions{userName: args[1], folderName: args[2]},
				fileName: args[3]}
			if len(args) >= 5 {
				parma.description = args[4]
			}
			msg, err = fs.createFile(parma)

		case "delete-file":
			parma := FileOptions{FolderOptions: FolderOptions{userName: args[1], folderName: args[2]}, fileName: args[3]}
			msg, err = fs.deleteFile(parma)

		case "list-files":
			parma := FileOptions{FolderOptions: FolderOptions{userName: args[1], folderName: args[2], sortCriteria: "--sort-name", sortby: "asc"}}
			if len(args) >= 4 {
				parma.sortCriteria = args[3]
			}
			if len(args) >= 5 {
				parma.sortby = args[4]
			}
			msg, err = fs.listFiles(parma)
		default:
			err = fmt.Errorf("Invalid command. Usage: command [%s]", args[0])
		}

		if err != nil {
			fmt.Fprint(os.Stdout, err)
			continue
		}
		fmt.Fprint(os.Stdout, msg)
	}
}

func main() {
	fs := createVirtaulFileSystem()
	fs.commandShell(UnitTestOptions{isUnitTest: 0})

}
