package router

import (
	"learning-go-rest-api/controller"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// アクセスを許可するオリジン
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		// 有効なヘッダー
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		// 有効なメソッド
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true, // cookieを使用可能にする
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode, // postmanで動作確認するときはオフにする　オンにするとsecure属性がオンになってしまう
		// CookieSameSite: http.SameSiteDefaultMode,
		// CookieMaxAge:   60,
	}))
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)
	e.GET("/csrf", uc.CsrfToken)
	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)
	return e
}
