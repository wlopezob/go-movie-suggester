package models

type InternalError struct {
	Message string `json:"message"`
}

type LoginCMD struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type WishMovieCMD struct {
	MovieId string `json:"movieid"`
	Comment string `json:"comment"`
}
