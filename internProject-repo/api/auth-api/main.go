package main

import (
	"example.com/m/api/auth-api/controllers"
	"example.com/m/api/auth-api/services"
	"example.com/m/pkg/data-access/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	//connecting to the database
	database.Connect()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:          nil,
		AllowOrigins:     []string{"*"},
		AllowOriginFunc:  nil,
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders:     nil,
		AllowCredentials: true,
		ExposeHeaders:    nil,
		MaxAge:           0,
	}))
	//e.GET("/google/login", controllers.GoogleLogin()) //this will redirect me to google login page
	//e.GET("google/callback", controllers.GoogleCallback) // this will be invoked by google after successful login
	e.POST("/user/register", controllers.Register())
	e.GET("/user/signin", controllers.SignInForm()).Name = "userSignInForm"
	e.POST("/user/signin", controllers.SignIn())
	e.GET("/user/user", controllers.User())

	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &services.Claims{},
		SigningKey:              []byte(services.GetJWTSecret()),
		TokenLookup:             "cookie:access-token",
		ErrorHandlerWithContext: services.JWTErrorChecker,
	}))

	// Attach jwt token refresher.
	adminGroup.Use(services.TokenRefresherMiddleware)

	adminGroup.GET("", controllers.Admin())

	e.Logger.Fatal(e.Start(":3001"))

}
