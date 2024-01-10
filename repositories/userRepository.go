package repositories

import (
	"user-system/config"
	"user-system/entities"
	"gorm.io/gorm"
)
type UserRepository struct{
	db *config.Database 
}
func NewUserRepository(database *config.Database) *UserRepository {
	return &UserRepository{db : database}
}
func (userRepository *UserRepository) CreateUser(User *entities.User) (*entities.User, error){
	var result *gorm.DB

	result = userRepository.db.DB.Create(User)
	if result.Error != nil {
		return &entities.User{}, result.Error
	}
	return User, nil 
}
func (userRepository *UserRepository) ListUser() (*[]entities.User, error){
	var result *gorm.DB
	var users *[]entities.User

	result = userRepository.db.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}