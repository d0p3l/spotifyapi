package handlers

import (
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify/v2"
)

func (auth *Authentication) UserInfo(ctx echo.Context) error {
	client := auth.GiveClient(ctx)
	user, err := client.CurrentUser(ctx.Request().Context())
	if err != nil {
		log.Println(err)
		return err
	}

	err = ctx.JSON(200, user.User)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (auth *Authentication) UserTopArtists(ctx echo.Context) error {
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	timerange := ctx.QueryParam("time_range")
	client := auth.GiveClient(ctx)
	user, err := client.CurrentUsersTopArtists(ctx.Request().Context(), spotify.Limit(limit), spotify.Timerange(spotify.Range(timerange)))
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

func (auth *Authentication) UserTopTracks(ctx echo.Context) error {
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	timerange := ctx.QueryParam("time_range")
	offset, _ := strconv.Atoi(ctx.QueryParam("offset"))
	client := auth.GiveClient(ctx)
	user, err := client.CurrentUsersTopTracks(ctx.Request().Context(), spotify.Limit(limit), spotify.Timerange(spotify.Range(timerange)), spotify.Offset(offset))
	if err != nil {
		log.Println(err)
		return err
	}

	err = ctx.JSON(200, user.Tracks)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
