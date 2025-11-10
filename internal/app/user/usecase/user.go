package usecase

import (
	"strings"

	"github.com/valentinusdelvin/savebite-be/internal/app/user/repository"
	"github.com/valentinusdelvin/savebite-be/internal/domain/dto"
	"github.com/valentinusdelvin/savebite-be/internal/domain/entity"
	"github.com/valentinusdelvin/savebite-be/internal/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseItf interface {
}

type userUsecase struct {
	userRepo repository.UserRepositoryItf
	jwt      jwt.JWTItf
}

func NewUserUsecase(userRepo repository.UserRepositoryItf, jwt jwt.JWTItf) UserUsecaseItf {
	return &userUsecase{
		userRepo: userRepo,
		jwt:      jwt,
	}
}

func (u *userUsecase) Register(register dto.Register) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := entity.User{
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     register.Email,
		Password:  string(hashedPassword),
	}

	if strings.Split(register.Email, "@")[1] == "savebite.com" {
		user.IsAdmin = true
	} else {
		user.IsAdmin = false
	}

	err = u.userRepo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) Login(login dto.Login) (string, error) {
	user, err := u.userRepo.GetUserByEmail(login.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return "", err
	}

	token, err := u.jwt.CreateToken(user.UserId, user.IsAdmin)
	if err != nil {
		return "", err
	}

	return token, nil
}
