package folder

import (
	"VFS/file"
	objectbaseinfo "VFS/object_base_info"
	"fmt"
)

type Folder struct {
	objectbaseinfo.ObjectBaseInfo
	Files []file.File
}

func (f *Folder) SelectFile(fileName string) (*file.File, error) {
	for i, file := range f.Files {
		if file.Name == fileName {
			return &f.Files[i], nil
		}
	}
	return nil, fmt.Errorf("Error: The [%s] doesn't exist.\n", fileName)
}

func (f *Folder) IsFileExists(folderName string) bool {
	for _, file := range f.Files {
		if file.Name == folderName {
			return true
		}
	}
	return false
}
