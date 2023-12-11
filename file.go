package main

// import (
// 	"fmt"
// 	"sort"
// 	"time"
// )

// type File struct {
// 	name        string
// 	description string
// 	created     string
// }

// type FileOptions struct {
// 	FolderOptions //繼承 folder 所需要的東西
// 	fileName      string
// }

// // GO 似乎不支持 指標陣列多型使用
// // https://stackoverflow.com/questions/54377597/how-to-make-a-function-that-receives-an-array-of-custom-interface

// func (fs *VirtualFileSystem) createFile(fileParma FileOptions) (string, error) {
// 	/* 會用到的
// 	type FileOptions struct {
// 		userName      string
// 		folderName    string
// 		fileName      string
// 	}
// 	*/
// 	if !isNameValid(fileParma.fileName) { //Error: The [filename] contains invalid chars
// 		return "", fmt.Errorf("Error: The [%s] contains invalid chars.\n", fileParma.fileName)
// 	}

// 	owner, err := fs.selectUser(fileParma.userName) //Error: The [username] doesn't exist
// 	if err != nil {
// 		return "", err
// 	}
// 	folder, err := owner.selectFolder(fileParma.folderName)
// 	if err != nil {
// 		return "", err
// 	}

// 	folder.files = append(folder.files, File{name: fileParma.fileName, description: fileParma.description,
// 		created: time.Now().Format("2006-01-02 15:04:05")})
// 	return fmt.Sprintf("Create [%s] in [%s] / [%s] successfull.\n",
// 		fileParma.fileName, fileParma.userName, fileParma.folderName), nil
// }

// func (fs *VirtualFileSystem) deleteFile(fileParma FileOptions) (string, error) {
// 	owner, err := fs.selectUser(fileParma.userName) //Error: The [username] doesn't exist
// 	if err != nil {
// 		return "", err
// 	}
// 	folder, err := owner.selectFolder(fileParma.folderName)
// 	if err != nil {
// 		return "", err
// 	}

// 	for i, file := range folder.files {
// 		if file.name == fileParma.folderName {
// 			folder.files = append(folder.files[:i], folder.files[i+1:]...)
// 			return fmt.Sprintf("Delete [%s] successfully.\n", fileParma.folderName), nil
// 		}
// 	}

// 	return "", fmt.Errorf("Error: The [%s] doesn't exist\n", fileParma.folderName)
// }

// func (fs *VirtualFileSystem) listFiles(fileParma FileOptions) (string, error) {
// 	/* 會用到的資訊
// 	type FolderOptions struct {
// 		fileName     string
// 		folderName   string
// 		userName     string
// 		sortCriteria string
// 		sortby       string
// 	}
// 	*/

// 	owner, err := fs.selectUser(fileParma.userName) //Error: The [username] doesn't exist
// 	if err != nil {
// 		return "", err
// 	}
// 	folder, err := owner.selectFolder(fileParma.folderName)
// 	if err != nil {
// 		return "", err
// 	}

// 	if len(folder.files) == 0 {
// 		return "", fmt.Errorf("Warning: The [%s] doesn't have any files.\n", fileParma.userName)
// 	}

// 	if fileParma.sortCriteria == "--sort-name" {
// 		sort.Slice(folder.files, func(i int, j int) bool {
// 			if fileParma.sortby == "asc" {
// 				return folder.files[i].name < folder.files[j].name
// 			} else {
// 				return folder.files[i].name > folder.files[j].name
// 			}
// 		})
// 	}
// 	if fileParma.sortCriteria == "--sort-created" {
// 		sort.Slice(folder.files, func(i int, j int) bool {
// 			if fileParma.sortby == "asc" {
// 				return folder.files[i].created < folder.files[j].created
// 			} else {
// 				return folder.files[i].created > folder.files[j].created
// 			}
// 		})
// 	}

// 	result := ""
// 	for _, folder := range folder.files {
// 		result += fmt.Sprintf("[%s] [%s] [%s] [%s] \n", folder.name, folder.description, folder.created, fileParma.userName)
// 	}
// 	return result, nil
// }
