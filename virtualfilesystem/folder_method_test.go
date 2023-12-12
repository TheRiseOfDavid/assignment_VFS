package virtualfilesystem

import (
	"VFS/options"
	"strings"
	"testing"
	"time"
)

func TestCreateFolderWithoutDescriptionForSuccess(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	output, _ := fs.createFolder(folderParma)

	expected := "Create [folder] successfully.\n"
	if output != expected {
		t.Errorf("CreateFolderWithoutDescriptionForSuccess \nreturned %v\nexpected %s",
			output, expected)
	}

}

func TestCreateFolderWithDescriptionForSuccess(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:    "david",
		FolderName:  "folder",
		Description: "This is a description.",
	}
	output, _ := fs.createFolder(folderParma)

	expected := "Create [folder] successfully.\n"
	if output != expected {
		t.Errorf("CreateFolderWithDescriptionForSuccess \nreturned %v\nexpected %s",
			output, expected)
	}

}

func TestCreateFolderForErrorWithInvalidChars(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder大",
	}
	_, err := fs.createFolder(folderParma)

	expected := "Error: The [folder大] contain invalid chars\n"
	if err.Error() != expected {
		t.Errorf("CreateFolderForErrorWithInvalidChars \nreturned %v\nexpected %s",
			err.Error(), expected)
	}

}

func TestCreateFolderForErrorWithUserNotFound(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	_, _ = fs.createFolder(folderParma)
	folderParma = options.FolderOptions{
		UserName:   "david2",
		FolderName: "folder",
	}
	_, err := fs.createFolder(folderParma)

	expected := "Error: The [david2] doesn't exist.\n"
	if err.Error() != expected {
		t.Errorf("CreateFolderForErrorWithUserNotFound \nreturned %v\nexpected %s",
			err.Error(), expected)
	}
}

func TestCreateFolderForErrorWithFolderExisted(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	_, _ = fs.createFolder(folderParma)
	folderParma = options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	_, err := fs.createFolder(folderParma)

	expected := "Error: The [folder] has already existed.\n"
	if err.Error() != expected {
		t.Errorf("CreateFolderForErrorWithFolderExisted \nreturned %v\nexpected %s",
			err.Error(), expected)
	}
}

func TestDeleteFolderForSuccess(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:    "david",
		FolderName:  "folder",
		Description: "This is a description.",
	}
	output, _ := fs.createFolder(folderParma)

	folderParma = options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	output, _ = fs.deleteFolder(folderParma)
	expected := "Delete [folder] successfully.\n"
	if output != expected {
		t.Errorf("TestDeleteFolderForSuccess \nreturned %v\nexpected %s",
			output, expected)
	}

}

func TestDeleteFolderForErrorWithUserNotFound(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:    "david",
		FolderName:  "folder",
		Description: "This is a description.",
	}
	_, _ = fs.createFolder(folderParma)

	folderParma = options.FolderOptions{
		UserName:   "david2",
		FolderName: "folder",
	}
	_, err := fs.deleteFolder(folderParma)
	expected := "Error: The [david2] doesn't exist.\n"
	if err.Error() != expected {
		t.Errorf("TestDeleteFolderForErrorWithNotFound \nreturned %v\nexpected %s",
			err.Error(), expected)
	}
}

func TestDeleteFolderForErrorWithFolderNotFound(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:    "david",
		FolderName:  "folder",
		Description: "This is a description.",
	}
	_, _ = fs.createFolder(folderParma)

	folderParma = options.FolderOptions{
		UserName:   "david",
		FolderName: "folder2",
	}
	_, err := fs.deleteFolder(folderParma)
	expected := "Error: The [folder2] doesn't exist.\n"
	if err.Error() != expected {
		t.Errorf("DeleteFolderForErrorWithFolderNotFound \nreturned %v\nexpected %s",
			err.Error(), expected)
	}
}

func TestListFoldersForSortNamebyAsc(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:    "david",
		FolderName:  "folder1",
		Description: "This is a description.",
	}
	output, _ := fs.createFolder(folderParma)
	folderParma = options.FolderOptions{
		UserName:   "david",
		FolderName: "folder2",
	}
	output, _ = fs.createFolder(folderParma)
	folderParma = options.FolderOptions{
		UserName:   "david",
		FolderName: "folder3",
	}
	output, _ = fs.createFolder(folderParma)

	folderParma = options.FolderOptions{
		UserName:     "david",
		SortCriteria: "--sort-name",
		Sortby:       "asc",
	}
	output, _ = fs.listFolders(folderParma)
	outputs := strings.Split(output, "\n")
	expected := [...]string{"[folder1]", "[folder2]", "[folder3]"}
	j := 0
	for i := 0; i <= 2; i++ {
		result := strings.Split(outputs[i], " ")
		if result[0] != expected[j] {
			t.Errorf("TestListFoldersForSortNamebyAsc Failed %s, %s\n",
				result[0], expected[j])
		}
		j += 1
	}
}

func TestListFoldersForSortNamebyDesc(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:    "david",
		FolderName:  "folder1",
		Description: "This is a description.",
	}
	output, _ := fs.createFolder(folderParma)
	folderParma = options.FolderOptions{
		UserName:   "david",
		FolderName: "folder2",
	}
	output, _ = fs.createFolder(folderParma)
	folderParma = options.FolderOptions{
		UserName:   "david",
		FolderName: "folder3",
	}
	output, _ = fs.createFolder(folderParma)

	folderParma = options.FolderOptions{
		UserName:     "david",
		SortCriteria: "--sort-name",
		Sortby:       "desc",
	}
	output, _ = fs.listFolders(folderParma)
	outputs := strings.Split(output, "\n")
	expected := [...]string{"[folder3]", "[folder2]", "[folder1]"}
	j := 0
	for i := 0; i <= 2; i++ {
		result := strings.Split(outputs[i], " ")
		if result[0] != expected[j] {
			t.Errorf("TestListFoldersForSortNamebyAsc Failed %s, %s\n",
				result[0], expected[j])
		}
		j += 1
	}
}

