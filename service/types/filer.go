package types

type FileData interface {
	Get()(name, path string)
}
