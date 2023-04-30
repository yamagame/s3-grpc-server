package repository

type ClientInterface interface {
	CreateFileInfo() error
	ReadFileInfo() (*FileInfo, error)
	UpdateFileInfo(file *FileInfo) error
	DeleteFileInfo(file *FileInfo) error
}
