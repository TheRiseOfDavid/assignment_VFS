package virtualfilesystem

import (
	"VFS/folder"
	objectbaseinfo "VFS/object_base_info"
	"VFS/user"
	"fmt"
)

func (fs *VirtualFileSystem) isUserExists(userName string) bool {
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
	return nil, fmt.Errorf("Error: The [%s] doesn't exist.\n", userName)
}

func (fs *VirtualFileSystem) registerUser(userName string) (string, error) {
	//需要再補一個 Error: The [username] contain invalid chars.
	if !isNameValid(userName) {
		return "", fmt.Errorf("Error: The [%s] contain invalid chars.\n", userName)
	}
	if fs.isUserExists(userName) {
		return "", fmt.Errorf("Error: The [%s] has already existed.\n", userName)
	}

	owner := &user.User{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{Name: userName},
		Folders: make([]folder.Folder, 0)}
	fs.owners = append(fs.owners, *owner)

	return fmt.Sprintf("Add [%s] successfully.\n", owner.Name), nil
}
