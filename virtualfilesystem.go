package virtualfilesystem

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type User struct {
	name string
}

// VirtualFileSystem 出發點
type VirtualFileSystem struct {
	owners []User
}

func createVirtaulFileSystem() *VirtualFileSystem {
	return &VirtualFileSystem{owners: make([]User, 0)}
}

func (fs *VirtualFileSystem) registerUser(userName string) (string, error) {
	for _, owner := range fs.owners {
		if owner.name == userName {
			//fmt.Fprintf(os.Stdout, "Error: The %v has already existed.\n", owner.name)
			return fmt.Sprintf("Error: The %v has already existed.\n", owner.name),
				fmt.Errorf("Error: The %s has already existed.", owner.name)
		}
	}

	owner := User{name: userName}
	fs.owners = append(fs.owners, owner)

	//fmt.Fprintf(os.Stdout, "Add %v successfull. \n", owner.name)
	return fmt.Sprintf("Add %v successfull. \n", owner.name),
		fmt.Errorf("Add %s successfull. ", userName)
}

type UnitTestOptions struct {
	isUnitTest int
	reader     io.Reader
}

// 提供line interface 給 VFS
func (fs *VirtualFileSystem) commandShell(utOption UnitTestOptions) {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("vfs> ")
		scanner.Scan()
		command := scanner.Text()

		if command == "exit" {
			break
		}

		args := strings.Fields(command)
		if len(args) < 2 {
			fmt.Println("Invalid command. Usage: command [arguments]")
			continue
		}

		switch args[0] {
		case "register":
			msg, err := fs.registerUser(args[1])
			if err != nil {
				fmt.Fprint(os.Stdout, "Invalid command. Usage: command [arguments]")
			}
			fmt.Fprint(os.Stdout, msg)

		}
	}
}

func main() {
	fs := createVirtaulFileSystem()
	fs.commandShell()

}
