package options

type FolderOptions struct {
	UserName      string
	FolderName    string
	Description   string
	SortCriteria  string
	Sortby        string
	NewFolderName string
}

type FileOptions struct {
	FolderOptions //繼承 folder 所需要的東西
	FileName      string
}
