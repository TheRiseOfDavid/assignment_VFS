package main

type FolderOptions struct {
	userName      string
	folderName    string
	description   string
	sortCriteria  string
	sortby        string
	newFolderName string
}

type FileOptions struct {
	FolderOptions
	fileName string
}
