package files

type FileServiceImpl interface {
	UploadFile() error
}

type fileService struct{}

func NewFileService() FileServiceImpl {
	return &fileService{}
}

// UploadFile implements FileServiceImpl.
func (f *fileService) UploadFile() error {
	panic("unimplemented")
}
