package tran

import "dragonsss.cn/lbk_user/internal/database"

// Transaction 事务的操作
type Transaction interface {
	Action(func(conn database.DbConn) error) error
}
