package folder

import (
	file "virtual_file_system/file"
	objectmanager "virtual_file_system/object_manager"
)

type Folder struct {
	*objectmanager.ObjectManager[file.File]
	files []file.File
}
