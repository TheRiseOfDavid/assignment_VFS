# Virtual File System (VFS)

[TOC]

## Obejctive

- Implement a virtual file system with user and file management capabilities using GoLang 1.20+.

## Technical Requirements

- The program must be implemented using GoLang 1.20+

## Features

- User Management
  - Allow users to register a unique, case insensitive username.
  - Users can have an arbitrary number of folders and files.
- Folder Management
  - Users can create, delete, and rename folders.
  - Folder names must be unique within the user's scope and are case insensitive.
  - Folders have an optional description field.
- File Management

  - Users can create, delete, and list all files within a specified folder.
  - File names must be unique within the same folder and are case insensitive.
  - Files have an optional description field.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed GoLang 1.20+.

### Installation

1. Clone the repository:

```bash
git clone https://github.com/TheRiseOfDavid/assignment_VFS
```

2. Change into the project directory:

```bash
cd assignment_VFS
```

## Usage

**Note:**

- All the messages of successful or warning command executions should output to STDOUT.
- All the Error messages should output to STDERR.
- The token surrounded by [...] is a user input/variable.
- The question mark(?) within the [...]? indicates that token/user input is optional.

### Initialize the VFS

- Created the VFS

```bash
go run .
# VFS> //View the VFS command as it runs.
```

1.  User Registeration
    - `register[username]`
      Response
      - `# Add [username] successfully.``
      - `# Error: The [username] has already existed.`
      - `# Error: The [username] contain invalid chars.``
2.  Folder Management

    - `create-folder [username] [foldername] [description]?`
      Response
      - `# Create [foldername] successfully.`
      - `# Error: The [username] doesn\'t exist.`
      - `# Error: The [foldername] contain invalid chars.`
    - `delete-folder [username] [foldername]`
      Response

      - `Delete [foldername] successfully.`
      - `Error: The [username] doesn\'t exist.`
      - `Error: The [foldername] doesn\'t exist.`

    - `list-folders [username] [--sort-name|--sort-created] [asc|desc]`
      Response
      - `List all the folders within the [username] scope in following formats:[foldername] [description] [created at] [username]`
        Each field should be separated by whitespace or tab characters.
        The `[created at]` is a human-readable date/time format.
        The order of printed folder information is determined by the
        `--sort-name` or `--sort-created` combined with asc or desc flags.
        The `--sort-name` flag means sorting by `[foldername]`.
        If neither `--sort-name` nor ` --sort-created` is provided, sort the list by `[foldername]` in ascending order.
      - Warning: The [username] doesn't have any folders.
      - Error: The [username] doesn't exist.
      - Prompt the user the usage of the command if there is an invalid flag.(should output to STDERR)
    - `rename-folder [username] [foldername] [new-folder-name]`
      Response:
      - `Rename [foldername] to [new-folder-name] successfully.`
      - `Error: The [username] doesn't exist.`
      - `Error: The [foldername] doesn't exist`

3.  File Management
    - `create-file [username] [foldername] [filename] [description]?`
      Response:
      - `Create [filename] in [username] / [foldername] successfully`
      - `Error: The [username] doesn't exist.`
      - `Error: The [foldername] doesn't exist.`
      - `Error: The [filename] contains invalid chars.`
    - `delete-file [username] [foldername] [filename]`
      Response:
      - `Delete [filename] in [username] / [foldername] successfully.`
      - `Error: The [username] doesn't exist.`
      - `Error: The [foldername] doesn't exist.`
      - `Error: The [filename] doesn't exist.`
    - `list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]`
      Response: List files with the following fields: `[filename] [description] [created at] [foldername] [username]`
      Each field should be separated by whitespace or tab characters.
      The `[created at]` is a human-readable date/time format.
      The order of printed file information is determined by the `--sort-name` or `--sort-created` combined with asc or desc flags.
      The `--sort-name` means sorting by `[filename]` .
      If neither `--sort-name` nor `--sort-created` is provided, sort the list by `[filename]` in ascending order.
      - `Warning: The folder is empty.`
      - `Error: The [username] doesn't exist.`
      - `Error: The [foldername] doesn't exist.`
      - `Prompt the user the usage of the command if there is an invalid flag.(should output to STDERR)`
