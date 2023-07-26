package app

import (
	"log"

	"github.com/d0p3l/spotifyapi/internal/app/handlers"
	"github.com/labstack/echo/v4"
)

type App struct {
	auth *handlers.Authentication
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.auth = handlers.New()
	a.echo = echo.New()

	// e.Use(authMiddleWare)
	a.echo.GET("api/login", a.auth.CompleteAuth)
	a.echo.GET("api/userinfo", a.auth.UserInfo)  // userInfo
	a.echo.GET("api/authurl", a.auth.GetAuthUrl) // getAuthUrl
	a.echo.GET("api/usertopartists", a.auth.UserTopArtists)

	return a, nil
}

func (a *App) Run() error {
	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
