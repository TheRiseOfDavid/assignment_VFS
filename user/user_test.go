package user

import (
	"VFS/file"
	"VFS/folder"
	objectbaseinfo "VFS/object_base_info"
	"reflect"
	"testing"
)

func TestSelectFolderForSuccess(t *testing.T) {
	fs := User{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder",
	},
		Folders: make([]folder.Folder, 0)}

	fs.Folders = append(fs.Folders, folder.Folder{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder1",
	},
		Files: make([]file.File, 0)})

	fs.Folders = append(fs.Folders, folder.Folder{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder2",
	},
		Files: make([]file.File, 0)})

	theFolder, _ := fs.SelectFolder("folder1")
	output := reflect.TypeOf(*theFolder)
	expected := reflect.TypeOf(folder.Folder{})
	if output != expected {
		t.Errorf("TestSelectUserForSuccess \nreturned %s\nexpected %s", output, expected)
	}
}

func TestSelectFolderForError(t *testing.T) {
	fs := User{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder",
	},
		Folders: make([]folder.Folder, 0)}

	fs.Folders = append(fs.Folders, folder.Folder{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder1",
	},
		Files: make([]file.File, 0)})

	fs.Folders = append(fs.Folders, folder.Folder{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder2",
	},
		Files: make([]file.File, 0)})

	_, err := fs.SelectFolder("folder3")
	expected := "Error: The [folder3] doesn't exist.\n"
	if err.Error() != expected {
		t.Errorf("TestSelectUserForError \nreturned %s\nexpected %s", err.Error(), expected)
	}
}

func TestIsFolderExistsForTrue(t *testing.T) {
	fs := User{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder",
	},
		Folders: make([]folder.Folder, 0)}

	fs.Folders = append(fs.Folders, folder.Folder{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder1",
	},
		Files: make([]file.File, 0)})

	fs.Folders = append(fs.Folders, folder.Folder{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder2",
	},
		Files: make([]file.File, 0)})

	output := fs.IsFolderExists("folder1")
	expected := true
	if output != expected {
		t.Errorf("TestSelectUserForSuccess \nreturned %t\nexpected %t", output, expected)
	}
}

func TestIsFolderExistsForFalse(t *testing.T) {
	fs := User{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder",
	},
		Folders: make([]folder.Folder, 0)}

	fs.Folders = append(fs.Folders, folder.Folder{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder1",
	},
		Files: make([]file.File, 0)})

	fs.Folders = append(fs.Folders, folder.Folder{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder2",
	},
		Files: make([]file.File, 0)})

	output := fs.IsFolderExists("folder3")
	expected := false
	if output != expected {
		t.Errorf("TestSelectUserForSuccess \nreturned %t\nexpected %t", output, expected)
	}
}
