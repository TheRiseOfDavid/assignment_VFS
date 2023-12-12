package virtualfilesystem

import (
	"VFS/file"
	objectbaseinfo "VFS/object_base_info"
	"VFS/options"
	"fmt"
	"sort"
	"time"
)

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

	owner, err := fs.selectUser(fileParma.UserName) //Error: The [username] doesn't exist
	if err != nil {
		return "", err
	}
	folder, err := owner.SelectFolder(fileParma.FolderName)
	if err != nil {
		return "", err
	}
	if folder.IsFileExists(fileParma.FileName) {
		return "", fmt.Errorf("Error: The [%s] has already existed.\n", fileParma.FileName)
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
		if file.Name == fileParma.FileName {
			folder.Files = append(folder.Files[:i], folder.Files[i+1:]...)
			return fmt.Sprintf("Delete [%s] successfully.\n", fileParma.FileName), nil
		}
	}

	return "", fmt.Errorf("Error: The [%s] doesn't exist.\n", fileParma.FileName)
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
		return "", fmt.Errorf("Warning: The [%s] is empty.\n", fileParma.FolderName)
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
		result += fmt.Sprintf("[%s] [%s] [%s] [%s] [%s]\n", folder.Name, folder.Description, folder.Created, fileParma.FolderName, fileParma.UserName)
	}
	return result, nil
}
