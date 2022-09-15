package models

type Books struct {
	ID        int    `json:"id"`
	Title     string `json:"title" validate:"required"`
	Isbn      string `json:"isbn"`
	Writer    string `json:"writer" validate:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type BooksResponse struct {
	Message string `json:"response"`
	Data    Books  `json:"data"`
}
