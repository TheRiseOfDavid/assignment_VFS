package user

import (
	folder "virtual_file_system/folder"
	objectmanager "virtual_file_system/object_manager"
)

type User struct {
	*objectmanager.ObjectManager[folder.Folder]
	Folders []folder.Folder
}
