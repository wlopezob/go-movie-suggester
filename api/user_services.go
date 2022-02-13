package api

import (
	"github.com/gofiber/fiber/v2/utils"
	"github.com/wlopezob/go-movie-suggester/internal/database"
	"github.com/wlopezob/go-movie-suggester/internal/logs"
	"github.com/wlopezob/go-movie-suggester/pkg/models"
)

type CreateUserCMD struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
}

type UserSummary struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	JWT      string `json:"token"`
}
type userGateway interface {
	SaveUser(cmd CreateUserCMD) (*UserSummary, error)
	Login(cmd models.LoginCMD) string
	AddWishMovie(userId, movieId string, comment string) error
}

type UserService struct {
	*database.MysqlClient
}

func (us *UserService) Login(cmd models.LoginCMD) string {
	id := ""
	smt, _ := us.Prepare(`select id from "user" where username=$1 and password=$2`)
	err := smt.QueryRow(cmd.Username, cmd.Password).Scan(&id)
	if err != nil {
		logs.Error("Cannot query login")
		panic(err.Error())
	}
	return id
}
func (us *UserService) SaveUser(cmd CreateUserCMD) (*UserSummary, error) {
	id := utils.UUID()
	smt, _ := us.Prepare(`insert into "user"(id,username,"password") values($1,$2,$3)`)
	_, err := smt.Exec(id, cmd.Username, cmd.Password)
	if err != nil {
		logs.Error("cannot insert user: " + err.Error())
		return nil, err
	}
	return &UserSummary{
		ID:       id,
		Username: cmd.Username,
		JWT:      "",
	}, nil
}

func (us *UserService) AddWishMovie(userId, movieId string, comment string) error {
	smt, _ := us.Prepare(`insert into wish_list(user_id, movie_id,"comment") values($1,$2,$3)`)
	_, err := smt.Exec(userId, movieId, comment)
	if err != nil {
		logs.Error("Cannot save movie")
		return err
	}
	return nil
}