func TestListFoldersForSortCreatedbyAsc(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:    "david",
		FolderName:  "folder2",
		Description: "This is a description.",
	}
	output, _ := fs.createFolder(folderParma)
	folderParma = options.FolderOptions{
		UserName:   "david",
		FolderName: "folder1",
	}
	output, _ = fs.createFolder(folderParma)
	folderParma = options.FolderOptions{
		UserName:   "david",
		FolderName: "folder3",
	}
	output, _ = fs.createFolder(folderParma)

	folderParma = options.FolderOptions{
		UserName:     "david",
		SortCriteria: "--sort-created",
		Sortby:       "asc",
	}
	output, _ = fs.listFolders(folderParma)
	outputs := strings.Split(output, "\n")
	expected := [...]string{"[folder2]", "[folder1]", "[folder3]"}
	j := 0
	for i := 0; i <= 2; i++ {
		result := strings.Split(outputs[i], " ")
		if result[0] != expected[j] {
			t.Errorf("TestListFoldersForSortNamebyAsc Failed %s, %s\n",
				result[0], expected[j])
		}
		j += 1
	}
}

func TestListFoldersForSortCreatedbyDesc(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder2",
	}

	output, _ := fs.createFolder(folderParma)
	folderParma = options.FolderOptions{
		UserName:   "david",
		FolderName: "folder1",
	}
	time.Sleep(1 * time.Second)
	output, _ = fs.createFolder(folderParma)
	folderParma = options.FolderOptions{
		UserName:   "david",
		FolderName: "folder3",
	}
	time.Sleep(1 * time.Second)
	output, _ = fs.createFolder(folderParma)
	time.Sleep(1 * time.Second)
	folderParma = options.FolderOptions{
		UserName:     "david",
		SortCriteria: "--sort-created",
		Sortby:       "desc",
	}
	output, _ = fs.listFolders(folderParma)

	outputs := strings.Split(output, "\n")
	expected := [...]string{"[folder3]", "[folder1]", "[folder2]"}
	j := 0
	for i := 0; i <= 2; i++ {
		result := strings.Split(outputs[i], " ")
		if result[0] != expected[j] {
			t.Errorf("TestListFoldersForSortNamebyAsc Failed %s, %s\n",
				result[0], expected[j])
		}
		j += 1
	}
}

func TestListFoldersForWarningWithoutFolder(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder2",
	}

	folderParma = options.FolderOptions{
		UserName:     "david",
		SortCriteria: "--sort-created",
		Sortby:       "desc",
	}
	_, err := fs.listFolders(folderParma)
	expected := "Warning: The [david] doesn't have any folders.\n"
	if err.Error() != expected {
		t.Errorf("TestListFoldersForWarningWithoutFolder return %sexpected %s",
			err.Error(), expected)
	}
}

func TestListFoldersForWarningWithUserNotFound(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder2",
	}

	folderParma = options.FolderOptions{
		UserName:     "david2",
		SortCriteria: "--sort-created",
		Sortby:       "desc",
	}
	_, err := fs.listFolders(folderParma)
	expected := "Error: The [david2] doesn't exist.\n"
	if err.Error() != expected {
		t.Errorf("TestListFoldersForWarningWithUserNotFound return %sexpected %s",
			err.Error(), expected)
	}

}

func TestRenameFolderForSuccess(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	output, _ := fs.createFolder(folderParma)

	folderParma = options.FolderOptions{
		UserName:      "david",
		FolderName:    "folder",
		NewFolderName: "newfolder",
	}
	output, _ = fs.renameFolder(folderParma)
	expected := "Rename [folder] to [newfolder] successfully.\n"
	if output != expected {
		t.Errorf("TestRenameFolderForSuccess \nreturned %v\nexpected %s",
			output, expected)
	}

}

func TestRenameFolderForErrorWithUserNotFound(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	_, _ = fs.createFolder(folderParma)

	folderParma = options.FolderOptions{
		UserName:      "david2",
		FolderName:    "folder",
		NewFolderName: "newfolder",
	}
	_, err := fs.renameFolder(folderParma)
	expected := "Error: The [david2] doesn't exist.\n"
	if err.Error() != expected {
		t.Errorf("RenameFolderForErrorWithUserNotFound \nreturned %v\nexpected %s",
			err.Error(), expected)
	}
}

func TestRenameFolderForErrorWithFolderNotFound(t *testing.T) {
	fs := CreateVirtaulFileSystem()
	_, _ = fs.registerUser("david")
	folderParma := options.FolderOptions{
		UserName:   "david",
		FolderName: "folder",
	}
	_, _ = fs.createFolder(folderParma)

	folderParma = options.FolderOptions{
		UserName:      "david",
		FolderName:    "folder2",
		NewFolderName: "newfolder",
	}
	_, err := fs.renameFolder(folderParma)
	expected := "Error: The [folder2] doesn't exist.\n"
	if err.Error() != expected {
		t.Errorf("RenameFolderForErrorWithFolderNotFound \nreturned %v\nexpected %s",
			err.Error(), expected)
	}
}
