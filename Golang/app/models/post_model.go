package models

type Post struct {
	ID     uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Title  string `json:"title"`
	Body   string `json:"body" gorm:"type:text"`
	UserID uint   `json:"user_id"`
	User   *User  `json:"user,omitempty"`
}

type CreatePostRequest struct {
	Title string `json:"title" validate:"required,min=8,max=30"`
	Body  string `json:"body" validate:"required,min=8"`
}

type UpdatePostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
