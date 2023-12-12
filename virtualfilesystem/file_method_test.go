package virtualfilesystem

import (
	"VFS/options"
	"strings"
	"testing"
	"time"
)

func TestCreateFileForSuccess(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	output, _ := fs.createFolder(folderParma)

	fileParma := options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file",
	}
	output, _ = fs.createFile(fileParma)
	expected := "Create [file] in [david] / [folder] successfull.\n"
	if output != expected {
		t.Errorf("TestCreateFileForSuccess \nreturned %v\nexpected %s",
			output, expected)
	}
}

func TestCreateFileForSuccessWithDescription(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	output, _ := fs.createFolder(folderParma)

	fileParma := options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:    "david",
			FolderName:  "folder",
			Description: "This is a file.",
		},
		FileName: "file",
	}
	output, _ = fs.createFile(fileParma)
	expected := "Create [file] in [david] / [folder] successfull.\n"
	if output != expected {
		t.Errorf("TestCreateFileForSuccess \nreturned %v\nexpected %s",
			output, expected)
	}
}

func TestCreateFileForErrorWithUserNotFound(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	_, _ = fs.createFolder(folderParma)

	fileParma := options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david2",
			FolderName: "folder",
		},
		FileName: "file",
	}
	_, err := fs.createFile(fileParma)
	expected := "Error: The [david2] doesn't exist.\n"
	if err.Error() != expected {
		t.Errorf("TestCreateFileForErrorWithUserNotFound \nreturned %v\nexpected %s",
			err.Error(), expected)
	}
}

func TestCreateFileForErrorWithFolderNotFound(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	_, _ = fs.createFolder(folderParma)

	fileParma := options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder2",
		},
		FileName: "file",
	}
	_, err := fs.createFile(fileParma)
	expected := "Error: The [folder2] doesn't exist.\n"
	if err.Error() != expected {
		t.Errorf("TestCreateFileForErrorWithFolderNotFound \nreturned %v\nexpected %s",
			err.Error(), expected)
	}
}

func TestDeleteFileForSuccess(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	output, _ := fs.createFolder(folderParma)

	fileParma := options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file",
	}
	output, _ = fs.createFile(fileParma)

	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file",
	}
	output, _ = fs.deleteFile(fileParma)
	expected := "Delete [file] successfully.\n"
	if output != expected {
		t.Errorf("TestDeleteFileForSuccess \nreturned %v\nexpected %s",
			output, expected)
	}
}

func TestDeleteFileForErrorWithUserNotFound(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	_, _ = fs.createFolder(folderParma)

	fileParma := options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file",
	}
	_, _ = fs.createFile(fileParma)

	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david2",
			FolderName: "folder",
		},
		FileName: "file",
	}
	_, err := fs.deleteFile(fileParma)
	expected := "Error: The [david2] doesn't exist.\n"
	if err.Error() != expected {
		t.Errorf("TestDeleteFileForErrorWithUserNotFound \nreturned %v\nexpected %s",
			err.Error(), expected)
	}
}

func TestDeleteFileForErrorWithFolderNotFound(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	_, _ = fs.createFolder(folderParma)

	fileParma := options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file",
	}
	_, _ = fs.createFile(fileParma)

	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder2",
		},
		FileName: "file",
	}
	_, err := fs.deleteFile(fileParma)
	expected := "Error: The [folder2] doesn't exist.\n"
	if err.Error() != expected {
		t.Errorf("TestDeleteFileForErrorWithFolderNotFound \nreturned %v\nexpected %s",
			err.Error(), expected)
	}
}

func TestDeleteFileForErrorWithFileNotFound(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	_, _ = fs.createFolder(folderParma)

	fileParma := options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file",
	}
	_, _ = fs.createFile(fileParma)

	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file2",
	}
	_, err := fs.deleteFile(fileParma)
	expected := "Error: The [file2] doesn't exist.\n"
	if err.Error() != expected {
		t.Errorf("TestDeleteFileForErrorWithFileNotFound \nreturned %v\nexpected %s",
			err.Error(), expected)
	}
}

func TestListFilesForSortNamebyAsc(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:    "david",
		FolderName:  "folder",
		Description: "This is a description.",
	}
	output, _ := fs.createFolder(folderParma)
	fileParma := options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file1",
	}
	output, _ = fs.createFile(fileParma)
	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file2",
	}
	output, _ = fs.createFile(fileParma)
	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file3",
	}
	output, _ = fs.createFile(fileParma)

	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:     "david",
			FolderName:   "folder",
			SortCriteria: "--sort-name",
			Sortby:       "asc",
		},
	}
	output, _ = fs.listFiles(fileParma)
	outputs := strings.Split(output, "\n")
	expected := [...]string{"[file1]", "[file2]", "[file3]"}
	j := 0
	//fmt.Print(output)
	for i := 0; i <= 2; i++ {
		result := strings.Split(outputs[i], " ")
		if result[0] != expected[j] {
			t.Errorf("TestListFilesForSortNamebyAsc Failed %s, %s\n",
				result[0], expected[j])
		}
		j += 1
	}
}

