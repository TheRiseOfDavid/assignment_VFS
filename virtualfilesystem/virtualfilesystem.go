package virtualfilesystem

import (
	"VFS/file"
	"VFS/folder"
	objectbaseinfo "VFS/object_base_info"
	"VFS/options"
	"VFS/user"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

// VirtualFileSystem 出發點
type VirtualFileSystem struct {
	owners []user.User
}

func CreateVirtaulFileSystem() *VirtualFileSystem {
	return &VirtualFileSystem{owners: make([]user.User, 0)}
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
		return "", fmt.Errorf("Error: The [%s] contain invalid chars.\n", userName)
	}
	if fs.isUserExist(userName) {
		return "", fmt.Errorf("Error: The [%s] has already existed.\n", userName)
	}

	owner := &user.User{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{Name: userName},
		Folders: make([]folder.Folder, 0)}
	fs.owners = append(fs.owners, *owner)

	return fmt.Sprintf("Add [%s] successfully.\n", owner.Name), nil
}

func (fs *VirtualFileSystem) createFolder(folderParma options.FolderOptions) (string, error) {
	if !isNameValid(folderParma.FolderName) {
		return "", fmt.Errorf("Error: The [%s] contain invalid chars\n", folderParma.FolderName)
	}

	owner, err := fs.selectUser(folderParma.UserName)
	if err != nil {
		return "", err // err 回傳 fmt.Errorf("Error: : The [%s] doesn't exist\n", userName)
	}

	if owner.IsFolderExists(folderParma.FolderName) {
		return "", fmt.Errorf("Error: [%s] has already existed\n", folderParma.FolderName)
	}

	folder := folder.Folder{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: folderParma.FolderName, Description: folderParma.Description,
		Created: time.Now().Format("2006-01-02 15:04:05"),
	},
		Files: make([]file.File, 0)}
	owner.Folders = append(owner.Folders, folder)
	//fmt.Print("update The len of folders is ", len(owner.folders), "\n")
	return fmt.Sprintf("Create [%s] successfully. \n", folderParma.FolderName), nil
}

func (fs *VirtualFileSystem) deleteFolder(folderParma options.FolderOptions) (string, error) {
	/* 會用到的資訊
	type FolderOptions struct {
		userName    string
		folderName  string
	}
	*/

	folderParma.UserName = strings.ToLower(folderParma.UserName)     //進來程式後都小寫
	folderParma.FolderName = strings.ToLower(folderParma.FolderName) //進來程式後都小寫

	owner, err := fs.selectUser(folderParma.UserName)
	if err != nil {
		return "", err // err 回傳 fmt.Errorf("Error: : The [%s] doesn't exist\n", userName)
	}

	for i, folder := range owner.Folders {
		if folder.Name == folderParma.FolderName {
			owner.Folders = append(owner.Folders[:i], owner.Folders[i+1:]...)
			return fmt.Sprintf("Delete [%s] successfully.\n", folderParma.FolderName), nil
		}
	}

	return "", fmt.Errorf("Error: The [%s] doesn't exist\n", folderParma.FolderName)
}

func (fs *VirtualFileSystem) listFolders(folderParma options.FolderOptions) (string, error) {
	/* 會用到的資訊
	type FolderOptions struct {
		userName     string
		sortCriteria string
		sortby       string
	}
	*/

	folderParma.UserName = strings.ToLower(folderParma.UserName)     //進來程式後都小寫
	folderParma.FolderName = strings.ToLower(folderParma.FolderName) //進來程式後都小寫

	owner, err := fs.selectUser(folderParma.UserName)
	if err != nil {
		return "", err
	}

	if len(owner.Folders) == 0 {
		return "", fmt.Errorf("Warning: The [%s] doesn't have any folders.\n", folderParma.UserName)
	}

	if folderParma.SortCriteria == "--sort-name" {
		sort.Slice(owner.Folders, func(i int, j int) bool {
			if folderParma.Sortby == "asc" {
				return owner.Folders[i].Name < owner.Folders[j].Name
			} else {
				return owner.Folders[i].Name > owner.Folders[j].Name
			}
		})
	}
	if folderParma.SortCriteria == "--sort-created" {
		sort.Slice(owner.Folders, func(i int, j int) bool {
			if folderParma.Sortby == "asc" {
				return owner.Folders[i].Created < owner.Folders[j].Created
			} else {
				return owner.Folders[i].Created > owner.Folders[j].Created
			}
		})
	}

	result := ""
	for _, folder := range owner.Folders {
		result += fmt.Sprintf("[%s] [%s] [%s] [%s] \n", folder.Name, folder.Description, folder.Created, folderParma.UserName)
	}
	return result, nil
}

func (fs *VirtualFileSystem) renameFolder(folderParma options.FolderOptions) (string, error) {
	//會用到的資訊
	/*
		type FolderOptions struct {
			userName      string
			folderName    string
			newFolderName string
		}
	*/

	folderParma.UserName = strings.ToLower(folderParma.UserName)     //進來程式後都小寫
	folderParma.FolderName = strings.ToLower(folderParma.FolderName) //進來程式後都小寫

	owner, err := fs.selectUser(folderParma.UserName)
	if err != nil {
		return "", err
	}

	for i, folder := range owner.Folders {
		if folder.Name == folderParma.FolderName {
			owner.Folders[i].Name = folderParma.NewFolderName
			return fmt.Sprintf("Rename [%s] to [%s] successfully.\n",
				folderParma.FolderName, folderParma.NewFolderName), nil
		}
	}
	return "", fmt.Errorf("Error: The [%s] doesn't exist\n", folderParma.FolderName)
}

// GO 似乎不支持 指標陣列多型使用
// https://stackoverflow.com/questions/54377597/how-to-make-a-function-that-receives-an-array-of-custom-interface

