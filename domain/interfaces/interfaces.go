package interfaces

type IWritter interface {
	WriteToDisk() error
}

type IFileOperations interface {
	CheckIfPathExists() error
	DeletePath() error
}
