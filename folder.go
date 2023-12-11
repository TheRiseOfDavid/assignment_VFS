package main

// import (
// 	"fmt"
// 	"io"
// 	"sort"
// 	"strings"
// 	"time"
// )

// type Folder struct {
// 	name        string
// 	description string
// 	created     string
// 	files       []File
// }

// type FolderOptions struct {
// 	userName      string
// 	folderName    string
// 	description   string
// 	sortCriteria  string
// 	sortby        string
// 	newFolderName string
// }

// func (f Folder) getName() string {
// 	return f.name
// }

// func (f *Folder) selectFile(fileName string) (*File, error) {
// 	for i, file := range f.files {
// 		if file.name == fileName {
// 			return &f.files[i], nil
// 		}
// 	}
// 	return nil, fmt.Errorf("Error: The [%s] doesn't exist.\n", fileName)
// }

// func (fs *VirtualFileSystem) createFolder(folderParma FolderOptions) (string, error) {
// 	if !isNameValid(folderParma.folderName) {
// 		return "", fmt.Errorf("Error: The [%s] contain invalid chars\n", folderParma.folderName)
// 	}

// 	folderParma.userName = strings.ToLower(folderParma.userName)     //進來程式後都小寫
// 	folderParma.folderName = strings.ToLower(folderParma.folderName) //進來程式後都小寫

// 	owner, err := fs.selectUser(folderParma.userName)
// 	if err != nil {
// 		return "", err // err 回傳 fmt.Errorf("Error: : The [%s] doesn't exist\n", userName)
// 	}

// 	//fmt.Print("The len of folders is ", len(owner.folders), "\n")
// 	/*
// 		for _, folder := range owner.folders {
// 			if folder.name == folderParma.folderName {
// 				fmt.Print(folderParma.folderName, "\n")
// 			}
// 		}
// 	*/

// 	if owner.isFolderExists(folderParma.folderName) {
// 		return "", fmt.Errorf("Error: [%s] has already existed\n", folderParma.folderName)
// 	}

// 	folder := Folder{name: folderParma.folderName, description: folderParma.description,
// 		created: time.Now().Format("2006-01-02 15:04:05"), files: make([]File, 0)}
// 	owner.folders = append(owner.folders, folder)
// 	//fmt.Print("update The len of folders is ", len(owner.folders), "\n")
// 	return fmt.Sprintf("Create [%s] successfully. \n", folderParma.folderName), nil
// }

// func (fs *VirtualFileSystem) deleteFolder(folderParma FolderOptions) (string, error) {
// 	/* 會用到的資訊
// 	type FolderOptions struct {
// 		userName    string
// 		folderName  string
// 	}
// 	*/

// 	folderParma.userName = strings.ToLower(folderParma.userName)     //進來程式後都小寫
// 	folderParma.folderName = strings.ToLower(folderParma.folderName) //進來程式後都小寫

// 	owner, err := fs.selectUser(folderParma.userName)
// 	if err != nil {
// 		return "", err // err 回傳 fmt.Errorf("Error: : The [%s] doesn't exist\n", userName)
// 	}

// 	for i, folder := range owner.folders {
// 		if folder.name == folderParma.folderName {
// 			owner.folders = append(owner.folders[:i], owner.folders[i+1:]...)
// 			return fmt.Sprintf("Delete [%s] successfully.\n", folderParma.folderName), nil
// 		}
// 	}

// 	return "", fmt.Errorf("Error: The [%s] doesn't exist\n", folderParma.folderName)
// }

// func (fs *VirtualFileSystem) listFolders(folderParma FolderOptions) (string, error) {
// 	/* 會用到的資訊
// 	type FolderOptions struct {
// 		userName     string
// 		sortCriteria string
// 		sortby       string
// 	}
// 	*/

// 	folderParma.userName = strings.ToLower(folderParma.userName)     //進來程式後都小寫
// 	folderParma.folderName = strings.ToLower(folderParma.folderName) //進來程式後都小寫

// 	owner, err := fs.selectUser(folderParma.userName)
// 	if err != nil {
// 		return "", err
// 	}

// 	if len(owner.folders) == 0 {
// 		return "", fmt.Errorf("Warning: The [%s] doesn't have any folders.\n", folderParma.userName)
// 	}

// 	if folderParma.sortCriteria == "--sort-name" {
// 		sort.Slice(owner.folders, func(i int, j int) bool {
// 			if folderParma.sortby == "asc" {
// 				return owner.folders[i].name < owner.folders[j].name
// 			} else {
// 				return owner.folders[i].name > owner.folders[j].name
// 			}
// 		})
// 	}
// 	if folderParma.sortCriteria == "--sort-created" {
// 		sort.Slice(owner.folders, func(i int, j int) bool {
// 			if folderParma.sortby == "asc" {
// 				return owner.folders[i].created < owner.folders[j].created
// 			} else {
// 				return owner.folders[i].created > owner.folders[j].created
// 			}
// 		})
// 	}

// 	result := ""
// 	for _, folder := range owner.folders {
// 		result += fmt.Sprintf("[%s] [%s] [%s] [%s] \n", folder.name, folder.description, folder.created, folderParma.userName)
// 	}
// 	return result, nil
// }

// func (fs *VirtualFileSystem) renameFolder(folderParma FolderOptions) (string, error) {
// 	//會用到的資訊
// 	/*
// 		type FolderOptions struct {
// 			userName      string
// 			folderName    string
// 			newFolderName string
// 		}
// 	*/

// 	folderParma.userName = strings.ToLower(folderParma.userName)     //進來程式後都小寫
// 	folderParma.folderName = strings.ToLower(folderParma.folderName) //進來程式後都小寫

// 	owner, err := fs.selectUser(folderParma.userName)
// 	if err != nil {
// 		return "", err
// 	}

// 	for i, folder := range owner.folders {
// 		if folder.name == folderParma.folderName {
// 			owner.folders[i].name = folderParma.newFolderName
// 			return fmt.Sprintf("Rename [%s] to [%s] successfully.\n",
// 				folderParma.folderName, folderParma.newFolderName), nil
// 		}
// 	}
// 	return "", fmt.Errorf("Error: The [%s] doesn't exist\n", folderParma.folderName)
// }

// type UnitTestOptions struct {
// 	isUnitTest int
// 	reader     io.Reader
// }
