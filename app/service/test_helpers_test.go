package service

import "fmt"

func (serv *Service) CheckPathsExists() error {
	for _, path := range serv.paths {
		errCheck := path.ObjectPath.CheckIfPathExists()
		if errCheck != nil {
			return errCheck
		}
	}

	return nil
}

func (serv *Service) DeletePaths() error {
	for index := len(serv.paths) - 1; index >= 0; index-- {
		err := serv.paths[index].ObjectPath.DeletePath()
		if err != nil {
			return err
		}
	}

	return nil
}

func (serv *Service) ChangeDirectory(newPath string) error {
	for _, path := range serv.paths {
		if err := path.ObjectPath.ChangeDirectory(newPath); err != nil {
			return fmt.Errorf("error : %w", err)
		}
	}

	return nil
}
