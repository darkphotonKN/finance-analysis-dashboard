package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/** User Controller
*
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
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created test.", "data": createdUser})
}
