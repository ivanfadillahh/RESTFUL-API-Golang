package user

import (
	"log"
	// "net/http"

	"github.com/gin-gonic/gin"

	Config "goapi/config"
	ModelUser "goapi/struct"
)

func AllUsers(c *gin.Context){
	var db, errdb = Config.Connectdb()
	if errdb != nil{
		c.JSON(500, gin.H{"result": "Error Connection"})
		log.Println("Error Connection")
		return
	}
	defer db.Close()
	var ResultUser ModelUser.ShowUser
	// query show all user on signle row
	// var stmtuser = db.QueryRow("select * from tb_user").Scan(&ResultUser.UserID, &ResultUser.UserName, &ResultUser.UserFullname, &ResultUser.UserEmail, &ResultUser.UserPassword, &ResultUser.DateCreated)
	// if stmtuser != nil {
	// 	c.JSON(400, gin.H{"result": &ResultUser, "message": stmtuser})
	// } else {
	// 	c.JSON(200, gin.H{"result": &ResultUser, "message": "Success get all users"})
	// }
	// end query

	res, err := db.Query("select * from tb_user")
	defer res.Close()

	if err != nil{
		c.JSON(500, gin.H{"result": "Error Connection"})
		log.Println("Error Connection")
		return
	}
	
	for res.Next(){
		err := res.Scan(&ResultUser.UserID, &ResultUser.UserName, &ResultUser.UserFullname, &ResultUser.UserEmail, &ResultUser.UserPassword, &ResultUser.DateCreated)
		
		if err != nil {
			c.JSON(400, gin.H{"result": &ResultUser, "message": "Fail to get users"})
		} else {
			c.JSON(200, gin.H{"result": &ResultUser})
		}
	}
}

func CreateUser(c *gin.Context){
	var db, errdb = Config.Connectdb()
	if errdb != nil{
		c.JSON(500, gin.H{"result": "Error Connection"})
		log.Println("Error Connection")
		return
	}

	defer db.Close()

	// param Input
	var txtuser ModelUser.ReqUser
	c.BindJSON(&txtuser)
	var UserName = txtuser.UserName
	var UserFullname = txtuser.UserFullname
	var UserEmail = txtuser.UserEmail
	var UserPassword = txtuser.UserPassword

	// query insert user
	if UserName != "" {
		_, err := db.Query("insert into tb_user (user_name,user_fullname,user_email,user_password) values (?,?,?,?)", UserName,UserFullname,UserEmail,UserPassword)
		if err != nil {
			c.JSON(400, gin.H{"result": err, "message": "Failed to insert user"})
		} else {
			c.JSON(201, gin.H{"result": nil, "message": "Success insert new user"})
		}	
	} else {
		c.JSON(400, gin.H{"result":"Bad Request Body"})
	}
	// end query
}

func UpdateUser(c *gin.Context){
	var db, errdb = Config.Connectdb()
	if errdb != nil {
		c.JSON(500, gin.H{"result": "Error Connection"})
		log.Println("Error Connection")
		return
	}

	defer db.Close()
	// param input
	var txtuser ModelUser.ReqUser
	c.BindJSON(&txtuser)
	var UserID = c.Param("id")
	// var UserID = txtuser.UserID
	var UserName = txtuser.UserName
	var UserFullname = txtuser.UserFullname
	var UserEmail = txtuser.UserEmail
	var UserPassword = txtuser.UserPassword

	// query update user
	_, err := db.Query("UPDATE tb_user SET user_name = ?,user_fullname = ?,user_email = ?,user_password = ? WHERE user_id = ?", UserName, UserFullname,UserEmail,UserPassword,UserID)
	if err != nil{
		c.JSON(400, gin.H{"result": err, "message": "Failed to update user"})
	} else{
		c.JSON(200, gin.H{"result": err, "message": "Success update user"})
	}
	// end query
}

func DeleteUser(c *gin.Context){
	var db, errdb = Config.Connectdb()
	if errdb != nil {
		c.JSON(500, gin.H{"result": "Error Connection"})
		log.Println("Error Connection")
		return
	}
	defer db.Close()

	// param input
	var UserID = c.Param("id")
	// query delete user
	_, err := db.Query("delete from tb_user where user_id = ?", UserID)
	if err != nil {
		c.JSON(400, gin.H{"result": err, "message": "Failed to delete user"})
	} else{
		c.JSON(200, gin.H{"result": nil, "message": "Success delete user"})
	}
	// end query
}

func DetailUser(c *gin.Context){
	var db, errdb = Config.Connectdb()
	if errdb != nil {
		c.JSON(500, gin.H{"result": "Error Connection"})
		log.Println("Error Connection")
		return
	}
	defer db.Close()

	var ResultUser ModelUser.ShowUser
	var UserID = c.Param("id")
	// query show user by id
	var stmtuser = db.QueryRow("select * from tb_user where user_id = ?", UserID).Scan(&ResultUser.UserID, &ResultUser.UserName, &ResultUser.UserFullname, &ResultUser.UserEmail, &ResultUser.UserPassword, &ResultUser.DateCreated)
	if stmtuser != nil {
		c.JSON(400, gin.H{"result": &ResultUser, "message": stmtuser})
	} else{
		c.JSON(200, gin.H{"result": &ResultUser, "message": "Success get user"})
	}
	// end query
}

// function test request headers
func TestHeaders(c *gin.Context){
	c.JSON(200, gin.H{"token": c.Request.Header["Token"]})
}