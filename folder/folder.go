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
