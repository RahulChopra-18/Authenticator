package controllers

import (
	"context"
	"example.com/m/pkg/data-access/database"
	user3 "example.com/m/pkg/models/user"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var data map[string]string
		client := database.DB
		usersCollection := client.Database("Login").Collection("credentials")
		// Parse the submitted data and populate the data map
		if err := c.Bind(&data); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

		user := user3.User{
			Password: password,
			Username: data["name"],
			Email:    data["email"],
		}

		if _, err := usersCollection.InsertOne(context.TODO(), user); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, user)
	}
}
