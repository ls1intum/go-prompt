package controller

import (
	"go-prompt/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAPI(s *service.Service) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/developer-applications", func(c *gin.Context) {
		developerApplications, err := s.GetAllDeveloperApplications(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, developerApplications)
	})

	return r
}
