package usermodel

// penamaan varibel harus huruf besar agar variabel atau function itu bisa exported

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Image string `json:"profile_img"`
}