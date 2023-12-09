package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
	"unicode"
)

type File struct {
	name        string
	description string
	created     string
}

type FileOptions struct {
	FolderOptions //繼承 folder 所需要的東西
	fileName      string
}

// VirtualFileSystem 出發點
type VirtualFileSystem struct {
	owners []User
}

func (f *Folder) selectFile(fileName string) (*File, error) {
	for i, file := range f.files {
		if file.name == fileName {
			return &f.files[i], nil
		}
	}
	return nil, fmt.Errorf("Error: The [%s] doesn't exist.\n", fileName)
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

// GO 似乎不支持 指標陣列多型使用
// https://stackoverflow.com/questions/54377597/how-to-make-a-function-that-receives-an-array-of-custom-interface

func (fs *VirtualFileSystem) createFile(fileParma FileOptions) (string, error) {
	/* 會用到的
	type FileOptions struct {
		userName      string
		folderName    string
		fileName      string
	}
	*/
	if !isNameValid(fileParma.fileName) { //Error: The [filename] contains invalid chars
		return "", fmt.Errorf("Error: The [%s] contains invalid chars.\n", fileParma.fileName)
	}

	owner, err := fs.selectUser(fileParma.userName) //Error: The [username] doesn't exist
	if err != nil {
		return "", err
	}
	folder, err := owner.selectFolder(fileParma.folderName)
	if err != nil {
		return "", err
	}

	folder.files = append(folder.files, File{name: fileParma.fileName, description: fileParma.description,
		created: time.Now().Format("2006-01-02 15:04:05")})
	return fmt.Sprintf("Create [%s] in [%s] / [%s] successfull.\n",
		fileParma.fileName, fileParma.userName, fileParma.folderName), nil
}

func (fs *VirtualFileSystem) deleteFile(fileParma FileOptions) (string, error) {
	owner, err := fs.selectUser(fileParma.userName) //Error: The [username] doesn't exist
	if err != nil {
		return "", err
	}
	folder, err := owner.selectFolder(fileParma.folderName)
	if err != nil {
		return "", err
	}

	for i, file := range folder.files {
		if file.name == fileParma.folderName {
			folder.files = append(folder.files[:i], folder.files[i+1:]...)
			return fmt.Sprintf("Delete [%s] successfully.\n", fileParma.folderName), nil
		}
	}

	return "", fmt.Errorf("Error: The [%s] doesn't exist\n", fileParma.folderName)
}

func (fs *VirtualFileSystem) listFiles(fileParma FileOptions) (string, error) {
	/* 會用到的資訊
	type FolderOptions struct {
		fileName     string
		folderName   string
		userName     string
		sortCriteria string
		sortby       string
	}
	*/

	owner, err := fs.selectUser(fileParma.userName) //Error: The [username] doesn't exist
	if err != nil {
		return "", err
	}
	folder, err := owner.selectFolder(fileParma.folderName)
	if err != nil {
		return "", err
	}

	if len(folder.files) == 0 {
		return "", fmt.Errorf("Warning: The [%s] doesn't have any files.\n", fileParma.userName)
	}

	if fileParma.sortCriteria == "--sort-name" {
		sort.Slice(folder.files, func(i int, j int) bool {
			if fileParma.sortby == "asc" {
				return folder.files[i].name < folder.files[j].name
			} else {
				return folder.files[i].name > folder.files[j].name
			}
		})
	}
	if fileParma.sortCriteria == "--sort-created" {
		sort.Slice(folder.files, func(i int, j int) bool {
			if fileParma.sortby == "asc" {
				return folder.files[i].created < folder.files[j].created
			} else {
				return folder.files[i].created > folder.files[j].created
			}
		})
	}

	result := ""
	for _, folder := range folder.files {
		result += fmt.Sprintf("[%s] [%s] [%s] [%s] \n", folder.name, folder.description, folder.created, fileParma.userName)
	}
	return result, nil
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
			fmt.Errorf("Invalid command. Usage: command [%s]", args[0])
			continue
			//return

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
