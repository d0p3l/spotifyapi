package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify/v2"
)

func (auth *Authentication) UserInfo(ctx echo.Context) error {
	tok := auth.GiveToken(ctx)

	client := spotify.New(auth.auth.Client(ctx.Request().Context(), tok))
	user, err := client.CurrentUser(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.JSON(200, user.User)
}
