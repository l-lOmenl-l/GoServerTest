package domain

type TypeProduct struct {
	Id   int    `json:"-" db:"id"`
	Name string `json:"type_product" binding:"required" db:"name"`
}
