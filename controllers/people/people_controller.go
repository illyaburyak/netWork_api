package people

import (
	"github.com/gin-gonic/gin"
	"github.com/goInter/goNetworkApi/domain/people"
	"github.com/goInter/goNetworkApi/services"
	"github.com/goInter/goNetworkApi/utils/errors"
	"net/http"
	"strconv"
)

func getParams(userIdParams string) (int64, *errors.RestErrors) {
	personId, personErr := strconv.ParseInt(userIdParams, 10, 64)
	if personErr != nil {
		return 0, errors.NewBadRequestError("Invalid user id")
	}
	return personId, nil
}

func CreatePerson(c *gin.Context) {
	var person people.Person

	// parse request body
	err := c.ShouldBindJSON(&person)
	if err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	// sending person to the services
	result, saveErr := services.CreatePerson(person)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	// sending back to the client person as json
	c.JSON(http.StatusCreated, result)
}

func GetPersonById(c *gin.Context) () {
	// url params
	personId, personErr := getParams(c.Param("person_id"))
	if personErr != nil {
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}
	// sending peopleId to the service
	person, getErr := services.GetPersonById(personId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	// respond to the client
	c.JSON(http.StatusOK, person)
}
func UpdatePerson(c *gin.Context) {
	// url params
	personId, idErr := getParams(c.Param("person_id"))
	if idErr != nil {
		err := errors.NewBadRequestError("Invalid user id")
		c.JSON(err.Status, err)
		return
	}

	// populate person
	var person people.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	person.Pid = personId

	// check the method
	isPartial := c.Request.Method == http.MethodPatch

	// update partial or not, depends on a request
	result, err := services.UpdatePerson(isPartial, person)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeletePerson(c *gin.Context) {
	personId, idErr := getParams(c.Param("person_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeletePerson(personId); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
