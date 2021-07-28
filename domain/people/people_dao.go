package people

import (
	"database/sql"
	"fmt"
	"github.com/GoMySelf/udemyMicroServices/utils/mysql_utils"
	"github.com/goInter/goNetworkApi/datasources/mysql"
	"github.com/goInter/goNetworkApi/utils/errors"
	"strings"
)

const (
	errorNoRows       = "no rows in result set"
	queryInsertPerson = "INSERT INTO Person(first_name, last_name, email, gender) VALUES(?,?,?,?);"
	queryGetPerson    = "SELECT pid, first_name, last_name, email, gender FROM Person WHERE pid=?;"
	queryUpdatePerson = "UPDATE Person SET first_name=?, last_name=?, email=?, gender=? WHERE pid=?;"
	queryDeletePerson = "DELETE FROM Person WHERE pid=?;"
)

// Get -> get single user
func (person *Person) Get() *errors.RestErrors {
	// preparing statement
	stmt, err := people_db.Client.Prepare(queryGetPerson)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(stmt)

	// get singe person where id == query id
	result := stmt.QueryRow(person.Pid)

	// take from db
	if getErr := result.Scan(&person.Pid, &person.FirstName, &person.LastName, &person.Email, &person.Gender); getErr != nil {
		if strings.Contains(getErr.Error(), errorNoRows) {
			return errors.NewNotFoundError("Person not found")
		}
		return errors.NewInternalServerError("error when trying to get user")
	}

	// check connection
	if err := people_db.Client.Ping(); err != nil {
		panic(err.Error())
	}

	return nil
}

// Save -> save person in DB
func (person *Person) Save() *errors.RestErrors {

	stmt, err := people_db.Client.Prepare(queryInsertPerson)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(stmt)

	// execute statement to the database
	insertResult, saveErr := stmt.Exec(person.FirstName, person.LastName, person.Email, person.Gender, person.Pid)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	// personID gonna be last id insert the db
	personId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("Error when trying to save person")
	}
	person.Pid = personId

	return nil

}

func (person *Person) Update() *errors.RestErrors {
	// open statement
	stmt, err := people_db.Client.Prepare(queryUpdatePerson)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(stmt)
	// first_name=?, last_name=?, email=?, gender=?
	_, err = stmt.Exec(person.FirstName, person.LastName, person.Email, person.Gender, person.Pid)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (person *Person) Delete() *errors.RestErrors {
	stmt, err := people_db.Client.Prepare(queryDeletePerson)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(stmt)

	if _, err := stmt.Exec(person.Pid); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}
