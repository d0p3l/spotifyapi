package handlers

import (
	"log"

	"github.com/labstack/echo/v4"
)

func (auth *Authentication) UserInfo(ctx echo.Context) error {
	client := auth.giveClient(ctx)
	user, err := client.CurrentUser(ctx.Request().Context())
	if err != nil {
		return err
	}

	err = ctx.JSON(200, user.User)
	if err != nil {
		return err
	}

	return nil
}

func (auth *Authentication) UserTopArtists(ctx echo.Context) error {
	client := auth.giveClient(ctx)
	user, err := client.CurrentUsersTopArtists(ctx.Request().Context())
	if err != nil {
		log.Println(err)
		return err
	}

	err = ctx.JSON(200, user.Artists)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
