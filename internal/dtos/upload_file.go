package dtos

import "mime/multipart"

type Form struct {
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}
