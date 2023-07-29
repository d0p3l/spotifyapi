package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/zmb3/spotify/v2"
)

func (auth *Authentication) RecentlyPlayedTracks(ctx echo.Context) error {
	client := auth.GiveClient(ctx)
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	tracks, err := client.PlayerRecentlyPlayedOpt(ctx.Request().Context(), &spotify.RecentlyPlayedOptions{Limit: limit})
	if err != nil {
		return err
	}

	err = ctx.JSON(http.StatusOK, tracks)
	if err != nil {
		return err
	}

	return nil
}
