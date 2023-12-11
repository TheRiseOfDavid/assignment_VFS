package user

import (
	"VFS/folder"
	objectbaseinfo "VFS/object_base_info"
	"fmt"
)

type User struct {
	objectbaseinfo.ObjectBaseInfo
	Folders []folder.Folder
}

func (owner *User) SelectFolder(folderName string) (*folder.Folder, error) {
	for i, folder := range owner.Folders {
		if folder.Name == folderName {
			return &owner.Folders[i], nil
		}
	}
	return nil, fmt.Errorf("Error: The [%s] doesn't exist.\n", folderName)
}

func (owner *User) IsFolderExists(folderName string) bool {
	for _, folder := range owner.Folders {
		if folder.Name == folderName {
			return true
		}
	}
	return false
}
