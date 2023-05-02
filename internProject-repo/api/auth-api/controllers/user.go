package controllers

import (
	"context"
	"example.com/m/pkg/data-access/database"
	user3 "example.com/m/pkg/models/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

func User() echo.HandlerFunc {
	return func(c echo.Context) error {
		claims, err := c.Cookie("user")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, nil)
		}

		var user user3.User

		client := database.DB
		usersCollection := client.Database("Login").Collection("credentials")

		searchAttributes := make(map[string]interface{})
		searchAttributes["username"] = claims.Value
		dbResponse := usersCollection.FindOne(context.TODO(), searchAttributes)
		if dbResponse.Err() != nil {
			return dbResponse.Err()
		}

		err = dbResponse.Decode(&user)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user)
	}
}
