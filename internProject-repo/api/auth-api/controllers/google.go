package controllers

//
//import (
//	"errors"
//	"example.com/m/api/auth-api/services"
//	"fmt"
//	"github.com/labstack/echo/v4"
//	"net/http"
//)
//
//func GoogleLogin() echo.HandlerFunc {
//	return func(c echo.Context) error {
//		googleConfig := services.SetupConfig()
//		url := googleConfig.AuthCodeURL("mystate")
//
//		fmt.Println(url)
//		return c.Redirect(http.StatusSeeOther, url)
//	}
//}
//
//func GoogleCallback() echo.HandlerFunc {
//	return func(c echo.Context) error {
//		//checking for the state
//		state := c.Request().URL.Query()["state"][0]
//		if state != "mystate" {
//			return errors.New("state dont match")
//		}
//		code := c.Request().URL.Query()["code"][0]
//
//		googleConfig := config
//	}
//}
