package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	Config "goapi/config"
	Connect "goapi/connect"
	User "goapi/user"

	"github.com/rs/cors"
)

func main(){
	addr, err := Config.MyPort()
	if err != nil{
		log.Fatal(err)
	}

	router := gin.Default()

	v1 := router.Group("/api/v1/goapi")
	{
		v1.GET("/testconnection", Connect.Tesconnect)
		// crud API
		v1.GET("/getallusers", User.AllUsers)
		v1.POST("/createuser", User.CreateUser)
		v1.PUT("/updateuser/:id", User.UpdateUser)
		v1.DELETE("/deleteuser/:id", User.DeleteUser)
		v1.GET("/getuser/:id", User.DetailUser)

		// test request headers
		v1.GET("/test", User.TestHeaders)
	}
	c := cors.AllowAll()

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(addr, handler))
}