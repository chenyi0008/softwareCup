package models

type FileModel struct {
	Id   int
	Path string
}

func (FileModel) TableName() string {
	return "file"
}
