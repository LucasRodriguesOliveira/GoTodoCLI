package filemanager

func (fm *FileManager) Write(data []byte) error {
  fm.OpenOrCreate()
	defer fm.File.Data.Close()

	_, err := fm.File.Data.Write(data)

	if err != nil {
		return err
	}

	return nil
}
