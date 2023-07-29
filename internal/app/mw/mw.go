package mw

// import (
// 	"github.com/d0p3l/spotifyapi/internal/app/handlers"
// 	"github.com/labstack/echo/v4"
// )

// var auth *handlers.Authentication

// func authMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		client := auth.GiveClient(ctx)

// 		err := next(ctx)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	}
// }