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
		for _, arg := range args[1:] {
			if isLengthExcessive(arg, 100) || !isNameValid(arg) {
				err := fmt.Errorf("Invalid command. Usage: command [%s]\n", arg)
				fmt.Fprint(os.Stderr, err)
				lengthFlag = false
				break
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
			err = fmt.Errorf("Invalid command. Usage: command [%s]\n", args[0])
		}
		if err != nil {
			fmt.Fprint(os.Stdout, err)
			continue
		}
		fmt.Fprint(os.Stdout, msg)
	}
}
