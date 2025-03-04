package mysql

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ThumbUpModel = (*customThumbUpModel)(nil)

type (
	// ThumbUpModel is an interface to be customized, add more methods here,
	// and implement the added methods in customThumbUpModel.
	ThumbUpModel interface {
		thumbUpModel
		withSession(session sqlx.Session) ThumbUpModel
	}

	customThumbUpModel struct {
		*defaultThumbUpModel
	}
)

// NewThumbUpModel returns a model for the database table.
func NewThumbUpModel(conn sqlx.SqlConn) ThumbUpModel {
	return &customThumbUpModel{
		defaultThumbUpModel: newThumbUpModel(conn),
	}
}

func (m *customThumbUpModel) withSession(session sqlx.Session) ThumbUpModel {
	return NewThumbUpModel(sqlx.NewSqlConnFromSession(session))
}
