package people

import (
	"github.com/goInter/goNetworkApi/utils/errors"
	"strings"
)

type Person struct {
	Pid       int64  `json:"pid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

// Validate -> validate email on a fly
func (person *Person) Validate() *errors.RestErrors {
	person.LastName = strings.TrimSpace(person.LastName)
	person.FirstName = strings.TrimSpace(person.FirstName)

	person.Email = strings.TrimSpace(strings.ToLower(person.Email))
	if person.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}
