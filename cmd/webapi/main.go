package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/project/webapi/internal/handlers"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/getArticles", func(ctx *gin.Context) {
		body := ctx.Request.Body

		var bodyValue string
		if body != nil {
			value := make([]byte, 1000)
			body.Read(value)
			bodyValue = string(value)
		}

		handler, err := handlers.New(strings.Trim(bodyValue, "\x00"))

		if err != nil {
			ctx.Error(err)
			return
		}

		article, err := handler.GetArticles()

		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": article,
		})
	})
	r.Run()
}
