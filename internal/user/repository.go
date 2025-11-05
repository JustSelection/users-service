package user

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id string) (*User, error)
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(id string) error
	EmailExists(email string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *User) (*User, error) {
	exists, err := r.EmailExists(user.Email)
	if err != nil {
		return nil, fmt.Errorf("email check failed: %w", err)
	}
	if exists {
		return nil, ErrEmailExists
	}
	err = r.db.Create(user).Error
	return user, err
}

func (r *userRepository) Delete(id string) error {
	result := r.db.Delete(&User{}, "id = ?", id)
	if result.RowsAffected == 0 {
		return ErrUserNotFound
	}
	return result.Error
}

func (r *userRepository) EmailExists(email string) (bool, error) {
	var count int64
	err := r.db.Model(&User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}

func (r *userRepository) GetAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetByID(id string) (*User, error) {
	var user User
	err := r.db.Where("id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrUserNotFound
	}
	return &user, err
}

func (r *userRepository) Update(user *User) (*User, error) {
	err := r.db.Save(user).Error
	return user, err
}
