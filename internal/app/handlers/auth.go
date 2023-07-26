package handlers

import (
	"net/http"
	"time"

	"github.com/d0p3l/spotifyapi/internal/app/envconfig"
	"github.com/labstack/echo/v4"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

type Authentication struct {
	auth  *spotifyauth.Authenticator
	state string
}

const redirectURI = "http://localhost:8080/login"

func New() *Authentication {
	confauth := envconfig.New()
	return &Authentication{
		auth: spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate,
			spotifyauth.ScopeUserTopRead, spotifyauth.ScopeUserReadRecentlyPlayed, spotifyauth.ScopeUserLibraryRead),
			spotifyauth.WithClientID(confauth.SPOTIFY_ID), spotifyauth.WithClientSecret(confauth.SPOTIFY_SECRET)),
		state: "abc1234",
	}
}

func (auth *Authentication) CompleteAuth(e echo.Context) error {
	tok, err := auth.auth.Token(e.Request().Context(), auth.state, e.Request())
	if err != nil {
		// http.Error(e.Response().Writer, "Couldn't get token", http.StatusForbidden)
		return err
	}
	if st := e.FormValue("state"); st != auth.state {
		// http.NotFound(e.Response().Writer, e.Request())
		// log.Fatalf("State mismatch: %s != %s\n", st, state)
		return err
	}

	return e.JSON(200, tok)
}

func (auth *Authentication) GiveToken(ctx echo.Context) *oauth2.Token {
	// date, _ := time.Parse(time.RFC1123, ctx.QueryParam("expiry"))
	// ctx.QueryParam("token_type")
	// ctx.QueryParam("refresh_token")

	tok := &oauth2.Token{
		AccessToken:  ctx.Request().Header.Get("Authorization"),
		TokenType:    ctx.FormValue("token_type"),
		RefreshToken: ctx.FormValue("refresh_token"),
		Expiry:       time.Now(), // TODO сделать нормальный time date
	}

	return tok
}

func (auth *Authentication) GetAuthUrl(ctx echo.Context) error {
	url := auth.auth.AuthURL(auth.state)
	return ctx.String(http.StatusOK, url)
}
