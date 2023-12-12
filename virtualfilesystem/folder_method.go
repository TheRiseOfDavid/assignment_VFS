package virtualfilesystem

import (
	"VFS/file"
	"VFS/folder"
	objectbaseinfo "VFS/object_base_info"
	"VFS/options"
	"fmt"
	"sort"
	"strings"
	"time"
)

func (fs *VirtualFileSystem) createFolder(folderParma options.FolderOptions) (string, error) {
	if !isNameValid(folderParma.FolderName) {
		return "", fmt.Errorf("Error: The [%s] contain invalid chars\n", folderParma.FolderName)
	}

	owner, err := fs.selectUser(folderParma.UserName)
	if err != nil {
		return "", err // err 回傳 fmt.Errorf("Error: : The [%s] doesn't exist\n", userName)
	}

	if owner.IsFolderExists(folderParma.FolderName) {
		return "", fmt.Errorf("Error: The [%s] has already existed.\n", folderParma.FolderName)
	}

	folder := folder.Folder{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: folderParma.FolderName, Description: folderParma.Description,
		Created: time.Now().Format("2006-01-02 15:04:05"),
	},
		Files: make([]file.File, 0)}
	owner.Folders = append(owner.Folders, folder)
	//fmt.Print("update The len of folders is ", len(owner.folders), "\n")
	return fmt.Sprintf("Create [%s] successfully.\n", folderParma.FolderName), nil
}

func (fs *VirtualFileSystem) deleteFolder(folderParma options.FolderOptions) (string, error) {
	/* 會用到的資訊
	type FolderOptions struct {
		userName    string
		folderName  string
	}
	*/

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

	return "", fmt.Errorf("Error: The [%s] doesn't exist.\n", folderParma.FolderName)
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

	owner, err := fs.selectUser(folderParma.UserName)
	if err != nil {
		return "", err
	}

	folder, err := owner.SelectFolder(folderParma.FolderName)
	if err != nil {
		return "", err
	}
	folder.Name = folderParma.NewFolderName
	return fmt.Sprintf("Rename [%s] to [%s] successfully.\n",
		folderParma.FolderName, folderParma.NewFolderName), nil

}
