package book

type InputBook struct {
	Title string `json:"title" binding:"required"`
	Price int `json:"price" binding:"required,number"`
	Author string `json:"author" binding:"required"`
}