package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
	"unicode"
	"virtualfilesystem/file"
	"virtualfilesystem/folder"
	objectbaseinfo "virtualfilesystem/object_base_info"
	"virtualfilesystem/user"
)

// VirtualFileSystem 出發點
type VirtualFileSystem struct {
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

func (fs *VirtualFileSystem) isUserExist(userName string) bool {
	for _, owner := range fs.owners {
		if owner.Name == userName {
			//fmt.Fprintf(os.Stdout, "Error: The %v has already existed.\n", owner.name)
			return true
		}
	}
	return false
}

func (fs *VirtualFileSystem) selectUser(userName string) (*user.User, error) {
	for i, owner := range fs.owners {
		if owner.Name == userName {
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
	if fs.isUserExist(userName) {
		return "", fmt.Errorf("Error: The %s has already existed.\n", userName)
	}

	owner := &user.User{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{Name: userName},
		Folders: make([]folder.Folder, 0)}
	fs.owners = append(fs.owners, *owner)

	return fmt.Sprintf("Add %s successfull. \n", owner.Name), nil
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

	//fmt.Print("The len of folders is ", len(owner.folders), "\n")
	/*
		for _, folder := range owner.folders {
			if folder.name == folderParma.folderName {
				fmt.Print(folderParma.folderName, "\n")
			}
		}
	*/

	if owner.IsFolderExists(folderParma.folderName) {
		return "", fmt.Errorf("Error: [%s] has already existed\n", folderParma.folderName)
	}

	folder := folder.Folder{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: folderParma.folderName, Description: folderParma.description,
		Created: time.Now().Format("2006-01-02 15:04:05"),
	},
		Files: make([]file.File, 0)}
	owner.Folders = append(owner.Folders, folder)
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

	folderParma.userName = strings.ToLower(folderParma.userName)     //進來程式後都小寫
	folderParma.folderName = strings.ToLower(folderParma.folderName) //進來程式後都小寫

	owner, err := fs.selectUser(folderParma.userName)
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

func (fs *VirtualFileSystem) listFolders(folderParma FolderOptions) (string, error) {
	/* 會用到的資訊
	type FolderOptions struct {
		userName     string
		sortCriteria string
		sortby       string
	}
	*/

	folderParma.userName = strings.ToLower(folderParma.userName)     //進來程式後都小寫
	folderParma.folderName = strings.ToLower(folderParma.folderName) //進來程式後都小寫

	owner, err := fs.selectUser(folderParma.userName)
	if err != nil {
		return "", err
	}

	if len(owner.Folders) == 0 {
		return "", fmt.Errorf("Warning: The [%s] doesn't have any folders.\n", folderParma.userName)
	}

	if folderParma.sortCriteria == "--sort-name" {
		sort.Slice(owner.Folders, func(i int, j int) bool {
			if folderParma.sortby == "asc" {
				return owner.Folders[i].Name < owner.Folders[j].Name
			} else {
				return owner.Folders[i].Name > owner.Folders[j].Name
			}
		})
	}
	if folderParma.sortCriteria == "--sort-created" {
		sort.Slice(owner.Folders, func(i int, j int) bool {
			if folderParma.sortby == "asc" {
				return owner.Folders[i].Created < owner.Folders[j].Created
			} else {
				return owner.Folders[i].Created > owner.Folders[j].Created
			}
		})
	}

	result := ""
	for _, folder := range owner.Folders {
		result += fmt.Sprintf("[%s] [%s] [%s] [%s] \n", folder.Name, folder.Description, folder.Created, folderParma.userName)
	}
	return result, nil
}

func (fs *VirtualFileSystem) renameFolder(folderParma FolderOptions) (string, error) {
	//會用到的資訊
	/*
		type FolderOptions struct {
			userName      string
			folderName    string
			newFolderName string
		}
	*/

	folderParma.userName = strings.ToLower(folderParma.userName)     //進來程式後都小寫
	folderParma.folderName = strings.ToLower(folderParma.folderName) //進來程式後都小寫

	owner, err := fs.selectUser(folderParma.userName)
	if err != nil {
		return "", err
	}

	for i, folder := range owner.Folders {
		if folder.Name == folderParma.folderName {
			owner.Folders[i].Name = folderParma.newFolderName
			return fmt.Sprintf("Rename [%s] to [%s] successfully.\n",
				folderParma.folderName, folderParma.newFolderName), nil
		}
	}
	return "", fmt.Errorf("Error: The [%s] doesn't exist\n", folderParma.folderName)
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
	folder, err := owner.SelectFolder(fileParma.folderName)
	if err != nil {
		return "", err
	}

	folder.Files = append(folder.Files, file.File{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: fileParma.fileName, Description: fileParma.description,
		Created: time.Now().Format("2006-01-02 15:04:05")},
	})
	return fmt.Sprintf("Create [%s] in [%s] / [%s] successfull.\n",
		fileParma.fileName, fileParma.userName, fileParma.folderName), nil
}

func (fs *VirtualFileSystem) deleteFile(fileParma FileOptions) (string, error) {
	owner, err := fs.selectUser(fileParma.userName) //Error: The [username] doesn't exist
	if err != nil {
		return "", err
	}
	folder, err := owner.SelectFolder(fileParma.folderName)
	if err != nil {
		return "", err
	}

	for i, file := range folder.Files {
		if file.Name == fileParma.folderName {
			folder.Files = append(folder.Files[:i], folder.Files[i+1:]...)
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
	folder, err := owner.SelectFolder(fileParma.folderName)
	if err != nil {
		return "", err
	}

	if len(folder.Files) == 0 {
		return "", fmt.Errorf("Warning: The [%s] doesn't have any files.\n", fileParma.userName)
	}

	if fileParma.sortCriteria == "--sort-name" {
		sort.Slice(folder.Files, func(i int, j int) bool {
			if fileParma.sortby == "asc" {
				return folder.Files[i].Name < folder.Files[j].Name
			} else {
				return folder.Files[i].Name > folder.Files[j].Name
			}
		})
	}
	if fileParma.sortCriteria == "--sort-created" {
		sort.Slice(folder.Files, func(i int, j int) bool {
			if fileParma.sortby == "asc" {
				return folder.Files[i].Created < folder.Files[j].Created
			} else {
				return folder.Files[i].Created > folder.Files[j].Created
			}
		})
	}

	result := ""
	for _, folder := range folder.Files {
		result += fmt.Sprintf("[%s] [%s] [%s] [%s] \n", folder.Name, folder.Description, folder.Created, fileParma.userName)
	}
	return result, nil
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
			fmt.Fprint(os.Stderr, err)
			continue
			//return

		}

		// lengthFlag := true // 檢查字串長度是否超出 100 字元(文件要求)
		// for _, arg := range args {
		// 	if isLengthExcessive(arg, 100) || !isNameValid(arg) {
		// 		err := fmt.Errorf("Invalid command. Usage: command [%s]\n", arg)
		// 		fmt.Fprint(os.Stderr, err)
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
	fs.commandShell()

}
