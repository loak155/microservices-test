package usecase

import (
	"os"
	"time"
	"user-service/domain"
	"user-service/repository"
	"user-service/validator"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Signup(user domain.SignupRequest) (domain.SignupResponse, error)
	Login(user domain.LoginRequest) (domain.LoginResponse, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) Signup(req domain.SignupRequest) (domain.SignupResponse, error) {
	if err := uu.uv.SignupRequestValidate(req); err != nil {
		return domain.SignupResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return domain.SignupResponse{}, err
	}
	newUser := domain.User{Username: req.Username, Email: req.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return domain.SignupResponse{}, err
	}
	res := domain.SignupResponse{
		ID:       newUser.ID,
		Username: newUser.Username,
		Email:    newUser.Email,
	}
	return res, nil
}

func (uu *userUsecase) Login(req domain.LoginRequest) (domain.LoginResponse, error) {
	if err := uu.uv.LoginRequestValidate(req); err != nil {
		return domain.LoginResponse{}, err
	}
	storedUser := domain.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, req.Email); err != nil {
		return domain.LoginResponse{}, err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(req.Password))
	if err != nil {
		return domain.LoginResponse{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return domain.LoginResponse{}, err
	}
	res := domain.LoginResponse{
		UserID: storedUser.ID,
		Token:  tokenString,
	}
	return res, nil
}
