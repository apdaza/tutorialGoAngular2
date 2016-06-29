package main

import (
	"github.com/gin-gonic/gin"

	. "github.com/apdaza/tutorial/utils"
	. "github.com/apdaza/tutorial/controllers"
)

func main() {
	r := gin.Default()

	r.Use(Cors())

	web := r.Group("api/contactos")
	{
		web.GET("/contactos", GetContactos)
		web.POST("/contactos", PostContactos)

		web.OPTIONS("/contactos", Options)        // POST
	}
	r.Run(":8080")
}
