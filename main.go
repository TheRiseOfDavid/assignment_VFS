package main

import (
	"VFS/virtualfilesystem"
)

func main() {
	fs := virtualfilesystem.CreateVirtaulFileSystem()
	fs.CommandShell()
}
