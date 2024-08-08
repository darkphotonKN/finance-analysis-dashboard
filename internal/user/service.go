package user

import (
	"fmt"

	"github.com/darkphotonKN/finance-analysis-dashboard/internal/shared/constants"
	"github.com/darkphotonKN/finance-analysis-dashboard/internal/shared/models"
	"github.com/darkphotonKN/finance-analysis-dashboard/internal/utils/auth"
)

type UserService interface {
	UserSignup(CreateUserReq) (*models.User, error)
	AuthenticateUser(userSignInReq UserSignInReq) (UserSignInRes, error)
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

	// find user from database
	user, err := s.userRepository.FindByEmail(createUserReq.Email)

	fmt.Printf("user %+v:", user)

	if user != nil {
		return user, fmt.Errorf("User already has an account.")
	}

	newUser := models.User{
		FirstName: createUserReq.FirstName,
		LastName:  createUserReq.LastName,
		Email:     createUserReq.Email,
		Password:  auth.HashPassword(createUserReq.Password),
		Role:      constants.UserRoleUser,
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
func (s *userService) AuthenticateUser(userSignInReq UserSignInReq) (UserSignInRes, error) {

	// find user from database
	user, err := s.userRepository.FindByEmail(userSignInReq.Email)

	if err != nil {
		return UserSignInRes{}, err
	}

	// authenticate password
	hashIncPw := auth.HashPassword(userSignInReq.Password)

	if hashIncPw != user.Password {
		return UserSignInRes{}, fmt.Errorf("Password was incorrect.")
	}

	// authenticated, generate jwt access token and refresh token
	accessToken, refreshToken, err := auth.GenerateJWT(user.ID)

	if err != nil {
		return UserSignInRes{}, err
	}

	// succesfully authenticated, return user and tokens
	return UserSignInRes{
		User: UserRes{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
