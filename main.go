package main

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", homePage)
	router.GET("/backup-db", backupDB)
	router.GET("/health", healthCheck)

	router.Run()
}

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "This is my home page")
}

func backupDB(c *gin.Context) {
	log.Println("DB Backup started ...")

	time.Sleep(1 * time.Second)
	err := sendPebbleNotification()
	if err != nil {
		log.Println("DB Backup failed!")
		c.String(http.StatusInternalServerError, "DB Backup failed")
		return
	}

	log.Println("DB Backup finished!")
	c.String(http.StatusOK, "DB Backup finished succesfully")
}

func sendPebbleNotification() error {
	cmd := exec.Command("/charm/bin/pebble", "notify", "guotiexin.com/db/backup", "path=/tmp/mydb.sql")
	if err := cmd.Run(); err != nil {
		return errors.Join(errors.New("couldn't execute a pebble notify: "), err)
	}
	return nil
}

func healthCheck(c *gin.Context) {
	if rand.Intn(10) == 0 {
		c.String(http.StatusInternalServerError, "Health check failed")
		return
	}
	c.String(http.StatusOK, "Health check passed")
}
