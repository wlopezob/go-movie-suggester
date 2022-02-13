package api

import "github.com/wlopezob/go-movie-suggester/internal/database"

type Services struct {
	search MovieSearch
	users  userGateway
}

func NewServices() Services {
	mysqlClient := database.NewMySQLClient()
	return Services{
		search: &MovieService{mysqlClient},
		users:  &UserService{mysqlClient},
	}
}

type WebServices struct {
	Services
	tokenKey string
}

func start(tokenKey string) *WebServices {
	return &WebServices{NewServices(), tokenKey}
}
