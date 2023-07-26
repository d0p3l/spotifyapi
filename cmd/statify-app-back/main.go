package main

import (
	"log"

	"github.com/d0p3l/spotifyapi/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// func authMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(ctx echo.Context) error {
// 		tok := giveToken(ctx)

// 		err := next(ctx)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	}
// }

// func refreshToken(ctx echo.Context) error {

// }
