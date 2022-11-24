package service

func (serv *Service) CheckPathsExists() error {
	for _, path := range serv.paths {
		errCheck := path.CheckIfPathExists()
		if errCheck != nil {
			return errCheck
		}
	}

	return nil
}

func (serv *Service) DeletePaths() error {
	for index := len(serv.paths) - 1; index >= 0; index-- {
		err := serv.paths[index].DeletePath()
		if err != nil {
			return err
		}
	}

	return nil
}
