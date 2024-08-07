package user

import (
	"fmt"

	"github.com/darkphotonKN/finance-analysis-dashboard/internal/shared/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	FindById(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// Inserts New User
func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	fmt.Printf("Creating user in repo: %+v\n", *user)

	if err := r.db.Create(user).Error; err != nil {

		return nil, err
	}

	return user, nil
}

// Find User By Id
func (r *userRepository) FindById(id uint) (*models.User, error) {
	var user models.User

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return &models.User{}, err
	}

	return &user, nil
}

// Find User By Email
func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
