package app

import "github.com/goInter/goNetworkApi/controllers/people"

func mapUrl() {
	router.POST("/person", people.CreatePerson)
	router.GET("/person/:person_id", people.GetPersonById)
	router.PUT("/person/:person_id", people.UpdatePerson)
	router.PATCH("/person/:person_id", people.UpdatePerson)
	router.DELETE("/person/:person_id", people.DeletePerson)
}