func (fs *VirtualFileSystem) createFile(fileParma options.FileOptions) (string, error) {
	/* 會用到的
	type FileOptions struct {
		userName      string
		folderName    string
		fileName      string
	}
	*/
	if !isNameValid(fileParma.FileName) { //Error: The [filename] contains invalid chars
		return "", fmt.Errorf("Error: The [%s] contains invalid chars.\n", fileParma.FileName)
	}

	owner, err := fs.selectUser(fileParma.UserName) //Error: The [username] doesn't exist
	if err != nil {
		return "", err
	}
	folder, err := owner.SelectFolder(fileParma.FolderName)
	if err != nil {
		return "", err
	}

	folder.Files = append(folder.Files, file.File{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: fileParma.FileName, Description: fileParma.Description,
		Created: time.Now().Format("2006-01-02 15:04:05")},
	})
	return fmt.Sprintf("Create [%s] in [%s] / [%s] successfull.\n",
		fileParma.FileName, fileParma.UserName, fileParma.FolderName), nil
}

func (fs *VirtualFileSystem) deleteFile(fileParma options.FileOptions) (string, error) {
	owner, err := fs.selectUser(fileParma.UserName) //Error: The [username] doesn't exist
	if err != nil {
		return "", err
	}
	folder, err := owner.SelectFolder(fileParma.FolderName)
	if err != nil {
		return "", err
	}

	for i, file := range folder.Files {
		if file.Name == fileParma.FolderName {
			folder.Files = append(folder.Files[:i], folder.Files[i+1:]...)
			return fmt.Sprintf("Delete [%s] successfully.\n", fileParma.FolderName), nil
		}
	}

	return "", fmt.Errorf("Error: The [%s] doesn't exist\n", fileParma.FolderName)
}

func (fs *VirtualFileSystem) listFiles(fileParma options.FileOptions) (string, error) {
	/* 會用到的資訊
	type FolderOptions struct {
		fileName     string
		folderName   string
		userName     string
		sortCriteria string
		sortby       string
	}
	*/

	owner, err := fs.selectUser(fileParma.UserName) //Error: The [username] doesn't exist
	if err != nil {
		return "", err
	}
	folder, err := owner.SelectFolder(fileParma.FolderName)
	if err != nil {
		return "", err
	}

	if len(folder.Files) == 0 {
		return "", fmt.Errorf("Warning: The [%s] doesn't have any files.\n", fileParma.UserName)
	}

	if fileParma.SortCriteria == "--sort-name" {
		sort.Slice(folder.Files, func(i int, j int) bool {
			if fileParma.Sortby == "asc" {
				return folder.Files[i].Name < folder.Files[j].Name
			} else {
				return folder.Files[i].Name > folder.Files[j].Name
			}
		})
	}
	if fileParma.SortCriteria == "--sort-created" {
		sort.Slice(folder.Files, func(i int, j int) bool {
			if fileParma.Sortby == "asc" {
				return folder.Files[i].Created < folder.Files[j].Created
			} else {
				return folder.Files[i].Created > folder.Files[j].Created
			}
		})
	}

	result := ""
	for _, folder := range folder.Files {
		result += fmt.Sprintf("[%s] [%s] [%s] [%s] \n", folder.Name, folder.Description, folder.Created, fileParma.UserName)
	}
	return result, nil
}

// 提供line interface 給 VFS
func (fs *VirtualFileSystem) CommandShell() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("vfs> ")
		scanner.Scan()
		command := scanner.Text()
		command = strings.TrimSpace(command) //消除前後的 space

		if command == "exit" {
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

		lengthFlag := true // 檢查字串長度是否超出 100 字元(文件要求)
		for _, arg := range args {
			if isLengthExcessive(arg, 100) || !isNameValid(arg) {
				err := fmt.Errorf("Invalid command. Usage: command [%s]\n", arg)
				fmt.Fprint(os.Stderr, err)
				lengthFlag = true
				continue
			}
		}
		if lengthFlag == false { //通過測驗就執行命令
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
			parma := options.FolderOptions{UserName: args[1], FolderName: args[2]}
			if len(args) >= 4 {
				parma.Description = args[3]
			}
			msg, err = fs.createFolder(parma)

		case "delete-folder":
			parma := options.FolderOptions{UserName: args[1], FolderName: args[2]}
			msg, err = fs.deleteFolder(parma)
		case "list-folders":
			parma := options.FolderOptions{UserName: args[1], SortCriteria: "--sort-name", Sortby: "asc"}
			if len(args) >= 3 {
				parma.SortCriteria = args[2]
			}
			if len(args) >= 4 {
				parma.Sortby = args[3]
			}
			msg, err = fs.listFolders(parma)

		case "rename-folder":
			parma := options.FolderOptions{UserName: args[1], FolderName: args[2], NewFolderName: args[3]}
			msg, err = fs.renameFolder(parma)
		case "create-file":
			parma := options.FileOptions{FolderOptions: options.FolderOptions{UserName: args[1], FolderName: args[2]},
				FileName: args[3]}
			if len(args) >= 5 {
				parma.Description = args[4]
			}
			msg, err = fs.createFile(parma)

		case "delete-file":
			parma := options.FileOptions{FolderOptions: options.FolderOptions{UserName: args[1], FolderName: args[2]}, FileName: args[3]}
			msg, err = fs.deleteFile(parma)

		case "list-files":
			parma := options.FileOptions{FolderOptions: options.FolderOptions{UserName: args[1], FolderName: args[2], SortCriteria: "--sort-name", Sortby: "asc"}}
			if len(args) >= 4 {
				parma.SortCriteria = args[3]
			}
			if len(args) >= 5 {
				parma.Sortby = args[4]
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
