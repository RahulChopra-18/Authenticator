package controllers

import (
	"example.com/m/api/auth-api/services"
	"example.com/m/pkg/data-access/database"
	user2 "example.com/m/pkg/models/user"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

// SignInForm responsible for signIn Form rendering.
func SignInForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		tmpl, err := template.ParseFiles("api/auth-api/templates/signIn.html")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if err := tmpl.Execute(c.Response().Writer, nil); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return nil
	}
}

// SignIn will be executed after SignInForm submission.
func SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Initiate a new User struct.
		u := new(user2.Check)
		// Parse the submitted data and fill the User struct with the data from the SignIn form.
		if err := c.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		storedUser, err := database.LoadUser(u)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "User not found")
		}
		// Compare the stored hashed password, with the hashed version of the password that was received
		if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(u.Password)); err != nil {
			// If the two passwords don't match, return a 401 status
			return echo.NewHTTPError(http.StatusUnauthorized, "Password is incorrect")
		}
		// If password is correct, generate tokens and set cookies.
		err = services.GenerateTokensAndSetCookies(storedUser, c)

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Token is incorrect")
		}

		return c.JSON(http.StatusOK, "success")
		//	return c.Redirect(http.StatusMovedPermanently, "/admin")
	}
}
