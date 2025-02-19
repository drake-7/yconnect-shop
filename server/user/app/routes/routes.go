package routes

import (
	"net/http"
	"github.com/NetSinx/yconnect-shop/server/user/app/config"
	"github.com/NetSinx/yconnect-shop/server/user/controller"
	"github.com/NetSinx/yconnect-shop/server/user/repository"
	"github.com/NetSinx/yconnect-shop/server/user/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ApiRoutes() *echo.Echo {
	userRepository := repository.UserRepository(config.ConnectDB())
	userService := service.UserService(userRepository)
	userController := controller.UserController(userService)

	router := echo.New()
	router.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "cookie:xsrf",
		CookiePath: "/",
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteStrictMode,
		CookieSecure: true,
	}))

	router.GET("/gencsrf", func(c echo.Context) error {
		csrfToken := c.Get("csrf")

		return c.JSON(200, map[string]interface{}{
			"csrf_token": csrfToken,
		})
	})
	router.POST("/user/sign-up", userController.RegisterUser)
	router.POST("/user/sign-in", userController.LoginUser)
	router.GET("/user", userController.ListUsers)
	router.GET("/user/:id", userController.GetUser)
	router.PUT("/user/:id", userController.UpdateUser)
	router.DELETE("/user/:id", userController.DeleteUser)

	return router
}