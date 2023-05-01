package repository

type RepositoryInterface interface {
	CreateFileInfo() error
	ReadFileInfo() (*FileInfo, error)
	UpdateFileInfo(file *FileInfo) error
	DeleteFileInfo(file *FileInfo) error
}
