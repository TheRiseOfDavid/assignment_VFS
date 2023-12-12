package virtualfilesystem

import (
	"VFS/options"
	"VFS/user"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// VirtualFileSystem 出發點
type VirtualFileSystem struct {
	owners []user.User
}

func CreateVirtaulFileSystem() *VirtualFileSystem {
	return &VirtualFileSystem{owners: make([]user.User, 0)}
}

func (fs *VirtualFileSystem) CommandShell() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("vfs> ")
		scanner.Scan()
		command := scanner.Text()
		if command == "exit" {
			return
		}

		command = strings.TrimSpace(command) //消除前後的 space
		msg, err := fs.scannerCommand(command)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			continue
		}
		fmt.Fprint(os.Stdout, msg)
	}
}

// 提供line interface 給 VFS
func (fs *VirtualFileSystem) scannerCommand(command string) (string, error) {

	var msg string = ""
	var err error = nil

	args := strings.Fields(command)
	//if len(os.Args) < 2 {
	if len(args) < 2 {
		err := fmt.Errorf("Invalid command because miss argument(s). Usage: command [%s]\n", args[0])
		return msg, err
	}

	lengthFlag := true // 檢查字串長度是否超出 100 字元(文件要求)
	for _, arg := range args[1:] {
		if isLengthExcessive(arg, 100) || !isNameValid(arg) {
			err = fmt.Errorf("Invalid command. Usage: command [%s]\n", arg)
			lengthFlag = false
			break
		}
	}
	if lengthFlag == false { //通過測驗就執行命令
		return msg, err
	}

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
			if args[2] != "--sort-name" && args[2] != "--sort-created" {
				return "", fmt.Errorf("Usage: list files [username] [foldername] [--sort-name|--sort-created] [asc|desc]\n")
			}
			parma.SortCriteria = args[2]
		}
		if len(args) >= 4 {
			if args[3] != "asc" && args[3] != "desc" {
				return "", fmt.Errorf("Usage: list files [username] [foldername] [--sort-name|--sort-created] [asc|desc]\n")
			}
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
			if args[3] != "--sort-name" && args[3] != "--sort-created" {
				return "", fmt.Errorf("Usage: list files [username] [foldername] [--sort-name|--sort-created] [asc|desc]\n")
			}
			parma.SortCriteria = args[3]
		}
		if len(args) >= 5 {
			if args[4] != "asc" && args[4] != "desc" {
				return "", fmt.Errorf("Usage: list files [username] [foldername] [--sort-name|--sort-created] [asc|desc]\n")
			}
			parma.Sortby = args[4]
		}
		msg, err = fs.listFiles(parma)
	default:
		err = fmt.Errorf("Error: Unrecognized command [%s]\n", args[0])
	}
	if err != nil {
		return "", err
	}
	return msg, nil
}
