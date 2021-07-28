package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {

	mapUrl()

	err := router.Run(":3000")
	if err != nil {
		return
	}

}
