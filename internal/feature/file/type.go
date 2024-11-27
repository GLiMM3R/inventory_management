package files

type PutObjectRequest struct {
	FileName string `json:"file_name"`
}

type PutObjectResponse struct {
	URL      string `json:"url"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
}

type GetObjectRequest struct {
	Directory string `json:"directory"`
	FileName  string `json:"file_name"`
}

type GetObjectResponse struct {
	Directory string `json:"directory"`
	FileName  string `json:"file_name"`
}
