package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/jwt-project/helper"
)

func main() {
	router := gin.Default()
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome Home")
	})

	server := &http.Server{
		Addr:    ":8090",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.Errorpanic(err)
}
