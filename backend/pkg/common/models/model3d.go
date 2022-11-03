package models

import "gorm.io/gorm"

type Model3D struct {
	gorm.Model         // adds ID, created_at etc.
	Title       			string `json:"title"`
	Author      			string `json:"author" gorm:"foreignKey"`
	Description 			string `json:"description"`
	Price 					float64  `json:"price"`
	SerializedPreviewFile   []byte  `json:"serialized_preview_file"`
	PreviewFileNameAndExtension string `json:"preview_file_name_and_extension"`
	SerializedFile3D    	[]byte  `json:"serialized_file_3d"`
    FileNameAndExtension    string  `json:"file_name_and_extension"`
}