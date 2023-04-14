package types

type DatabaseFile struct {
	Name, Path string
}

func (df *DatabaseFile) Get() (name, path string) {
	return df.Name, df.Path
}

