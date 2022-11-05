package interfaces

type IWritter interface {
	WriteToDisk() error
	CheckIfExists() error
	DeletePath() error
	GetPath() string
}
