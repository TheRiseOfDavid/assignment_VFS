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
      - Description: To register the user.
      - Response
      ```bash
      # Add [username] successfully.
      # Error: The [username] has already existed.
      # Error: The [username] contain invalid chars.
      ```
2.  Folder Management
    - `create-folder [username] [foldername] [description]?`

## 要寫的東西

- Include unit tests for the implemented features
- Options 的資料大小寫不分
- UnitTest 全部都要補
- 看要不要把重複地寫成 interface(C++ 的 template)

## issue

我看這篇文章說
https://stackoverflow.com/questions/54377597/how-to-make-a-function-that-receives-an-array-of-custom-interface

go 沒辦法在 interface method 中的 parma 使用 slice.
因為她只能轉型 object 不能轉型 object's slice.

所以我沒辦法寫一個指標 Interface 來多型? 對嘛

我的 package 全部都要是 main 嗎
