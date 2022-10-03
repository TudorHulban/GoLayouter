# Go Layout Creator 
Application creates go app layout folders and files as per the specified file.   
The operations in the input file are as below.
## Change Directory 
Prefix the new folder with an exclamation mark.
```sh
! path
```
## File creation
### Create files for specific package:
```sh
# package <package name>
```
Example :
```sh
subfolder1
 # package main
 file1.go file2.go
    file3.go
 file4.go
subfolder2
```
The app will create all for specified files as part of the secified Go package.
### Create test files
Example :
```sh
subfolder1
 # package main
 file1.go file2.go
    file3.go
 # t
 file4.go
subfolder2
```
would create for file4.go also the test file for it as file4_test.go.

## Line indendation
Line indentation creates subfolder as per indentation and as per below example:
```
folder1
 subfolder1
 subfolder2
  subsubfolder2
 subfolder3
folder2
```
`folder1` will be created with  `subfolder1` and `subfolder2`.  
`subfolder2` will be parent of `subsubfolder1`.  
`folder2` will be created in same folder as `folder1`.
