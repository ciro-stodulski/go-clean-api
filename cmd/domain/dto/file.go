package dto

// Example of form file
import "mime/multipart"

type File struct {
	FileName string                `form:"file_name" binding:"required"`
	File     *multipart.FileHeader `form:"file" binding:"required"`
}
