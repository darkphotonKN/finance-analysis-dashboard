package user

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/darkphotonKN/finance-analysis-dashboard/internal/shared/constants"
	"github.com/gin-gonic/gin"
)

/**
* User Controller
* Note: Controller, Service, Repository need to follow dependency inversion principle
**/

type UserController interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
	FindAllUsers(c *gin.Context)
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
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message:": fmt.Sprintf("validation error: ", err.Error())})
		return
	}

	// create user
	createdUser, err := ctrl.userService.UserSignup(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message:": fmt.Sprintf("Could not create user, error: %v", err.Error())})
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

	authenticatedInfo, err := ctrl.userService.AuthenticateUser(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message:": fmt.Sprintf("Error while authenticating: %v", err.Error())})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message:": "Success. User Authenticated.", "data": authenticatedInfo})
}

/**
* Gets all users (Admin)
**/
func (ctrl *userController) FindAllUsers(c *gin.Context) {
	// Get and validate query parameters
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	sort := c.DefaultQuery("sort", "first_name")
	order := constants.SortOrder(c.DefaultQuery("order", "asc"))

	fmt.Printf("page: %d, pageSize: %d, sort: %s order: %s\n", page, pageSize, sort, order)

	keyword := c.Query("keyword")

	users, err := ctrl.userService.FindAllUsers(page, pageSize, keyword, sort, order)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message:": "Could not retreive users."})
		return
	}

	// serialize users response
	var usersResponseSlice []UserRes

	for _, user := range *users {
		usersResponseSlice = append(usersResponseSlice, UserRes{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": usersResponseSlice})
}
