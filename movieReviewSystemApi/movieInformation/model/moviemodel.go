package model

import "github.com/zeromicro/go-zero/core/stores/mon"

var _ MovieModel = (*customMovieModel)(nil)

type (
	// MovieModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMovieModel.
	MovieModel interface {
		movieModel
	}

	customMovieModel struct {
		*defaultMovieModel
	}
)

// NewMovieModel returns a model for the mongo.
func NewMovieModel(url, db, collection string) MovieModel {
	conn := mon.MustNewModel(url, db, collection)
	return &customMovieModel{
		defaultMovieModel: newDefaultMovieModel(conn),
	}
}
