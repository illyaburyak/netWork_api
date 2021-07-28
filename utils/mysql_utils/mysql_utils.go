package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/goInter/goNetworkApi/utils/errors"
	"strings"
)

const (
	errorNoRows = "no rows in the result set"
)

func parseError(err error) *errors.RestErrors {
	// try to convert our error to an sql error
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		// if data already exists
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("No record matching given id")
		}
		return errors.NewInternalServerError("Error parsing database response")
	}
	switch sqlErr.Number {
	// error saying duplicated key
	case 1062:
		return errors.NewBadRequestError("Email already exists")
	}
	return errors.NewInternalServerError("Error processing request")
}
