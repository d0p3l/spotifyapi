package handlers

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify/v2"
	"golang.org/x/oauth2"
)

func (auth *Authentication) giveClient(ctx echo.Context) *spotify.Client {
	tok := giveToken(ctx)
	return spotify.New(auth.spotifyauth.Client(ctx.Request().Context(), tok))
}

func giveToken(ctx echo.Context) *oauth2.Token {

	tok := &oauth2.Token{
		AccessToken:  ctx.Request().Header.Get("Authorization"),
		TokenType:    ctx.FormValue("token_type"),
		RefreshToken: ctx.FormValue("refresh_token"),
		Expiry:       time.Now(), // TODO сделать нормальный time date
	}

	return tok
}
