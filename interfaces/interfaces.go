package interfaces

type IWritter interface {
	WriteToDisk() error
	CheckIfExists() error
	DeletePath() error
	GetPath() string
}

// TODO:
// add objects that satisfy the interface

func Write(paths []IWritter) error {
	for _, path := range paths {
		err := path.WriteToDisk()
		if err != nil {
			return err
		}
	}

	return nil
}

func CheckInterface(paths []IWritter) error {
	for _, path := range paths {
		errCheck := path.CheckIfExists()
		if errCheck != nil {
			return errCheck
		}
	}

	return nil
}

func DeleteInterface(paths []IWritter) error {
	for index := len(paths) - 1; index >= 0; index-- {
		err := paths[index].DeletePath()
		if err != nil {
			return err
		}
	}

	return nil
}
