package main

import (
	"errors"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", homePage)
	router.GET("/backup-db", backupDB)

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
