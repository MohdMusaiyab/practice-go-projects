package dto

type CreatePostDTO struct {
	Title   string `json:"title" validate:"required,min=5,max=100"`
	Content string `json:"content" validate:"required"`
}

type UpdatePostDTO struct {
	Title   string `json:"title" validate:"omitempty,min=5,max=100"`
	Content string `json:"content" validate:"omitempty"`
}
