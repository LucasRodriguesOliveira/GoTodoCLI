package filemanager

func (fm *FileManager) OpenOrCreate() error {
  ok, err := fm.Check()

	if err != nil {
		return err
	}

	if ok {
		err = fm.open()
	} else {
		err = fm.create()
	}

	if err != nil {
		return err
	}

	return nil
}
