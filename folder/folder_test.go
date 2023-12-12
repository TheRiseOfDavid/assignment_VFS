package folder

import (
	"VFS/file"
	objectbaseinfo "VFS/object_base_info"
	"reflect"
	"testing"
)

func TestSelectUserForSuccess(t *testing.T) {
	fs := Folder{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder",
	},
		Files: make([]file.File, 0)}

	fs.Files = append(fs.Files, file.File{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "file1",
	}})
	fs.Files = append(fs.Files, file.File{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "file2",
	}})

	theFile, _ := fs.SelectFile("file1")
	output := reflect.TypeOf(*theFile)
	expected := reflect.TypeOf(file.File{})
	if output != expected {
		t.Errorf("TestSelectUserForSuccess \nreturned %s\nexpected %s", output, expected)
	}
}

func TestSelectFileForError(t *testing.T) {
	fs := Folder{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder",
	},
		Files: make([]file.File, 0)}

	fs.Files = append(fs.Files, file.File{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "file1",
	}})
	fs.Files = append(fs.Files, file.File{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "file2",
	}})

	_, err := fs.SelectFile("file3")
	expected := "Error: The [file3] doesn't exist.\n"
	if err.Error() != expected {
		t.Errorf("TestSelectUserForError \nreturned %s\nexpected %s", err.Error(), expected)
	}
}

func TestIsFileExistsForTrue(t *testing.T) {
	fs := Folder{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder",
	},
		Files: make([]file.File, 0)}

	fs.Files = append(fs.Files, file.File{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "file1",
	}})
	fs.Files = append(fs.Files, file.File{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "file2",
	}})

	output := fs.IsFileExists("file1")
	expected := true
	if output != expected {
		t.Errorf("TestSelectUserForSuccess \nreturned %t\nexpected %t", output, expected)
	}
}

func TestIsFileExistsForFalse(t *testing.T) {
	fs := Folder{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "folder",
	},
		Files: make([]file.File, 0)}

	fs.Files = append(fs.Files, file.File{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "file1",
	}})
	fs.Files = append(fs.Files, file.File{ObjectBaseInfo: objectbaseinfo.ObjectBaseInfo{
		Name: "file2",
	}})

	output := fs.IsFileExists("file3")
	expected := false
	if output != expected {
		t.Errorf("TestSelectUserForSuccess \nreturned %t\nexpected %t", output, expected)
	}
}
