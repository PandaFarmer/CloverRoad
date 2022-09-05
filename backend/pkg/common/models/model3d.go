package models

type Model3D struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Author      string `json:"author" gorm:"foreignKey"`
	Description string `json:"description"`
	Price 		float64  `json:"price"`
	BlobData    []byte  `json:"blob_data"`
    FileName    string  `json:"file_name"`
}