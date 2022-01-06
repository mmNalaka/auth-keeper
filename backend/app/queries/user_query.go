package queries

import (
	"errors"
	"github.com/google/uuid"
	"github.com/mmnalaka/auth-keeper/app/models"
	"github.com/mmnalaka/auth-keeper/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CreateUser(email string, password string, role string) (*models.User, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		ID:       id,
		Email:    email,
		Password: hashedPassword,
		Role:     role,
	}

	if err := database.Postgres.Db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(e string) (*models.User, error) {
	var user *models.User

	if err := database.Postgres.Db.Where(&models.User{Email: e}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	if user.ID == uuid.Nil {
		return nil, nil
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
