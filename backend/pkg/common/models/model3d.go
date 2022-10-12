package models

type Model3D struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Author      string `json:"author" gorm:"foreignKey"`
	Description string `json:"description"`
	Price 		float64  `json:"price"`
	SerializedFile3D    []byte  `json:"serialized_file_3d"`
    FileNameAndExtension    string  `json:"file_name_and_extension"`
}