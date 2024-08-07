package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
* User Controller
* Note: Controller, Service, Repository need to follow dependency inversion principle
**/

type UserController interface {
	SignUp(c *gin.Context)
}

type userController struct {
	userService UserService
}

func NewUserController(userService UserService) UserController {
	return &userController{
		userService: userService,
	}

}

/**
* Sign Up User
**/
func (ctrl *userController) SignUp(c *gin.Context) {
	var req CreateUserReq

	// bind incoming JSON for validation
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation error:": err.Error()})
		return
	}

	// create user
	createdUser, err := ctrl.userService.UserSignup(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Could not create user. Error:": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created test.", "data": createdUser})
}

/**
* Authenticates User and Signs them in by Providing Access Token / Refresh Token
**/
func (ctrl *userController) SignIn(c *gin.Context) {
	var req UserSignInReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Validation Errorr:": err.Error()})
		return
	}

	user, err := ctrl.userService.AuthenticateUser(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error while authenticating:": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message:": "Success. User Authenticated.", "data": user})
}
