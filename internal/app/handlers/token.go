package handlers

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
)

func (auth *Authentication) GiveClient(ctx echo.Context) *spotify.Client {
	tok := giveFullToken(ctx)
	return spotify.New(auth.spotifyauth.Client(ctx.Request().Context(), tok))
}

func giveFullToken(ctx echo.Context) *oauth2.Token {
	expiry, _ := time.Parse(time.RFC3339, ctx.FormValue("expiry"))

	tok := &oauth2.Token{
		AccessToken:  ctx.Request().Header.Get("Authorization"),
		TokenType:    ctx.FormValue("token_type"),
		RefreshToken: ctx.FormValue("refresh_token"),
		Expiry:       expiry,
	}

	return tok
}
