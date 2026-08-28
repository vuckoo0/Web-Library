package auth

import (
	"fmt"
	config "main/config"
	models "main/models"
	storage "main/storage"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func GenerateToken(userId int64) (string, error) {

	errorChanel := make(chan error, 1)
	resultChanel := make(chan int, 1)

	go storage.GetPrivilege(userId, errorChanel, resultChanel)

	err := <-errorChanel
	if err != nil {
		return "", err
	}

	privilege := <-resultChanel

	claims := jwt.MapClaims{
		"user_id":   userId,
		"privilege": privilege,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config().JWT_Secret))
}

func CheckUserCredentials(u *models.User) error {

	userId, err := storage.IsUsernameTaken(u.Name)
	if err != nil {
		return err
	}

	if userId == -1 {
		return fmt.Errorf("Bad credentials!")
	}

	hashedPassword, err := storage.GetPassword(userId)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(u.Password))
	if err != nil {
		return fmt.Errorf("Bad credentials!")
	}

	errorChanel := make(chan error, 1)
	resultChanel := make(chan int, 1)

	go storage.GetPrivilege(userId, errorChanel, resultChanel)

	err = <-errorChanel
	if err != nil {
		return err
	}

	u.Privilege = <-resultChanel
	u.Id = userId
	return nil
}
