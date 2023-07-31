package app

import (
	"log"

	"github.com/d0p3l/spotifyapi/internal/app/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	auth *handlers.Authentication
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.auth = handlers.New()
	a.echo = echo.New()

	a.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// e.Use(authMiddleWare)
	a.echo.GET("api/login", a.auth.CompleteAuth)
	a.echo.POST("api/userinfo", a.auth.UserInfo) // userInfo
	a.echo.GET("api/authurl", a.auth.GetAuthUrl) // getAuthUrl
	a.echo.POST("api/usertopartists", a.auth.UserTopArtists)
	a.echo.POST("api/usertoptracks", a.auth.UserTopTracks)
	a.echo.POST("api/recentlyplayedtracks", a.auth.RecentlyPlayedTracks)

	return a, nil
}

func (a *App) Run() error {
	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
