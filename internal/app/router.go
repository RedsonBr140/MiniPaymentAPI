package app

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(gin.Logger())

	// v1 := r.Group("/api/v1")

	return r
}
