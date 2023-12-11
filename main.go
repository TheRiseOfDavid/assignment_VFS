package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
	"virtual_file_system/file"
	"virtual_file_system/folder"
	objectbaseinfo "virtual_file_system/object_base_info"
	objectmanager "virtual_file_system/object_manager"
	user "virtual_file_system/user"
)

type VirtualFileSystem struct {
	*objectmanager.ObjectManager[user.User]
	owners []user.User
}

func createVirtaulFileSystem() *VirtualFileSystem {
	return &VirtualFileSystem{owners: make([]user.User, 0)}
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

func (fs *VirtualFileSystem) registerUser(userName string) (string, error) {
	//需要再補一個 Error: The [username] contain invalid chars.
	if !isNameValid(userName) {
		return "", fmt.Errorf("Error: The %s contain invalid chars.\n", userName)
	}

	userName = strings.ToLower(userName) //進來程式後都小寫
	if fs.IsExists(userName) {
		return "", fmt.Errorf("Error: The %s has already existed.\n", userName)
	}

	owner := &user.User{
		ObjectManager: &objectmanager.ObjectManager[folder.Folder]{
			ObjectBaseInfo: &objectbaseinfo.ObjectBaseInfo{
				Name: userName,
			},
		},
		Folders: make([]folder.Folder, 0),
	}

	fs.owners = append(fs.owners, *owner)
	return fmt.Sprintf("Add %s successfull. \n", userName), nil
}

func (fs *VirtualFileSystem) createFolder(folderParma FolderOptions) (string, error) {
	/* 會用到的資訊
	type FolderOptions struct {
		userName    string
		folderName  string
		description string
	}
	*/

	owner, err := fs.SelectByName(folderParma.userName)
	if err != nil {
		return "", err // err 回傳 fmt.Errorf("Error: : The [%s] doesn't exist\n", userName)
	}
	if owner.IsExists(folderParma.folderName) {
		return "", fmt.Errorf("Error: [%s] has already existed\n", folderParma.folderName)
	}

	folder := &folder.Folder{
		ObjectManager: &objectmanager.ObjectManager[file.File]{
			ObjectBaseInfo: &objectbaseinfo.ObjectBaseInfo{
				Name:        folderParma.folderName,
				Description: folderParma.description,
				Created:     time.Now().Format("2006-01-02 15:04:05"),
			},
		},
		Files: make([]file.File, 0),
	}

	owner.Folders = append(owner.Folders, *folder)

	//fmt.Print("update The len of folders is ", len(owner.folders), "\n")
	return fmt.Sprintf("Create [%s] successfully. \n", folderParma.folderName), nil
}

func (fs *VirtualFileSystem) deleteFolder(folderParma FolderOptions) (string, error) {
	/* 會用到的資訊
	type FolderOptions struct {
		userName    string
		folderName  string
	}
	*/

	owner, err := fs.SelectByName(folderParma.userName)
	if err != nil {
		return "", err // err 回傳 fmt.Errorf("Error: : The [%s] doesn't exist\n", userName)
	}

	for i, folder := range owner.Folders {
		if folder.Name == folderParma.folderName {
			owner.Folders = append(owner.Folders[:i], owner.Folders[i+1:]...)
			return fmt.Sprintf("Delete [%s] successfully.\n", folderParma.folderName), nil
		}
	}

	return "", fmt.Errorf("Error: The [%s] doesn't exist\n", folderParma.folderName)
}

// 提供line interface 給 VFS
func (fs *VirtualFileSystem) commandShell() {

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
			err := fmt.Errorf("Invalid command. Usage: command [%s]\n", args[0])
			fmt.Fprint(os.Stdout, err)
			continue
			//return

		}

		// lengthFlag := true // 檢查字串長度是否超出 100 字元(文件要求)
		// for _, arg := range args {
		// 	if isLengthExcessive(arg, 100) || !isNameValid(arg) {
		// 		fmt.Errorf("Invalid command. Usage: command [%s]\n", arg)
		// 		lengthFlag = true
		// 		continue
		// 	}
		// }
		// if lengthFlag == true {
		// 	continue
		// }

		var msg string = ""
		var err error = nil

		//switch os.Args[0] {
		switch args[0] {
		case "register":
			msg, err = fs.registerUser(args[1])
			fmt.Println("hi")

		case "create-folder":
			parma := FolderOptions{userName: args[1], folderName: args[2]}
			if len(args) >= 4 {
				parma.description = args[3]
			}
			msg, err = fs.createFolder(parma)

		case "delete-folder":
			parma := FolderOptions{userName: args[1], folderName: args[2]}
			msg, err = fs.deleteFolder(parma)
		/*
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
		*/
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
	fs.commandShell()

}
