package dto

type CreateCommentDTO struct {
	Content string `json:"content" validate:"required,min=1,max=500"`
}
