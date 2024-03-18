package tests

import (
	"restapi/internal"
	"restapi/utils"
	"testing"

	"github.com/kataras/iris/v12/httptest"
)

type UserCredentials struct {
	login string
	password string
}

var credentials = UserCredentials{"test", "test"}

func TestGetProducts(t *testing.T) {
app := internal.IrisMainApp()
e := httptest.New(t, app)

e.GET("/products").Expect().Status(httptest.StatusOK).
JSON().Object().HasValue("fajne", "produkty")
}

func TestRegisterUser(t *testing.T) {
	app := internal.IrisMainApp()
	e := httptest.New(t, app)
	
	e.POST(utils.ApiPaths.Register).
	WithQuery("login", credentials.login).
	WithQuery( "password" , credentials.password).
	WithHeader("Content-Type","application/json").
	Expect().Status(httptest.StatusOK).
	JSON().Object().ContainsKey("message").HasValue("message", "Account created")

}

