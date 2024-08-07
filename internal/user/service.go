package user

import (
	"fmt"

	"github.com/darkphotonKN/finance-analysis-dashboard/internal/shared/models"
	"github.com/darkphotonKN/finance-analysis-dashboard/internal/util/auth"
)

type UserService interface {
	UserSignup(CreateUserReq) (*models.User, error)
	AuthenticateUser(userSignInReq UserSignInReq) (*models.User, error)
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
		Password:  auth.HashPassword(createUserReq.Password),
		Role:      "user",
	}

	fmt.Printf("Creating user in service: %+v\n", newUser)

	createdUser, err := s.userRepository.CreateUser(&newUser)

	if err != nil {
		return createdUser, err
	}

	return createdUser, nil
}

/**
* Authenticate User
**/
func (s *userService) AuthenticateUser(userSigninReq UserSignInReq) (*models.User, error) {

	// find user from database
	user, err := s.userRepository.FindByEmail(userSigninReq.Email)

	if err != nil {
		return user, err
	}

	// authenticate password
	hashIncPw := auth.HashPassword(userSigninReq.Password)

	if hashIncPw != user.Password {
		return user, fmt.Errorf("Password was incorrect.")
	}

	// succesfully authenticated, return user
	return user, nil
}
