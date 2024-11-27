package media

type CreateMediaDto struct {
	FileName    string `json:"file_name"`
	FilePath    string `json:"file_path"`
	FileType    string `json:"file_type"`
	FileSize    uint   `json:"file_size"`
	MediaType   string `json:"media_type"`
	Description string `json:"description"`
}

type CreateMediaRequest struct {
	FileName    string `json:"file_name"`
	FileType    string `json:"file_type"`
	FileSize    uint   `json:"file_size"`
	MediaType   string `json:"media_type"`
	Description string `json:"description"`
}

type MediaResponse struct {
	MediaID     string `json:"media_id"`
	FileURL     string `json:"file_url"`
	FileName    string `json:"file_name"`
	FilePath    string `json:"file_path"`
	FileType    string `json:"file_type"`
	FileSize    uint   `json:"file_size"`
	MediaType   string `json:"media_type"`
	Description string `json:"description"`
	CreatedAt   int64  `json:"created_at"`
	UpdatedAt   int64  `json:"updated_at"`
}
