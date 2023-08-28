package sequence

import (
	"database/sql"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type MySQL struct {
	conn sqlx.SqlConn
}

func NewMySQL(dsn string) Sequence {
	conn := sqlx.NewMysql(dsn)
	return &MySQL{
		conn: conn,
	}
}

func (m *MySQL) Next() (seq uint64, err error) {
	var stmt sqlx.StmtSession
	stmt, err = m.conn.Prepare(sqlReplaceStub)
	if err != nil {
		logx.Errorw("mysql prepare failed",
			logx.Field("err", err),
		)
		return 0, err
	}
	defer stmt.Close()

	var rest sql.Result
	rest, err = stmt.Exec()
	if err != nil {
		logx.Errorw("mysql stmt exec failed",
			logx.Field("err", err),
		)
		return 0, err
	}

	var lid int64
	lid, err = rest.LastInsertId()
	if err != nil {
		logx.Errorw("mysql get last insert id failed",
			logx.Field("err", err),
		)
		return 0, err
	}
	return uint64(lid), nil
}
