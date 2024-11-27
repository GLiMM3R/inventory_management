package schema

type Media struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement:true;column:id"`
	MediaID     string `json:"media_id" gorm:"primaryKey;unique;column:media_id"`
	FileName    string `json:"file_name" gorm:"column:file_name"`
	FilePath    string `json:"file_path" gorm:"column:file_path"`
	FileType    string `json:"file_type" gorm:"column:file_type"`
	FileSize    uint   `json:"file_size" gorm:"column:file_size"`
	MediaType   string `json:"media_type" gorm:"column:media_type"`
	Description string `json:"description" gorm:"column:description"`
	CreatedAt   int64  `json:"created_at" gorm:"autoCreateTime;column:created_at"`
	UpdatedAt   int64  `json:"updated_at" gorm:"autoUpdateTime;column:updated_at"`
}
