package domain

type UserPayload struct {
	Id       int    `json:"id"`
	Fname    string `json:"fname" binding:"required,min=3"`
	Lname    string `json:"lname" binding:"required,min=3"`
	Username string `json:"username" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,min=10,email"`
	Password string `json:"password" binding:"required,min=12"`
}
