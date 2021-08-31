package helper

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

func IsMySQLErrorDuplicate(err error) bool {
	var mysqlErr *mysql.MySQLError

	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return true
	}

	return false
}
