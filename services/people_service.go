package services

import (
	"github.com/goInter/goNetworkApi/domain/people"
	"github.com/goInter/goNetworkApi/utils/errors"
)

func CreatePerson(person people.Person) (*people.Person, *errors.RestErrors) {
	// validate
	err := person.Validate()
	if err != nil {
		return nil, err
	}

	// saving person in db
	if err := person.Save(); err != nil {
		return nil, err
	}

	return &person, nil
}

func GetPersonById(peopleId int64) (*people.Person, *errors.RestErrors) {
	result := &people.Person{Pid: peopleId}
	// getting person
	if err := result.Get(); err != nil {
		return nil, err
	}
	// return person
	return result, nil
}

func UpdatePerson(isPartial bool, person people.Person) (*people.Person, *errors.RestErrors) {
	// look i db to get user with matching id
	current, err := GetPersonById(person.Pid)
	if err != nil {
		return nil, err
	}

	if err := person.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if person.FirstName != "" {
			current.FirstName = person.FirstName
		}
		if person.LastName != "" {
			current.LastName = person.LastName
		}
		if person.Email != "" {
			current.Email = person.Email
		}
		if person.Gender != "" {
			current.Gender = person.Gender
		}

	} else {
		current.FirstName = person.FirstName
		current.LastName = person.LastName
		current.Email = person.Email
		current.Gender = person.Gender
	}

	// execute update
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil

}

func DeletePerson(personId int64) *errors.RestErrors {
	person := &people.Person{Pid: personId}
	return person.Delete()
}
