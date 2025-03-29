package controller

import (
	"learning-go-rest-api/model"
	"learning-go-rest-api/useCase"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

type IUserController interface {
	SignUp(c echo.Context) error
	Login(c echo.Context) error
	Logout(c echo.Context) error
}

type userController struct {
	uu useCase.IUserUseCase
}

func NewUserController(uu useCase.IUserUseCase) IUserController {
	return &userController{uu: uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}

	// リクエストボディを構造体にマッピングさせる
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) Login(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// cookieにJWTトークンを書き込み
	cookie := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(24 * time.Hour),
		Domain:  os.Getenv("API_DOMAIN"),
		Path:    "/", // サーバー全体でcookieが有効
		// Secure:   true, // trueにするとhttpsでのみ送信される 開発中はhttpで通信するのでfalseにする
		HttpOnly: true,                  // trueにするとJavaScriptからアクセスできなくなる
		SameSite: http.SameSiteNoneMode, // フロント/バックエンドでドメインが異なるのでSamesiteNoneモードにしておく
	}
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

func (uc *userController) Logout(c echo.Context) error {
	// cookieからtokenを削除
	cookie := &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Now(),
		Domain:  os.Getenv("API_DOMAIN"),
		Path:    "/",
		// Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}
