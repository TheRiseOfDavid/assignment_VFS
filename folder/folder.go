package folder

import (
	"fmt"
	"virtualfilesystem/file"
	objectbaseinfo "virtualfilesystem/object_base_info"
)

type Folder struct {
	objectbaseinfo.ObjectBaseInfo
	Files []file.File
}

func (f *Folder) selectFile(fileName string) (*file.File, error) {
	for i, file := range f.Files {
		if file.Name == fileName {
			return &f.Files[i], nil
		}
	}
	return nil, fmt.Errorf("Error: The [%s] doesn't exist.\n", fileName)
}
