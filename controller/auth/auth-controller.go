package authcontroller

import (
	authservices "finder-job/services/auth"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *authservices.AuthService
}


func NewAuthControlle(authService *authservices.AuthService) *AuthController {
	return &AuthController{AuthService: authService}
}

func (c *AuthController) Register(ctx *gin.Context){}
func (c *AuthController) Login(ctx *gin.Context){}


func (c *AuthController) Logout(ctx *gin.Context){}