func TestListFilesForSortNamebyDesc(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:    "david",
		FolderName:  "folder",
		Description: "This is a description.",
	}
	output, _ := fs.createFolder(folderParma)
	fileParma := options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file1",
	}
	output, _ = fs.createFile(fileParma)
	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file2",
	}
	output, _ = fs.createFile(fileParma)
	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file3",
	}
	output, _ = fs.createFile(fileParma)

	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:     "david",
			FolderName:   "folder",
			SortCriteria: "--sort-name",
			Sortby:       "desc",
		},
	}
	output, _ = fs.listFiles(fileParma)
	outputs := strings.Split(output, "\n")
	expected := [...]string{"[file3]", "[file2]", "[file1]"}
	j := 0
	//fmt.Print(output)
	for i := 0; i <= 2; i++ {
		result := strings.Split(outputs[i], " ")
		if result[0] != expected[j] {
			t.Errorf("TestListFilesForSortNamebyAsc Failed %s, %s\n",
				result[0], expected[j])
		}
		j += 1
	}
}

func TestListFilesForSortCreatedbyAsc(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:    "david",
		FolderName:  "folder",
		Description: "This is a description.",
	}
	output, _ := fs.createFolder(folderParma)
	fileParma := options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file2",
	}
	time.Sleep(1 * time.Second)
	output, _ = fs.createFile(fileParma)
	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file3",
	}
	time.Sleep(1 * time.Second)
	output, _ = fs.createFile(fileParma)
	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file1",
	}
	output, _ = fs.createFile(fileParma)

	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:     "david",
			FolderName:   "folder",
			SortCriteria: "--sort-created",
			Sortby:       "asc",
		},
	}
	output, _ = fs.listFiles(fileParma)
	outputs := strings.Split(output, "\n")
	expected := [...]string{"[file2]", "[file3]", "[file1]"}
	j := 0
	//fmt.Print(output)
	for i := 0; i <= 2; i++ {
		result := strings.Split(outputs[i], " ")
		if result[0] != expected[j] {
			t.Errorf("TestListFilesForSortNamebyAsc Failed %s, %s\n",
				result[0], expected[j])
		}
		j += 1
	}
}

func TestListFilesForSortCreatedbyDesc(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:    "david",
		FolderName:  "folder",
		Description: "This is a description.",
	}
	output, _ := fs.createFolder(folderParma)
	fileParma := options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file2",
	}

	output, _ = fs.createFile(fileParma)
	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file3",
	}
	time.Sleep(1 * time.Second)
	output, _ = fs.createFile(fileParma)
	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file1",
	}
	time.Sleep(1 * time.Second)
	output, _ = fs.createFile(fileParma)

	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:     "david",
			FolderName:   "folder",
			SortCriteria: "--sort-created",
			Sortby:       "desc",
		},
	}
	output, _ = fs.listFiles(fileParma)
	outputs := strings.Split(output, "\n")
	expected := [...]string{"[file1]", "[file3]", "[file2]"}
	j := 0
	//fmt.Print(output)
	for i := 0; i <= 2; i++ {
		result := strings.Split(outputs[i], " ")
		if result[0] != expected[j] {
			t.Errorf("TestListFilesForSortNamebyAsc Failed %s, %s\n",
				result[0], expected[j])
		}
		j += 1
	}
}

func TestListFileForWarningWithUserNotFound(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:    "david",
		FolderName:  "folder",
		Description: "This is a description.",
	}
	_, _ = fs.createFolder(folderParma)
	fileParma := options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file1",
	}
	_, _ = fs.createFile(fileParma)

	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:     "david2",
			FolderName:   "folder",
			SortCriteria: "--sort-created",
			Sortby:       "desc",
		},
	}
	_, err := fs.listFiles(fileParma)

	expected := "Error: The [david2] doesn't exist.\n"
	if err.Error() != expected {
		t.Errorf("TestListFoldersForWarningWithUserNotFound return %sexpected %s",
			err.Error(), expected)
	}
}

func TestListFileForWarningWithFolderNotFound(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:    "david",
		FolderName:  "folder",
		Description: "This is a description.",
	}
	_, _ = fs.createFolder(folderParma)
	fileParma := options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:   "david",
			FolderName: "folder",
		},
		FileName: "file1",
	}
	_, _ = fs.createFile(fileParma)

	fileParma = options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:     "david",
			FolderName:   "folder2",
			SortCriteria: "--sort-created",
			Sortby:       "desc",
		},
	}
	_, err := fs.listFiles(fileParma)

	expected := "Error: The [folder2] doesn't exist.\n"
	if err.Error() != expected {
		t.Errorf("TestListFileForWarningWithFolderNotFound return %sexpected %s",
			err.Error(), expected)
	}
}

func TestListFileForWarningWithoutFile(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:    "david",
		FolderName:  "folder",
		Description: "This is a description.",
	}
	_, _ = fs.createFolder(folderParma)

	fileParma := options.FileOptions{
		FolderOptions: options.FolderOptions{
			UserName:     "david",
			FolderName:   "folder",
			SortCriteria: "--sort-created",
			Sortby:       "desc",
		},
	}
	_, err := fs.listFiles(fileParma)

	expected := "Warning: The [folder] is empty.\n"
	if err.Error() != expected {
		t.Errorf("TestListFileForWarningWithFolderNotFound return %sexpected %s",
			err.Error(), expected)
	}
}
