package connect

import (
	"log"
	// "net/http"

	"github.com/gin-gonic/gin"

	Config "goapi/config"
)

func Tesconnect(c *gin.Context){
	var db, errdb = Config.Connectdb()
	if errdb != nil {
		c.JSON(500, gin.H{"result":"Error Connection"})
		log.Println("Error Connection")
		return
	}
	defer db.Close()
	c.JSON(200, gin.H{"result":"Connected"})
}