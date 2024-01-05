package model

type FileResponse struct {
	Uid    string `json:"uid"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Url    string `json:"url"`
	Size   int64  `json:"size"`
}
type File struct {
	ID        uint   `gorm:";primary_key;column:id" json:"id"`
	Uid       string `gorm:"column:uid" json:"uid"`
	Name      string `gorm:"column:name" json:"name"`
	Url       string `gorm:"column:url" json:"url"`
	Size      int64  `gorm:"column:size" json:"size"`
	FileType  string `gorm:"column:type" json:"file_type"`
	CreatedAt int64  `gorm:"column:created_at" json:"created_at"`
}
type Files []File

func (f File) TableName() string {
	return "files"
}

type AddFile struct {
	Uid       string `json:"uid"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Size      int64  `json:"size"`
	FileType  string `json:"type"`
	CreatedAt int64  `json:"created_at"`
}

func (a AddFile) GetFile() *File {
	return &File{
		Uid:       a.Uid,
		Name:      a.Name,
		Url:       a.Url,
		Size:      a.Size,
		FileType:  a.FileType,
		CreatedAt: a.CreatedAt,
	}
}
