package handlers

import (
	"errors"
	"net/http"

	"github.com/d0p3l/spotifyapi/internal/app/envconfig"
	"github.com/labstack/echo/v4"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

type Authentication struct {
	spotifyauth *spotifyauth.Authenticator
	state       string
}

const redirectURI = "http://localhost:8080/api/login"

func New() *Authentication {
	confauth := envconfig.New()
	return &Authentication{
		spotifyauth: spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopeUserTopRead, spotifyauth.ScopeUserReadRecentlyPlayed, spotifyauth.ScopeUserLibraryRead),
			spotifyauth.WithClientID(confauth.SPOTIFY_ID), spotifyauth.WithClientSecret(confauth.SPOTIFY_SECRET)),
		state: "abc1234", // сделать рандомным?
	}
}

func (auth *Authentication) CompleteAuth(ctx echo.Context) error {
	code := ctx.QueryParam("code")
	actualState := ctx.QueryParam("state")
	if actualState != auth.state {
		return errors.New("spotify: redirect state parameter doesn't match")
	}
	tok, err := auth.spotifyauth.Exchange(ctx.Request().Context(), code)

	// tok, err := auth.spotifyauth.Token(ctx.Request().Context(), auth.state, ctx.Request())
	if err != nil {
		return err
	}
	// if st := ctx.FormValue("state"); st != auth.state {
	// 	return err
	// }

	err = ctx.JSON(200, tok)
	if err != nil {
		return err
	}

	return nil
}

func (auth *Authentication) GetAuthUrl(ctx echo.Context) error {
	url := auth.spotifyauth.AuthURL(auth.state)
	return ctx.String(http.StatusOK, url)
}
