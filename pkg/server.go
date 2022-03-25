package server

import (
	"github.com/alexbt/sample-golang/pkg/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Start() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	r.GET("/json", controller.GetJson)
	r.GET("/user", controller.GetUser)

	r.Run()
}
