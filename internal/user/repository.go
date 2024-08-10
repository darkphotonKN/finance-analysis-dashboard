package user

import (
	"fmt"

	"github.com/darkphotonKN/finance-analysis-dashboard/internal/shared/constants"
	"github.com/darkphotonKN/finance-analysis-dashboard/internal/shared/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	FindById(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindAllUsers(page int, pageSize int, keyword string, sort string, order constants.SortOrder) (*[]models.User, error)
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

// Find All Users
func (r *userRepository) FindAllUsers(page int, pageSize int, keyword string, sort string, order constants.SortOrder) (*[]models.User, error) {

	fmt.Printf("page: %d, pageSize: %d, sort: %s order: %s\n", page, pageSize, sort, order)

	var users []models.User

	// initialize query
	query := r.db.Model(&models.User{}) // a query for the table User, represented by the model

	// pagination
	if page > 0 && pageSize > 0 {
		offset := (page - 1) * pageSize
		query = query.Offset(offset).Limit(pageSize)
	}

	// filter by "keyword"
	query = query.Where("first_name LIKE ?", "%"+keyword)

	// sort by "sort" and direction "order"
	if sort != "" && (order == constants.ASC || order == constants.DESC) {
		fmt.Printf("sorting with %s, with order of %s\n", sort, order)
		query = query.Order(sort + " " + string(order))
	} else {
		fmt.Print("sorting with default email and asc.")
		query = query.Order("email asc")
	}

	query = query.Order("email asc")
	query = query.Debug()

	result := query.Find(&users)

	if result.Error != nil {

		return &users, result.Error
	}

	return &users, nil
}
