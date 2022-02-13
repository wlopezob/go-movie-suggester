package api

import (
	"github.com/wlopezob/go-movie-suggester/internal/database"
	"github.com/wlopezob/go-movie-suggester/internal/logs"
)

type MovieFilter struct {
	Title    string `json:"title,omitempty"`
	Genre    string `json:"genre"`
	Director string `json:"drector"`
}

type Movie struct {
	Id          string `json:"id"`
	Title       string `json:"tile"`
	Cast        string `json:"cast"`
	ReleaseDate string `json:"release_date"`
	Genre       string `json:"genre"`
	Director    string `json:"director"`
}
type MovieSearch interface {
	Search(filter MovieFilter) ([]Movie, error)
}

type MovieService struct {
	*database.MysqlClient
}

func (s *MovieService) Search(filter MovieFilter) ([]Movie, error) {
	tx, err := s.Begin()
	if err != nil {
		logs.Error("cannot create transaction: " + err.Error())
		return nil, err
	}
	rows, err := tx.Query("select id, title, 'cast', release_date, genre, director from movie")
	if err != nil {
		logs.Error("cannot read movies:" + err.Error())
		_ = tx.Rollback()
		return nil, err
	}

	var movies []Movie
	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.Id, &movie.Title, &movie.Cast, &movie.ReleaseDate, &movie.Genre,
			&movie.Director)
		if err != nil {
			logs.Error("cannot read moview " + err.Error())
		}
		movies = append(movies, movie)

	}
	tx.Commit()
	return movies, nil
}
