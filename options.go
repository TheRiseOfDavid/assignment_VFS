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
	FolderOptions //繼承 folder 所需要的東西
	fileName      string
}
