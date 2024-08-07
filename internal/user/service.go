package user

import (
	"fmt"

	"github.com/darkphotonKN/finance-analysis-dashboard/internal/shared/models"
	"github.com/darkphotonKN/finance-analysis-dashboard/internal/util"
)

type UserService interface {
	UserSignup(CreateUserReq) (*models.User, error)
}

type userService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}

}

func (s *userService) UserSignup(createUserReq CreateUserReq) (*models.User, error) {

	newUser := models.User{
		FirstName: createUserReq.FirstName,
		LastName:  createUserReq.LastName,
		Email:     createUserReq.Email,
		Password:  util.HashPassword(createUserReq.Password),
		Role:      "user",
	}

	fmt.Printf("Creating user in service: %+v\n", newUser)

	createdUser, err := s.userRepository.CreateUser(&newUser)

	if err != nil {
		return createdUser, err
	}

	return createdUser, nil
}
