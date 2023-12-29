package controller

import (
	"net/http"
	"os"
	"time"
	"user-service/domain"
	"user-service/usecase"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	Signup(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu usecase.IUserUsecase
}

func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

func (uc *userController) Signup(c echo.Context) error {
	signupRequest := domain.SignupRequest{}
	if err := c.Bind(&signupRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	signupResponse, err := uc.uu.Signup(signupRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, signupResponse)
}

func (uc *userController) Login(c echo.Context) error {
	loginRequest := domain.LoginRequest{}
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	loginResponse, err := uc.uu.Login(loginRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = loginResponse.Token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) Logout(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
