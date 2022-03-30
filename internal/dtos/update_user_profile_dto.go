package dtos

type UpdateUserProfileDTO struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	OldPassword string `json:"old_password" binding:"min=6,max=200"`
	Password    string `json:"password" binding:"min=6,max=200"`
}
