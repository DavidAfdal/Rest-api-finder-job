package main

import (
	dbconfig "finder-job/config/db"
	authcontroller "finder-job/controller/auth"
	userrepo "finder-job/repository/user"
	authservices "finder-job/services/auth"

	"github.com/gin-gonic/gin"
)

func main() {
    server := gin.Default()
	db := dbconfig.ConnectDB()
	userRepo := userrepo.NewUserRepo(db)
	authServices := authservices.NewAuthService(userRepo)
	authController := authcontroller.NewAuthControlle(authServices)

	server.GET("/", authController.Login)

	server.Run(":5000")
}