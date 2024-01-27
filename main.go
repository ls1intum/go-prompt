package main

import (
	"context"
	"go-prompt/config"
	"go-prompt/db"
	"go-prompt/model/enum"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	if os.Getenv("DEBUG") == "true" {
		log.SetLevel(log.DebugLevel)
		log.Warn("DEBUG is true")
	}
	config.Load()

	log.Debugf("Config: %+v", config.GlobalConfig)

	pool, err := db.Connect(context.Background())
	if err != nil {
		log.WithError(err).Fatal("Error connecting to database")
	}
	defer pool.Close()

	var courses []enum.Course
	err = pool.QueryRow(context.Background(), "SELECT courses_taken FROM developer_application WHERE $1 = ANY(devices) LIMIT 1", enum.MACBOOK).Scan(&courses)
	if err != nil {
		log.WithError(err).Fatal("Error querying database")
	}
	log.Info(courses[0])

	// gin.SetMode(gin.ReleaseMode)
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// log.Panic(r.Run(fmt.Sprintf(":%d", 8080)))
}
