package user

import "github.com/darkphotonKN/finance-analysis-dashboard/internal/shared/constants"

type CreateUserReq struct {
	Email     string `json:"email" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type UserSignInReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserSignInRes struct {
	User         UserRes `json:"user"`
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
}

type UserRes struct {
	FirstName string             `json:"firstName"`
	LastName  string             `json:"lastName"`
	Email     string             `json:"email"`
	Role      constants.UserRole `json:"role"`
}

type QueryFindAllUsers struct {
	Page     int                 `json:"page"`
	PageSize int                 `json:"pageSize"`
	Keyword  string              `json:"keyword"`
	Sort     string              `json:"sort"`
	Order    constants.SortOrder `json:"order"`
}
